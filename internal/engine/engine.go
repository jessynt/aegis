package engine

import (
	"context"
	"fmt"
	"strings"
	"time"

	gkitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"

	"aegis/internal/engine/abstraction"
	"aegis/internal/engine/activation"
	"aegis/internal/engine/aggregation"
	"aegis/internal/engine/bag"
	"aegis/internal/engine/collection"
	"aegis/internal/engine/contract"
	"aegis/internal/engine/expression"
	"aegis/internal/engine/metric"
	"aegis/internal/engine/model"
	"aegis/internal/engine/prepare/plugins"
	"aegis/internal/engine/property"
	"aegis/internal/engine/rule"
	"aegis/internal/engine/store"
	"aegis/internal/engine/validation"
	eventDao "aegis/internal/module/event_dao"
	"aegis/internal/proto"
	"aegis/pkg/json"
	"aegis/pkg/profile"
)

var (
	ErrModelNotExists = errors.New("model not exists")
)

type Opt func(e *Engine)

type Engine struct {
	store              *store.Store
	logger             gkitLog.Logger
	modelManager       contract.ModelManager
	propertyManager    contract.PropertyManager
	abstractionManager contract.AbstractionManager
	activationManager  contract.ActivationManager
	ruleManager        contract.RuleManager
	aggregator         contract.Aggregator
	collectionManager  contract.CollectionManager
}

