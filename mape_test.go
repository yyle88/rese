package rese_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/rese"
)

func TestM0(t *testing.T) {
	run := func() error {
		return nil
	}

	rese.M0(run())
}

func TestM1(t *testing.T) {
	run := func() (map[int]string, error) {
		return map[int]string{0: "a"}, nil
	}

	m1 := rese.M1(run())
	require.Equal(t, map[int]string{0: "a"}, m1)
}

func TestM2(t *testing.T) {
	run := func() (map[int]string, map[string]int, error) {
		return map[int]string{3: "a"}, map[string]int{"x": 1}, nil
	}

	m1, m2 := rese.M2(run())
	require.Equal(t, map[int]string{3: "a"}, m1)
	require.Equal(t, map[string]int{"x": 1}, m2)
}

func TestM3(t *testing.T) {
	run := func() (map[int8]bool, map[float32]string, map[string]byte, error) {
		return map[int8]bool{1: true}, map[float32]string{1.5: "float"}, map[string]byte{"key": 'A'}, nil
	}

	m1, m2, m3 := rese.M3(run())
	require.Equal(t, map[int8]bool{1: true}, m1)
	require.Equal(t, map[float32]string{1.5: "float"}, m2)
	require.Equal(t, map[string]byte{"key": 'A'}, m3)
}

func TestM4(t *testing.T) {
	run := func() (map[int64]float64, map[bool]int, map[rune]string, map[uint8]bool, error) {
		return map[int64]float64{100: 3.14}, map[bool]int{true: 1}, map[rune]string{'x': "rune"}, map[uint8]bool{255: false}, nil
	}

	m1, m2, m3, m4 := rese.M4(run())
	require.Equal(t, map[int64]float64{100: 3.14}, m1)
	require.Equal(t, map[bool]int{true: 1}, m2)
	require.Equal(t, map[rune]string{'x': "rune"}, m3)
	require.Equal(t, map[uint8]bool{255: false}, m4)
}

func TestM5(t *testing.T) {
	run := func() (map[float64]int32, map[int16]complex64, map[string]uint, map[bool]byte, map[uint16]float32, error) {
		return map[float64]int32{2.718: 42}, map[int16]complex64{10: 1 + 2i}, map[string]uint{"num": 100}, map[bool]byte{false: '0'}, map[uint16]float32{1: 0.5}, nil
	}

	m1, m2, m3, m4, m5 := rese.M5(run())
	require.Equal(t, map[float64]int32{2.718: 42}, m1)
	require.Equal(t, map[int16]complex64{10: 1 + 2i}, m2)
	require.Equal(t, map[string]uint{"num": 100}, m3)
	require.Equal(t, map[bool]byte{false: '0'}, m4)
	require.Equal(t, map[uint16]float32{1: 0.5}, m5)
}

func TestM6(t *testing.T) {
	run := func() (map[uint32]string, map[int]float64, map[complex128]bool, map[string]int8, map[float32]uint64, map[bool]rune, error) {
		return map[uint32]string{1000: "test"}, map[int]float64{5: 2.5}, map[complex128]bool{1 + 1i: true}, map[string]int8{"x": -1}, map[float32]uint64{0.1: 999}, map[bool]rune{true: 'Z'}, nil
	}

	m1, m2, m3, m4, m5, m6 := rese.M6(run())
	require.Equal(t, map[uint32]string{1000: "test"}, m1)
	require.Equal(t, map[int]float64{5: 2.5}, m2)
	require.Equal(t, map[complex128]bool{1 + 1i: true}, m3)
	require.Equal(t, map[string]int8{"x": -1}, m4)
	require.Equal(t, map[float32]uint64{0.1: 999}, m5)
	require.Equal(t, map[bool]rune{true: 'Z'}, m6)
}

func TestM7(t *testing.T) {
	run := func() (map[int32]uint32, map[uint64]float64, map[string]complex128, map[bool]int16, map[rune]byte, map[float64]string, map[int8]bool, error) {
		return map[int32]uint32{-1: 10}, map[uint64]float64{100: 1.23}, map[string]complex128{"c": 3 + 4i}, map[bool]int16{true: 8}, map[rune]byte{'a': 65}, map[float64]string{0.0: "zero"}, map[int8]bool{127: false}, nil
	}

	m1, m2, m3, m4, m5, m6, m7 := rese.M7(run())
	require.Equal(t, map[int32]uint32{-1: 10}, m1)
	require.Equal(t, map[uint64]float64{100: 1.23}, m2)
	require.Equal(t, map[string]complex128{"c": 3 + 4i}, m3)
	require.Equal(t, map[bool]int16{true: 8}, m4)
	require.Equal(t, map[rune]byte{'a': 65}, m5)
	require.Equal(t, map[float64]string{0.0: "zero"}, m6)
	require.Equal(t, map[int8]bool{127: false}, m7)
}

func TestM8(t *testing.T) {
	run := func() (map[uint16]int64, map[float32]uint8, map[string]float64, map[bool]complex64, map[int]rune, map[complex128]string, map[uint32]bool, map[int16]int, error) {
		return map[uint16]int64{1: 1000}, map[float32]uint8{2.2: 255}, map[string]float64{"pi": 3.14}, map[bool]complex64{false: 2 + 3i}, map[int]rune{0: 'x'}, map[complex128]string{5 + 6i: "complex"}, map[uint32]bool{10: true}, map[int16]int{-5: 42}, nil
	}

	m1, m2, m3, m4, m5, m6, m7, m8 := rese.M8(run())
	require.Equal(t, map[uint16]int64{1: 1000}, m1)
	require.Equal(t, map[float32]uint8{2.2: 255}, m2)
	require.Equal(t, map[string]float64{"pi": 3.14}, m3)
	require.Equal(t, map[bool]complex64{false: 2 + 3i}, m4)
	require.Equal(t, map[int]rune{0: 'x'}, m5)
	require.Equal(t, map[complex128]string{5 + 6i: "complex"}, m6)
	require.Equal(t, map[uint32]bool{10: true}, m7)
	require.Equal(t, map[int16]int{-5: 42}, m8)
}

func TestM9(t *testing.T) {
	run := func() (map[int8]uint16, map[float64]int32, map[string]bool, map[uint]float32, map[complex64]string, map[bool]int64, map[rune]uint8, map[int16]complex128, map[uint64]string, error) {
		return map[int8]uint16{1: 100}, map[float64]int32{1.1: 50}, map[string]bool{"yes": true}, map[uint]float32{5: 0.5}, map[complex64]string{1 + 1i: "c64"}, map[bool]int64{true: 999}, map[rune]uint8{'r': 82}, map[int16]complex128{10: 2 + 3i}, map[uint64]string{1000: "end"}, nil
	}

	m1, m2, m3, m4, m5, m6, m7, m8, m9 := rese.M9(run())
	require.Equal(t, map[int8]uint16{1: 100}, m1)
	require.Equal(t, map[float64]int32{1.1: 50}, m2)
	require.Equal(t, map[string]bool{"yes": true}, m3)
	require.Equal(t, map[uint]float32{5: 0.5}, m4)
	require.Equal(t, map[complex64]string{1 + 1i: "c64"}, m5)
	require.Equal(t, map[bool]int64{true: 999}, m6)
	require.Equal(t, map[rune]uint8{'r': 82}, m7)
	require.Equal(t, map[int16]complex128{10: 2 + 3i}, m8)
	require.Equal(t, map[uint64]string{1000: "end"}, m9)
}
