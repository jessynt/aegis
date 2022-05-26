package contract

import (
	"time"

	"aegis/internal/proto"
)

func SubTimeByInterval(t time.Time, intervalType proto.AggregateIntervalType, intervalValue int64) time.Time {
	switch intervalType {
	case proto.AggregateIntervalTypeYear:
		return t.AddDate(-int(intervalValue), 0, 0)
	case proto.AggregateIntervalTypeMonth:
		return t.AddDate(0, -int(intervalType), 0)
	case proto.AggregateIntervalTypeDay:
		return t.AddDate(0, 0, -int(intervalValue))
	case proto.AggregateIntervalTypeHour:
		return t.Add(-(time.Duration(intervalValue) * time.Hour))
	case proto.AggregateIntervalTypeMinute:
		return t.Add(-(time.Duration(intervalValue) * time.Minute))
	case proto.AggregateIntervalTypeSecond:
		return t.Add(-(time.Duration(intervalValue) * time.Second))
	default:
		// no filter
		return time.Time{}
	}
}
