package mysqlx

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

// This implementation is based on:
//
// - github.com/russross/meddler
// - github.com/jmoiron/sqlx/reflectx
//
// KNOWN ISSUES:
//
// - nested struct fields are not supported

const (
	pkOptionName = "pk"
)

var (
	errPointerRequired   = errors.New("must pass a pointer, not a value")
	errNilPassed         = errors.New("nil pointer passed")
	errNilMapper         = errors.New("nil mapper passed")
	errNotEmptyPk        = errors.New("primary key must be empty")
	errUnsupportedPkType = errors.New("unsupported pk type")
	errNoColumns         = errors.New("no columns can be inserted")
)

func isEmptyPkValue(field reflect.Value) (bool, error) {
	switch field.Type().Kind() {
	// NOTE: current only int is supported
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	default:
		// unsupported type
		return false, errUnsupportedPkType
	}

	return field.Int() == 0, nil
}

func isPrimaryKeyField(field *reflectx.FieldInfo) bool {
	_, exists := field.Options[pkOptionName]
	return exists
}

func isFirstLevelField(field *reflectx.FieldInfo) bool {
	return len(field.Index) == 1
}

func getPrimaryKey(
	tm *reflectx.StructMap,
	srcValue reflect.Value,
) (hasPk bool, pkKey string, isEmptyPk bool, err error) {
	hasPk = false
	for _, fi := range tm.Index {
		if !isPrimaryKeyField(fi) {
			// not a pk field, skip
			continue
		}

		if hasPk {
			err = fmt.Errorf("duplicate pk found: %s, %s", pkKey, fi.Name)
			return
		}
		hasPk = true
		pkKey = fi.Name
		isEmptyPk, err = isEmptyPkValue(reflectx.FieldByIndexesReadOnly(srcValue, fi.Index))
		if err != nil {
			return
		}
	}

	return
}

// NOTE: pkValue only support int64/uint64 for now, as is `sql.Result.LastInsertId`'s value.
func setPrimaryKey(tm *reflectx.StructMap, src interface{}, pkKey string, pkValue interface{}) error {
	field, exists := tm.Names[pkKey]
	if !exists {
		return fmt.Errorf("missing primary key field: %s", pkKey)
	}
	if !isPrimaryKeyField(field) {
		return fmt.Errorf("not a primary key field: %s", pkKey)
	}

	var value = reflect.ValueOf(src).Elem()
	for _, index := range field.Index {
		value = value.Field(index)
	}
	switch value.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(pkValue.(int64))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(pkValue.(uint64))
	default:
		return errUnsupportedPkType
	}

	return nil
}

func gatherValues(tm *reflectx.StructMap, srcValue reflect.Value) (values []interface{}) {
	for _, fi := range tm.Index {
		if isPrimaryKeyField(fi) {
			// NOTE: exclude pk field
			continue
		}

		if !isFirstLevelField(fi) {
			// NOTE: exclude non first level fields
			continue
		}

		field := reflectx.FieldByIndexesReadOnly(srcValue, fi.Index)
		// TODO(hbc): prepare value for insert/update
		values = append(values, field.Interface())
	}

	return
}

func generateInsertQuery(table string, tm *reflectx.StructMap) (string, error) {
	var names []string
	for _, fi := range tm.Index {
		if isPrimaryKeyField(fi) {
			// NOTE: exclude pk field
			continue
		}

		if !isFirstLevelField(fi) {
			// NOTE: exclude non first level fields
			continue
		}

		names = append(names, quoted(fi.Name))
	}

	if len(names) < 1 {
		return "", errNoColumns
	}

	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		quoted(table),
		strings.Join(names, ","),
		strings.Repeat(placeholderString+",", len(names)-1)+placeholderString,
	), nil
}

// NOTE: extracted from https://github.com/jmoiron/sqlx/blob/cf35089a197953c69420c8d0cecda90809764b1d/sqlx.go#L151
// ongoing issue: https://github.com/jmoiron/sqlx/issues/397
func mapperFor(i interface{}) *reflectx.Mapper {
	switch i := i.(type) {
	case sqlx.DB:
		return i.Mapper
	case *sqlx.DB:
		return i.Mapper
	case sqlx.Tx:
		return i.Mapper
	case *sqlx.Tx:
		return i.Mapper
	default:
		return nil
	}
}

// 插入一个数据到指定数据表.
func Insert(
	ctx context.Context,
	db sqlx.ExecerContext,
	table string,
	src interface{},
) error {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Ptr {
		return errPointerRequired
	}
	if srcValue.IsNil() {
		return errNilPassed
	}

	mapper := mapperFor(db)
	if mapper == nil {
		return errNilMapper
	}
	tm := mapper.TypeMap(reflect.Indirect(srcValue).Type())

	hasPk, pkKey, isEmptyPk, err := getPrimaryKey(tm, srcValue)
	if err != nil {
		return err
	}
	if hasPk && !isEmptyPk {
		return errNotEmptyPk
	}

	values := gatherValues(tm, srcValue)

	q, err := generateInsertQuery(table, tm)
	if err != nil {
		return err
	}

	if hasPk {
		result, err := db.ExecContext(ctx, q, values...)
		if err != nil {
			return err
		}

		newPk, err := result.LastInsertId()
		if err != nil {
			return err
		}
		if err = setPrimaryKey(tm, src, pkKey, newPk); err != nil {
			return err
		}
	} else {
		// no primary key, so no need to lookup new value
		_, err := db.ExecContext(ctx, q, values...)
		if err != nil {
			return err
		}
	}

	return nil
}
