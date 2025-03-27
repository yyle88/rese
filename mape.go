package rese

import (
	"github.com/yyle88/done"
	"github.com/yyle88/must/mustmap"
)

func M0(err error) {
	done.Done(err)
}

func M1[K1 comparable, V1 any](m1 map[K1]V1, err error) map[K1]V1 {
	done.Done(err)
	mustmap.Have(m1)
	return m1
}

func M2[K1, K2 comparable, V1, V2 any](m1 map[K1]V1, m2 map[K2]V2, err error) (map[K1]V1, map[K2]V2) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	return m1, m2
}

func M3[K1, K2, K3 comparable, V1, V2, V3 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, err error) (map[K1]V1, map[K2]V2, map[K3]V3) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	return m1, m2, m3
}

func M4[K1, K2, K3, K4 comparable, V1, V2, V3, V4 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	return m1, m2, m3, m4
}

func M5[K1, K2, K3, K4, K5 comparable, V1, V2, V3, V4, V5 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	return m1, m2, m3, m4, m5
}

func M6[K1, K2, K3, K4, K5, K6 comparable, V1, V2, V3, V4, V5, V6 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, m6 map[K6]V6, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5, map[K6]V6) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	mustmap.Have(m6)
	return m1, m2, m3, m4, m5, m6
}

func M7[K1, K2, K3, K4, K5, K6, K7 comparable, V1, V2, V3, V4, V5, V6, V7 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, m6 map[K6]V6, m7 map[K7]V7, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5, map[K6]V6, map[K7]V7) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	mustmap.Have(m6)
	mustmap.Have(m7)
	return m1, m2, m3, m4, m5, m6, m7
}

func M8[K1, K2, K3, K4, K5, K6, K7, K8 comparable, V1, V2, V3, V4, V5, V6, V7, V8 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, m6 map[K6]V6, m7 map[K7]V7, m8 map[K8]V8, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5, map[K6]V6, map[K7]V7, map[K8]V8) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	mustmap.Have(m6)
	mustmap.Have(m7)
	mustmap.Have(m8)
	return m1, m2, m3, m4, m5, m6, m7, m8
}

func M9[K1, K2, K3, K4, K5, K6, K7, K8, K9 comparable, V1, V2, V3, V4, V5, V6, V7, V8, V9 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, m6 map[K6]V6, m7 map[K7]V7, m8 map[K8]V8, m9 map[K9]V9, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5, map[K6]V6, map[K7]V7, map[K8]V8, map[K9]V9) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	mustmap.Have(m6)
	mustmap.Have(m7)
	mustmap.Have(m8)
	mustmap.Have(m9)
	return m1, m2, m3, m4, m5, m6, m7, m8, m9
}
