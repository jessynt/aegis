package plugins

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIP2Location(t *testing.T) {
	rv, err := IP2Location("121.35.101.91")
	require.NoError(t, err)

	require.Equal(t, "中国", rv["Country"])
	require.Equal(t, "广东", rv["Province"])
	require.Equal(t, "深圳", rv["City"])
}

func BenchmarkIP2Location(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IP2Location("121.35.101.91")
	}
}
