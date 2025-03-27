package rese_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/rese"
)

func TestF0(t *testing.T) {
	run := func() error {
		return nil
	}

	rese.F0(run)
}

func TestF1(t *testing.T) {
	run := func() (int, error) {
		return 100, nil
	}

	v1 := rese.F1(run)
	require.Equal(t, 100, v1)
}

func TestF2(t *testing.T) {
	run := func() (string, map[string]int, error) {
		return "abc", map[string]int{"x": 1}, nil
	}

	v1, v2 := rese.F2(run)
	require.Equal(t, "abc", v1)
	require.Equal(t, map[string]int{"x": 1}, v2)
}

func TestF3(t *testing.T) {
	run := func() (bool, float64, []int, error) {
		return true, 3.14, []int{1, 2, 3}, nil
	}

	v1, v2, v3 := rese.F3(run)
	require.Equal(t, true, v1)
	require.Equal(t, 3.14, v2)
	require.Equal(t, []int{1, 2, 3}, v3)
}

func TestF4(t *testing.T) {
	run := func() (int8, uint16, complex64, string, error) {
		return 42, 1000, 1 + 2i, "test", nil
	}

	v1, v2, v3, v4 := rese.F4(run)
	require.Equal(t, int8(42), v1)
	require.Equal(t, uint16(1000), v2)
	require.Equal(t, complex64(1+2i), v3)
	require.Equal(t, "test", v4)
}

func TestF5(t *testing.T) {
	run := func() (float32, byte, map[int]bool, int64, rune, error) {
		return 0.5, 'A', map[int]bool{1: true}, 999, 'x', nil
	}

	v1, v2, v3, v4, v5 := rese.F5(run)
	require.Equal(t, float32(0.5), v1)
	require.Equal(t, byte('A'), v2)
	require.Equal(t, map[int]bool{1: true}, v3)
	require.Equal(t, int64(999), v4)
	require.Equal(t, rune('x'), v5)
}

func TestF6(t *testing.T) {
	run := func() (uint32, complex128, []string, bool, float64, int16, error) {
		return 500, 3 + 4i, []string{"a", "b"}, false, 2.718, -10, nil
	}

	v1, v2, v3, v4, v5, v6 := rese.F6(run)
	require.Equal(t, uint32(500), v1)
	require.Equal(t, complex128(3+4i), v2)
	require.Equal(t, []string{"a", "b"}, v3)
	require.Equal(t, false, v4)
	require.Equal(t, 2.718, v5)
	require.Equal(t, int16(-10), v6)
}

func TestF7(t *testing.T) {
	run := func() (string, uint8, map[float32]int, int32, bool, complex64, []byte, error) {
		return "xyz", 255, map[float32]int{1.1: 11}, -100, true, 5 + 6i, []byte{1, 2}, nil
	}

	v1, v2, v3, v4, v5, v6, v7 := rese.F7(run)
	require.Equal(t, "xyz", v1)
	require.Equal(t, uint8(255), v2)
	require.Equal(t, map[float32]int{1.1: 11}, v3)
	require.Equal(t, int32(-100), v4)
	require.Equal(t, true, v5)
	require.Equal(t, complex64(5+6i), v6)
	require.Equal(t, []byte{1, 2}, v7)
}

