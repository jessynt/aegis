package hashset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet_Size(t *testing.T) {
	set := New()
	set.Add("foo")
	require.EqualValues(t, 1, set.Size())
	set.Add("foo")
	require.EqualValues(t, 1, set.Size())
	set.Add("bar")
	require.EqualValues(t, 2, set.Size())
}

func TestSet_Clear(t *testing.T) {
	set := New()
	set.Add("foo")
	require.EqualValues(t, 1, set.Size())
	set.Add("foo")
	require.EqualValues(t, 1, set.Size())
	set.Add("bar")
	require.EqualValues(t, 2, set.Size())
	set.Clear()
	require.EqualValues(t, 0, set.Size())
}

func TestSet_Has(t *testing.T) {
	set := New()
	set.Add("foo")
	require.True(t, set.Contains("foo"))
	require.False(t, set.Contains("bar"))
	require.EqualValues(t, 1, set.Size())
}

func TestSet_Add(t *testing.T) {
	set := New()
	set.Add("foo")
	require.True(t, set.Contains("foo"))
	require.False(t, set.Contains("bar"))
}

func TestSet_Delete(t *testing.T) {
	set := New()

	set.Add("foo")
	set.Add("bar")
	set.Remove("foo")

	require.False(t, set.contains("foo"))
	require.True(t, set.contains("bar"))
}

func TestSet_Items(t *testing.T) {
	set := New()
	set.Add("foo")
	set.Add("bar")

	require.ElementsMatch(t, []string{"foo", "bar"}, set.Values())
}

func benchmarkHas(b *testing.B, set *Set, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.contains(n)
		}
	}
}

func benchmarkAdd(b *testing.B, set *Set, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, set *Set, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkHashSetHas100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkHashSetHas1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkHashSetHas10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkHashSetHas100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkHashSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}

	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}

	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}

	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}

	b.StartTimer()
	benchmarkRemove(b, set, size)
}
