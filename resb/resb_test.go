package resb_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/rese/resb"
)

func TestV0(t *testing.T) {
	run := func() bool {
		return true
	}

	resb.V0(run())
}

func TestV1(t *testing.T) {
	run := func() (string, bool) {
		return "a", true
	}

	v1 := resb.V1(run())
	require.Equal(t, "a", v1)
}

func TestV2(t *testing.T) {
	run := func() (string, uint64, bool) {
		return "a", 2, true
	}

	v1, v2 := resb.V2(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
}

func TestV3(t *testing.T) {
	run := func() (string, int, float64, bool) {
		return "a", 42, 3.14, true
	}

	v1, v2, v3 := resb.V3(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
}

func TestV4(t *testing.T) {
	run := func() (string, int, float64, bool, bool) {
		return "a", 42, 3.14, true, true
	}

	v1, v2, v3, v4 := resb.V4(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
}

func TestV5(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, bool) {
		return "a", 42, 3.14, true, 100, true
	}

	v1, v2, v3, v4, v5 := resb.V5(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
}

func TestV6(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, int64, bool) {
		return "a", 42, 3.14, true, 100, 1000, true
	}

	v1, v2, v3, v4, v5, v6 := resb.V6(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
	require.Equal(t, int64(1000), v6)
}

func TestV7(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, int64, rune, bool) {
		return "a", 42, 3.14, true, 100, 1000, 'x', true
	}

	v1, v2, v3, v4, v5, v6, v7 := resb.V7(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
	require.Equal(t, int64(1000), v6)
	require.Equal(t, rune('x'), v7)
}

func TestV8(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, int64, rune, byte, bool) {
		return "a", 42, 3.14, true, 100, 1000, 'x', 'y', true
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := resb.V8(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
	require.Equal(t, int64(1000), v6)
	require.Equal(t, rune('x'), v7)
	require.Equal(t, byte('y'), v8)
}

func TestV9(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, int64, rune, byte, complex64, bool) {
		return "a", 42, 3.14, true, 100, 1000, 'x', 'y', 1 + 2i, true
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := resb.V9(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
	require.Equal(t, int64(1000), v6)
	require.Equal(t, rune('x'), v7)
	require.Equal(t, byte('y'), v8)
	require.Equal(t, complex64(1+2i), v9)
}

func TestP0(t *testing.T) {
	run := func() bool {
		return true
	}

	resb.P0(run())
}

func TestP1(t *testing.T) {
	run := func() (*string, bool) {
		v1 := "hello"
		return &v1, true
	}

	p1 := resb.P1(run())
	require.Equal(t, "hello", *p1)
}

func TestP2(t *testing.T) {
	run := func() (*int, *float64, bool) {
		v1 := 42
		v2 := 3.14
		return &v1, &v2, true
	}

	v1, v2 := resb.P2(run())
	require.Equal(t, 42, *v1)
	require.Equal(t, 3.14, *v2)
}

func TestP3(t *testing.T) {
	run := func() (*string, *int, *float64, bool) {
		v1 := "a"
		v2 := 42
		v3 := 3.14
		return &v1, &v2, &v3, true
	}

	v1, v2, v3 := resb.P3(run())
	require.Equal(t, "a", *v1)
	require.Equal(t, 42, *v2)
	require.Equal(t, 3.14, *v3)
}

func TestP4(t *testing.T) {
	run := func() (*string, *int, *float64, *bool, bool) {
		v1 := "a"
		v2 := 42
		v3 := 3.14
		v4 := true
		return &v1, &v2, &v3, &v4, true
	}

	v1, v2, v3, v4 := resb.P4(run())
	require.Equal(t, "a", *v1)
	require.Equal(t, 42, *v2)
	require.Equal(t, 3.14, *v3)
	require.Equal(t, true, *v4)
}

func TestP5(t *testing.T) {
	run := func() (*string, *int, *float64, *bool, *uint32, bool) {
		v1 := "a"
		v2 := 42
		v3 := 3.14
		v4 := true
		v5 := uint32(100)
		return &v1, &v2, &v3, &v4, &v5, true
	}

	v1, v2, v3, v4, v5 := resb.P5(run())
	require.Equal(t, "a", *v1)
	require.Equal(t, 42, *v2)
	require.Equal(t, 3.14, *v3)
	require.Equal(t, true, *v4)
	require.Equal(t, uint32(100), *v5)
}

func TestP6(t *testing.T) {
	run := func() (*string, *int, *float64, *bool, *uint32, *int64, bool) {
		v1 := "a"
		v2 := 42
		v3 := 3.14
		v4 := true
		v5 := uint32(100)
		v6 := int64(1000)
		return &v1, &v2, &v3, &v4, &v5, &v6, true
	}

	v1, v2, v3, v4, v5, v6 := resb.P6(run())
	require.Equal(t, "a", *v1)
	require.Equal(t, 42, *v2)
	require.Equal(t, 3.14, *v3)
	require.Equal(t, true, *v4)
	require.Equal(t, uint32(100), *v5)
	require.Equal(t, int64(1000), *v6)
}

func TestP7(t *testing.T) {
	run := func() (*string, *int, *float64, *bool, *uint32, *int64, *rune, bool) {
		v1 := "a"
		v2 := 42
		v3 := 3.14
		v4 := true
		v5 := uint32(100)
		v6 := int64(1000)
		v7 := rune('x')
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, true
	}

	v1, v2, v3, v4, v5, v6, v7 := resb.P7(run())
	require.Equal(t, "a", *v1)
	require.Equal(t, 42, *v2)
	require.Equal(t, 3.14, *v3)
	require.Equal(t, true, *v4)
	require.Equal(t, uint32(100), *v5)
	require.Equal(t, int64(1000), *v6)
	require.Equal(t, rune('x'), *v7)
}

func TestP8(t *testing.T) {
	run := func() (*string, *int, *float64, *bool, *uint32, *int64, *rune, *byte, bool) {
		v1 := "a"
		v2 := 42
		v3 := 3.14
		v4 := true
		v5 := uint32(100)
		v6 := int64(1000)
		v7 := rune('x')
		v8 := byte('y')
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, true
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := resb.P8(run())
	require.Equal(t, "a", *v1)
	require.Equal(t, 42, *v2)
	require.Equal(t, 3.14, *v3)
	require.Equal(t, true, *v4)
	require.Equal(t, uint32(100), *v5)
	require.Equal(t, int64(1000), *v6)
	require.Equal(t, rune('x'), *v7)
	require.Equal(t, byte('y'), *v8)
}

func TestP9(t *testing.T) {
	run := func() (*string, *int, *float64, *bool, *uint32, *int64, *rune, *byte, *complex64, bool) {
		v1 := "a"
		v2 := 42
		v3 := 3.14
		v4 := true
		v5 := uint32(100)
		v6 := int64(1000)
		v7 := rune('x')
		v8 := byte('y')
		v9 := complex64(1 + 2i)
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, &v9, true
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := resb.P9(run())
	require.Equal(t, "a", *v1)
	require.Equal(t, 42, *v2)
	require.Equal(t, 3.14, *v3)
	require.Equal(t, true, *v4)
	require.Equal(t, uint32(100), *v5)
	require.Equal(t, int64(1000), *v6)
	require.Equal(t, rune('x'), *v7)
	require.Equal(t, byte('y'), *v8)
	require.Equal(t, complex64(1+2i), *v9)
}

func TestC0(t *testing.T) {
	run := func() bool {
		return true
	}

	resb.C0(run())
}

func TestC1(t *testing.T) {
	run := func() (string, bool) {
		return "hello", true
	}

	v1 := resb.C1(run())
	require.Equal(t, "hello", v1)
}

func TestC2(t *testing.T) {
	run := func() (int, float64, bool) {
		return 42, 3.14, true
	}

	v1, v2 := resb.C2(run())
	require.Equal(t, 42, v1)
	require.Equal(t, 3.14, v2)
}

func TestC3(t *testing.T) {
	run := func() (string, int, float64, bool) {
		return "a", 42, 3.14, true
	}

	v1, v2, v3 := resb.C3(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
}

func TestC4(t *testing.T) {
	run := func() (string, int, float64, bool, bool) {
		return "a", 42, 3.14, true, true
	}

	v1, v2, v3, v4 := resb.C4(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
}

func TestC5(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, bool) {
		return "a", 42, 3.14, true, 100, true
	}

	v1, v2, v3, v4, v5 := resb.C5(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
}

func TestC6(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, int64, bool) {
		return "a", 42, 3.14, true, 100, 1000, true
	}

	v1, v2, v3, v4, v5, v6 := resb.C6(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
	require.Equal(t, int64(1000), v6)
}

func TestC7(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, int64, rune, bool) {
		return "a", 42, 3.14, true, 100, 1000, 'x', true
	}

	v1, v2, v3, v4, v5, v6, v7 := resb.C7(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
	require.Equal(t, int64(1000), v6)
	require.Equal(t, rune('x'), v7)
}

func TestC8(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, int64, rune, byte, bool) {
		return "a", 42, 3.14, true, 100, 1000, 'x', 'y', true
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := resb.C8(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
	require.Equal(t, int64(1000), v6)
	require.Equal(t, rune('x'), v7)
	require.Equal(t, byte('y'), v8)
}

func TestC9(t *testing.T) {
	run := func() (string, int, float64, bool, uint32, int64, rune, byte, uint8, bool) {
		return "a", 42, 3.14, true, 100, 1000, 'x', 'y', 255, true
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := resb.C9(run())
	require.Equal(t, "a", v1)
	require.Equal(t, 42, v2)
	require.Equal(t, 3.14, v3)
	require.Equal(t, true, v4)
	require.Equal(t, uint32(100), v5)
	require.Equal(t, int64(1000), v6)
	require.Equal(t, rune('x'), v7)
	require.Equal(t, byte('y'), v8)
	require.Equal(t, uint8(255), v9)
}