func TestF8(t *testing.T) {
	run := func() (int, float32, uint64, string, rune, map[bool]string, complex128, []int8, error) {
		return 42, 0.123, 10000, "hello", 'Z', map[bool]string{true: "yes"}, 7 + 8i, []int8{1, -1}, nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := rese.F8(run)
	require.Equal(t, 42, v1)
	require.Equal(t, float32(0.123), v2)
	require.Equal(t, uint64(10000), v3)
	require.Equal(t, "hello", v4)
	require.Equal(t, rune('Z'), v5)
	require.Equal(t, map[bool]string{true: "yes"}, v6)
	require.Equal(t, complex128(7+8i), v7)
	require.Equal(t, []int8{1, -1}, v8)
}

func TestF9(t *testing.T) {
	run := func() (bool, int16, float64, uint32, string, complex64, []float32, map[int]string, byte, error) {
		return false, 32767, 1.414, 1234, "end", 2 + 3i, []float32{0.1, 0.2}, map[int]string{1: "one"}, 'B', nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := rese.F9(run)
	require.Equal(t, false, v1)
	require.Equal(t, int16(32767), v2)
	require.Equal(t, 1.414, v3)
	require.Equal(t, uint32(1234), v4)
	require.Equal(t, "end", v5)
	require.Equal(t, complex64(2+3i), v6)
	require.Equal(t, []float32{0.1, 0.2}, v7)
	require.Equal(t, map[int]string{1: "one"}, v8)
	require.Equal(t, byte('B'), v9)
}

func TestR0(t *testing.T) {
	run := func() error {
		return nil
	}

	rese.R0(run)
}

func TestR1(t *testing.T) {
	run := func() (*string, error) {
		v1 := "hello"
		return &v1, nil
	}

	p1 := rese.R1(run)
	require.Equal(t, "hello", *p1)
}

func TestR2(t *testing.T) {
	run := func() (*int, *float64, error) {
		v1 := 42
		v2 := 3.14
		return &v1, &v2, nil
	}

	v1, v2 := rese.R2(run)
	require.Equal(t, 42, *v1)
	require.Equal(t, 3.14, *v2)
}

func TestR3(t *testing.T) {
	run := func() (*bool, *uint16, *[]string, error) {
		v1 := true
		v2 := uint16(1000)
		v3 := []string{"a", "b"}
		return &v1, &v2, &v3, nil
	}

	v1, v2, v3 := rese.R3(run)
	require.Equal(t, true, *v1)
	require.Equal(t, uint16(1000), *v2)
	require.Equal(t, []string{"a", "b"}, *v3)
}

func TestR4(t *testing.T) {
	run := func() (*int8, *float32, *complex64, *map[string]int, error) {
		v1 := int8(-10)
		v2 := float32(0.5)
		v3 := complex64(1 + 2i)
		v4 := map[string]int{"x": 1}
		return &v1, &v2, &v3, &v4, nil
	}

	v1, v2, v3, v4 := rese.R4(run)
	require.Equal(t, int8(-10), *v1)
	require.Equal(t, float32(0.5), *v2)
	require.Equal(t, complex64(1+2i), *v3)
	require.Equal(t, map[string]int{"x": 1}, *v4)
}

func TestR5(t *testing.T) {
	run := func() (*uint32, *byte, *int64, *rune, *[]int, error) {
		v1 := uint32(500)
		v2 := byte('A')
		v3 := int64(999)
		v4 := rune('x')
		v5 := []int{1, 2, 3}
		return &v1, &v2, &v3, &v4, &v5, nil
	}

	v1, v2, v3, v4, v5 := rese.R5(run)
	require.Equal(t, uint32(500), *v1)
	require.Equal(t, byte('A'), *v2)
	require.Equal(t, int64(999), *v3)
	require.Equal(t, rune('x'), *v4)
	require.Equal(t, []int{1, 2, 3}, *v5)
}

func TestR6(t *testing.T) {
	run := func() (*float64, *int16, *complex128, *bool, *string, *map[int]bool, error) {
		v1 := 2.718
		v2 := int16(-100)
		v3 := complex128(3 + 4i)
		v4 := false
		v5 := "test"
		v6 := map[int]bool{1: true}
		return &v1, &v2, &v3, &v4, &v5, &v6, nil
	}

	v1, v2, v3, v4, v5, v6 := rese.R6(run)
	require.Equal(t, 2.718, *v1)
	require.Equal(t, int16(-100), *v2)
	require.Equal(t, complex128(3+4i), *v3)
	require.Equal(t, false, *v4)
	require.Equal(t, "test", *v5)
	require.Equal(t, map[int]bool{1: true}, *v6)
}

func TestR7(t *testing.T) {
	run := func() (*uint8, *float32, *int32, *rune, *complex64, *[]byte, *map[bool]string, error) {
		v1 := uint8(255)
		v2 := float32(1.23)
		v3 := int32(-50)
		v4 := rune('Z')
		v5 := complex64(5 + 6i)
		v6 := []byte{1, 2}
		v7 := map[bool]string{true: "yes"}
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, nil
	}

	v1, v2, v3, v4, v5, v6, v7 := rese.R7(run)
	require.Equal(t, uint8(255), *v1)
	require.Equal(t, float32(1.23), *v2)
	require.Equal(t, int32(-50), *v3)
	require.Equal(t, rune('Z'), *v4)
	require.Equal(t, complex64(5+6i), *v5)
	require.Equal(t, []byte{1, 2}, *v6)
	require.Equal(t, map[bool]string{true: "yes"}, *v7)
}

func TestR8(t *testing.T) {
	run := func() (*int, *uint64, *float64, *string, *complex128, *[]float32, *map[string]uint, *bool, error) {
		v1 := 100
		v2 := uint64(10000)
		v3 := 1.414
		v4 := "world"
		v5 := complex128(7 + 8i)
		v6 := []float32{0.1, 0.2}
		v7 := map[string]uint{"n": 42}
		v8 := true
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := rese.R8(run)
	require.Equal(t, 100, *v1)
	require.Equal(t, uint64(10000), *v2)
	require.Equal(t, 1.414, *v3)
	require.Equal(t, "world", *v4)
	require.Equal(t, complex128(7+8i), *v5)
	require.Equal(t, []float32{0.1, 0.2}, *v6)
	require.Equal(t, map[string]uint{"n": 42}, *v7)
	require.Equal(t, true, *v8)
}

func TestR9(t *testing.T) {
	run := func() (*int16, *uint32, *float32, *string, *complex64, *[]int8, *map[int]float64, *bool, *rune, error) {
		v1 := int16(32767)
		v2 := uint32(1234)
		v3 := float32(0.618)
		v4 := "end"
		v5 := complex64(2 + 3i)
		v6 := []int8{1, -1}
		v7 := map[int]float64{1: 1.1}
		v8 := false
		v9 := rune('B')
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, &v9, nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := rese.R9(run)
	require.Equal(t, int16(32767), *v1)
	require.Equal(t, uint32(1234), *v2)
	require.Equal(t, float32(0.618), *v3)
	require.Equal(t, "end", *v4)
	require.Equal(t, complex64(2+3i), *v5)
	require.Equal(t, []int8{1, -1}, *v6)
	require.Equal(t, map[int]float64{1: 1.1}, *v7)
	require.Equal(t, false, *v8)
	require.Equal(t, rune('B'), *v9)
}
