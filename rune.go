package rese

func F0(run func() error) {
	V0(run())
}

func F1[T1 any](run func() (T1, error)) T1 {
	return V1(run())
}

func F2[T1, T2 any](run func() (T1, T2, error)) (T1, T2) {
	return V2(run())
}

func F3[T1, T2, T3 any](run func() (T1, T2, T3, error)) (T1, T2, T3) {
	return V3(run())
}

func F4[T1, T2, T3, T4 any](run func() (T1, T2, T3, T4, error)) (T1, T2, T3, T4) {
	return V4(run())
}

func F5[T1, T2, T3, T4, T5 any](run func() (T1, T2, T3, T4, T5, error)) (T1, T2, T3, T4, T5) {
	return V5(run())
}

func F6[T1, T2, T3, T4, T5, T6 any](run func() (T1, T2, T3, T4, T5, T6, error)) (T1, T2, T3, T4, T5, T6) {
	return V6(run())
}

func F7[T1, T2, T3, T4, T5, T6, T7 any](run func() (T1, T2, T3, T4, T5, T6, T7, error)) (T1, T2, T3, T4, T5, T6, T7) {
	return V7(run())
}

func F8[T1, T2, T3, T4, T5, T6, T7, T8 any](run func() (T1, T2, T3, T4, T5, T6, T7, T8, error)) (T1, T2, T3, T4, T5, T6, T7, T8) {
	return V8(run())
}

func F9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](run func() (T1, T2, T3, T4, T5, T6, T7, T8, T9, error)) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return V9(run())
}

func R0(run func() error) {
	P0(run())
}

func R1[T1 any](run func() (*T1, error)) *T1 {
	return P1(run())
}

func R2[T1, T2 any](run func() (*T1, *T2, error)) (*T1, *T2) {
	return P2(run())
}

func R3[T1, T2, T3 any](run func() (*T1, *T2, *T3, error)) (*T1, *T2, *T3) {
	return P3(run())
}

func R4[T1, T2, T3, T4 any](run func() (*T1, *T2, *T3, *T4, error)) (*T1, *T2, *T3, *T4) {
	return P4(run())
}

func R5[T1, T2, T3, T4, T5 any](run func() (*T1, *T2, *T3, *T4, *T5, error)) (*T1, *T2, *T3, *T4, *T5) {
	return P5(run())
}

func R6[T1, T2, T3, T4, T5, T6 any](run func() (*T1, *T2, *T3, *T4, *T5, *T6, error)) (*T1, *T2, *T3, *T4, *T5, *T6) {
	return P6(run())
}

func R7[T1, T2, T3, T4, T5, T6, T7 any](run func() (*T1, *T2, *T3, *T4, *T5, *T6, *T7, error)) (*T1, *T2, *T3, *T4, *T5, *T6, *T7) {
	return P7(run())
}

func R8[T1, T2, T3, T4, T5, T6, T7, T8 any](run func() (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8, error)) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8) {
	return P8(run())
}

func R9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](run func() (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8, *T9, error)) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8, *T9) {
	return P9(run())
}
