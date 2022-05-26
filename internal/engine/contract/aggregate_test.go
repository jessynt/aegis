package contract

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"aegis/internal/proto"
)

func TestSubTimeByInterval(t *testing.T) {
	t1 := time.Unix(1136214245, 0)

	testCases := []struct {
		intervalType  proto.AggregateIntervalType
		intervalValue int64
		exceptResult  time.Time
	}{
		{proto.AggregateIntervalTypeYear, 1, time.Unix(1136214245-86400*365, 0)},
		{proto.AggregateIntervalTypeDay, 1, time.Unix(1136214245-86400, 0)},
		{proto.AggregateIntervalTypeDay, 3, time.Unix(1136214245-86400*3, 0)},
		{proto.AggregateIntervalTypeHour, 3, time.Unix(1136214245-3600*3, 0)},
		{proto.AggregateIntervalTypeMinute, 30, time.Unix(1136214245-60*30, 0)},
		{proto.AggregateIntervalTypeMinute, 3000, time.Unix(1136214245-60*3000, 0)},
		{proto.AggregateIntervalTypeMinute, 1, time.Unix(1136214245-60*1, 0)},
		{proto.AggregateIntervalTypeSecond, 3, time.Unix(1136214245-3, 0)},
	}
	for _, testCase := range testCases {
		rv := SubTimeByInterval(t1, testCase.intervalType, testCase.intervalValue)

		require.True(t, rv.Equal(testCase.exceptResult))
	}
}
