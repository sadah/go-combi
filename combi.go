// Package combi implements some combinatoric functions
package combi

const (
	errNegativeInput = "combi: Don't use negative"
	errBadSize       = "combi: n < k"
)

// P calculates P(n, k)
// P(n, k) = n! / k!
func P(n, k int) int {
	if n < 0 || k < 0 {
		panic(errNegativeInput)
	}
	if n < k {
		panic(errBadSize)
	}
	p := 1
	for i := 1; i <= k; i++ {
		p *= n
		n--
	}
	return p
}

// C returns C(n,k)
// C(n,k) = n!/((n-k)!k!)
func C(n, k int) int {
	return P(n, k) / P(k, k)
}

// Factorical returns n!
func Factorical(n int) int {
	return P(n, n)
}
