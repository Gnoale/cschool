package utils

func Pow(n, m int) int {
	if m == 0 {
		return 1
	}
	b := n
	for i := 1; i < m; i++ {
		n *= b
	}
	return n
}

// need to check if n > m...
func Gcd(n, m int) int {
	for m != 0 {
		t := m
		m = n % m
		n = t
	}
	return n
}

func GcdRec(n, m int) int {
	if m == 0 {
		return n
	}
	return GcdRec(m, n%m)
}
