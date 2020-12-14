package utils

import "regexp"

func ModuloInverse(a, m int) int {
	if m == 1 {
		return 0
	}
	m0, x0, x1 := m, 0, 1

	for a > 1 {
		q := a / m
		t := m

		m = a % m
		a = t

		t = x0
		x0 = x1 - q*x0
		x1 = t
	}
	if x1 < 0 {
		x1 = x1 + m0
	}

	return x1
}

func ParseIds(ids string) []string {
	reg := regexp.MustCompile(`,`)
	split := reg.Split(ids, -1)
	ans := make([]string, 0)
	for _, id := range split {
		ans = append(ans, id)
	}
	return ans
}
