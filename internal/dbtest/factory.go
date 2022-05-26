package dbtest

import (
	"fmt"
	"time"

	"github.com/bluele/factory-go/factory"

	"aegis/internal/model"
	"aegis/internal/proto"
)

var (
	atAttr = func(args factory.Args) (interface{}, error) { return time.Now(), nil }
	// seqInt64Attr   = func(n int) (interface{}, error) { return int64(n), nil }
	sprintfSeqAttr = func(layout string) func(int) (interface{}, error) {
		return func(n int) (interface{}, error) {
			return fmt.Sprintf(layout, n), nil
		}
	}
	identityAttr = func(i interface{}) func(factory.Args) (interface{}, error) {
		return func(args factory.Args) (interface{}, error) {
			return i, nil
		}
	}
)

var PropertyFactory = factory.NewFactory(&model.Property{}).
	SeqInt("Name", sprintfSeqAttr("property_%d")).
	SeqInt("Label", sprintfSeqAttr("属性_%d")).
	Attr("Type", identityAttr(int32(proto.PropertyTypeString))).
	Attr("ValidateType", identityAttr(int32(proto.ValidateTypeString))).
	Attr("CreatedAt", atAttr).
	Attr("UpdatedAt", atAttr)