func NewEngine(opts ...Opt) *Engine {
	e := &Engine{
		logger:             gkitLog.NewNopLogger(),
		modelManager:       model.NewManager(),
		propertyManager:    property.NewManager(),
		abstractionManager: abstraction.NewManager(),
		activationManager:  activation.NewManager(),
		ruleManager:        rule.NewManager(),
		aggregator:         aggregation.NewAggregator(),
		collectionManager:  collection.NewManager(),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func WithLogger(logger gkitLog.Logger) Opt {
	return func(e *Engine) {
		e.logger = logger
	}
}

func WithModelManager(modelManager contract.ModelManager) Opt {
	return func(e *Engine) {
		e.modelManager = modelManager
	}
}

func WithPropertyManager(propertyManager contract.PropertyManager) Opt {
	return func(e *Engine) {
		e.propertyManager = propertyManager
	}
}

func WithAbstractionManager(abstractionManager contract.AbstractionManager) Opt {
	return func(e *Engine) {
		e.abstractionManager = abstractionManager
	}
}

func WithActivationManager(activationManager contract.ActivationManager) Opt {
	return func(e *Engine) {
		e.activationManager = activationManager
	}
}

func WithRuleManager(ruleManager contract.RuleManager) Opt {
	return func(e *Engine) {
		e.ruleManager = ruleManager
	}
}

func WithAggregator(aggregator contract.Aggregator) Opt {
	return func(e *Engine) {
		e.aggregator = aggregator
	}
}

func WithCollectionManager(collectionManager contract.CollectionManager) Opt {
	return func(e *Engine) {
		e.collectionManager = collectionManager
	}
}

func WithStore(s *store.Store) Opt {
	return func(e *Engine) {
		e.store = s
	}
}

func (e *Engine) Report(ctx context.Context) error {
	b := bag.FromContext(ctx)

	rGUID, ok := b.GetProperty("GUID")
	if !ok {
		return ErrModelNotExists
	}

	m, ok := e.modelManager.GetModelByGuid(cast.ToString(rGUID))
	if !ok {
		return ErrModelNotExists
	}

	if err := e.validateProperties(ctx); err != nil {
		return err
	}

	if err := e.executePrepares(ctx); err != nil {
		return err
	}
	properties := b.GetProperties()
	t := time.Unix(0, cast.ToInt64(properties["Time"])*int64(time.Millisecond))

	record := make(map[string]interface{})
	record["ModelId"] = m.GetModelId()
	record["Year"] = t.Year()
	record["Quarter"] = (uint8(t.Month()) + 2) / 3
	record["Month"] = uint8(t.Month())
	record["DayOfMonth"] = t.Day()
	record["DayOfWeek"] = uint8(t.Weekday())

	for k, v := range properties {
		if k == "GUID" {
			continue
		}
		p, exists := e.propertyManager.GetByName(k)
		if !exists {
			continue
		}

		var value interface{}
		switch p.GetType() {
		case proto.PropertyTypeInteger:
			value = cast.ToInt64(v)
		case proto.PropertyTypeDouble:
			value = cast.ToFloat64(v)
		case proto.PropertyTypeString:
			value = cast.ToString(v)
		case proto.PropertyTypeBool:
			value = cast.ToBool(v)
		case proto.PropertyTypeDate:
			value = cast.ToString(v)
		case proto.PropertyTypeDateTime:
			value = time.Unix(0, cast.ToInt64(v)*int64(time.Millisecond))
		default:
			continue
		}
		record[k] = value
	}

	return eventDao.InsertEvent(ctx, e.store.ClickHouseConn(), record)
}

func (e *Engine) Check(ctx context.Context) error {
	b := bag.FromContext(ctx)

	rGUID, ok := b.GetProperty("GUID")
	if !ok {
		return ErrModelNotExists
	}

	m, ok := e.modelManager.GetModelByGuid(cast.ToString(rGUID))
	if !ok {
		return ErrModelNotExists
	}

	metric.EngineModelEvaluateCount.WithLabelValues(m.GetModelGUID()).Inc()

	modelId := m.GetModelId()

	if err := e.validateProperties(ctx); err != nil {
		return err
	}

	if err := e.executePrepares(ctx); err != nil {
		return err
	}

	if err := e.executeAbstractions(ctx, modelId); err != nil {
		return err
	}
	_ = level.Debug(e.logger).Log("abstractions", string(json.MustMarshal(b.GetAbstractions())))

	if err := e.executeActivation(ctx, modelId); err != nil {
		return err
	}

	return nil
}

func (e *Engine) executePrepares(ctx context.Context) error {
	b := bag.FromContext(ctx)

	if ip, exists := b.GetProperty("IP"); exists {
		locationMap, err := plugins.IP2Location(cast.ToString(ip))
		if err != nil {
			_ = level.Warn(e.logger).Log(
				"message", fmt.Sprintf("IP2Location: %#v process failed", ip),
				"error", err,
				"culprit", "engine.executePrepares",
			)
			return nil
		}

		if _, ok := b.GetProperty("Country"); !ok {
			b.StoreProperty("Country", locationMap["Country"])
		}

		if _, ok := b.GetProperty("Province"); !ok {
			b.StoreProperty("Province", locationMap["Province"])
		}

		if _, ok := b.GetProperty("City"); !ok {
			b.StoreProperty("City", locationMap["City"])
		}
	}
	return nil
}

func (e *Engine) executeAbstractions(ctx context.Context, modelId int64) error {
	defer profile.Duration(time.Now(), "executeAbstractions")

	b := bag.FromContext(ctx)

	abstractions := e.abstractionManager.GetAbstractionsByModelId(modelId)
	for _, a := range abstractions {
		// 搜索字段
		searchField := a.GetSearchField()
		searchField = strings.ReplaceAll(searchField, "properties.", "")
		searchFieldValue, ok := b.GetProperty(searchField)
		if !ok {
			return fmt.Errorf("property `%s` not exists", searchField)
		}

		// 聚合类型
		aggregateType := a.GetAggregateType()

		// 聚合字段
		aggregateField := a.GetAggregateField()
		aggregateField = strings.ReplaceAll(aggregateField, "properties.", "")
		aggregateFieldValue, ok := b.GetProperty(aggregateField)
		if !ok && aggregateType != proto.AggregateTypeCount {
			return fmt.Errorf("property `%s` not exists", aggregateField)
		}

		// 时间片
		intervalType := a.GetAggregateIntervalType()
		intervalValue := a.GetAggregateIntervalValue()

		// 过滤条件表达式
		// filterExpression := a.GetFilterExpression()
		propertyTime, ok := b.GetProperty("Time")
		if !ok {
			return fmt.Errorf("property `Time` not exists")
		}

		endsAt := time.Unix(0, cast.ToInt64(propertyTime)*int64(time.Millisecond))
		startsAt := contract.SubTimeByInterval(endsAt, intervalType, intervalValue)

		switch aggregateType {
		case proto.AggregateTypeCount:
			c, err := e.aggregator.Count(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt)
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		case proto.AggregateTypeDistinctCount:
			c, err := e.aggregator.DistinctCount(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		case proto.AggregateTypeAverage:
			c, err := e.aggregator.Average(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		case proto.AggregateTypeSum:
			c, err := e.aggregator.Sum(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		case proto.AggregateTypeMin:
			c, err := e.aggregator.Min(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		case proto.AggregateTypeMax:
			c, err := e.aggregator.Max(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		case proto.AggregateTypeSD:
			c, err := e.aggregator.SD(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		case proto.AggregateTypeVariance:
			c, err := e.aggregator.Variance(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		case proto.AggregateTypeDeviation:
			c, err := e.aggregator.Deviation(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField, cast.ToFloat64(aggregateFieldValue))
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		case proto.AggregateTypeMedian:
			c, err := e.aggregator.Median(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
			if err != nil {
				return err
			}
			b.StoreAbstraction(a.GetAbstractionName(), c)
		default:
			return errors.New("aggregate type not support")
		}
	}

	return nil
}

func (e *Engine) executeActivation(ctx context.Context, modelId int64) error {
	defer profile.Duration(time.Now(), "executeActivation")
	activations := e.activationManager.GetActivationsByModelId(modelId)
	b := bag.FromContext(ctx)

	env := make(expression.Env)
	env["abstractions"] = b.GetAbstractions()
	env["properties"] = b.GetProperties()

	cs := make(map[string]interface{})
	for _, c := range e.collectionManager.GetCollections() {
		cs[c.Name()] = c
	}
	env["collections"] = cs

	for _, a := range activations {
		rules := e.ruleManager.GetRulesByActivationId(a.Id())
		score := decimal.Zero
		for _, r := range rules {
			expr := r.GetExpression()
			matched, err := expression.Check(expr, env)
			if err != nil {
				_ = level.Error(e.logger).Log(
					"message", "expression eval failed",
					"error", err,
					"expr", expr,
					"env", env,
				)
				continue
			}
			// 命中风控规则，计算得分
			if matched {
				operator := r.GetOperatorType()
				baseScope := decimal.New(r.GetBaseScore(), 0)
				rate := decimal.NewFromFloat(r.GetRate())
				baseNum := decimal.New(r.GetBaseNum(), 0)
				extra := decimal.Zero

				ab := e.abstractionManager.GetAbstractionById(r.GetAbstractionId())
				if ab != nil {
					abs, ok := b.GetAbstractionByKey(ab.GetAbstractionName())
					if ok {
						extra = decimal.NewFromFloat(abs.(float64))
					}
				}

				extra = extra.Mul(rate)
				switch operator {
				case proto.OperatorTypeAdd:
					extra = baseNum.Add(extra)
				case proto.OperatorTypeSub:
					extra = baseNum.Sub(extra)
				case proto.OperatorTypeMul:
					extra = baseNum.Mul(extra)
				case proto.OperatorTypeDiv:
					extra = baseNum.DivRound(extra, 2)
				}
				score = score.Add(baseScope.Add(extra))
				_ = e.logger.Log("message", fmt.Sprintf("%s 匹配, 得分: %s", r.GetRuleLabel(), baseScope.Add(extra)))
			}
		}
		blockScore := decimal.New(a.GetBlockScore(), 0)
		warningScore := decimal.New(a.GetWarningScore(), 0)

		scoreInFloat64, _ := score.Round(2).Float64()
		riskObject := &proto.RiskObject{
			Score: scoreInFloat64,
		}

		if score.GreaterThanOrEqual(blockScore) {
			riskObject.RickType = proto.RiskDeny
		} else if score.GreaterThanOrEqual(warningScore) {
			riskObject.RickType = proto.RiskWarning
		} else {
			riskObject.RickType = proto.RiskPass
		}

		b.StoreActivation(a.GetActivationName(), riskObject)
	}
	return nil
}

func (e *Engine) validateProperties(ctx context.Context) error {
	b := bag.FromContext(ctx)

	for k, v := range b.GetProperties() {
		if k == "GUID" {
			continue
		}

		p, ok := e.propertyManager.GetByName(k)
		if !ok {
			_ = level.Error(e.logger).Log("message", fmt.Sprintf("property `%s` not exists in our database", k))
			continue
		}

		if err := validation.Validate(v, p); err != nil {
			return fmt.Errorf("property `%s` validate failed: %w", k, err)
		}
	}
	return nil
}
