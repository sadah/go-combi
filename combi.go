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
	if k > n-k {
		k = n - k
	}
	return P(n, k) / P(k, k)
}

// Factorial returns n!
func Factorial(n int) int {
	return P(n, n)
}

// PowerSetIndex returns indexes of power sets
func PowerSetIndex(n int) [][]int {
	if n < 0 {
		panic(errNegativeInput)
	}
	if n == 0 {
		return [][]int{{}}
	}

	size := 1 << n
	ret := make([][]int, size)

	for bit := 0; bit < size; bit++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) != 0 {
				subset = append(subset, i)
			}
		}
		ret[bit] = subset
	}
	return ret
}

// PowerSetInts returns power sets
func PowerSetInts(ints []int) [][]int {
	n := len(ints)
	if n == 0 {
		return [][]int{{}}
	}

	size := 1 << n
	ret := make([][]int, size)

	for bit := 0; bit < size; bit++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) != 0 {
				subset = append(subset, ints[i])
			}
		}
		ret[bit] = subset
	}
	return ret
}

// PowerSetStrs returns power sets
func PowerSetStrs(strs []string) [][]string {
	n := len(strs)
	if n == 0 {
		return [][]string{{}}
	}

	size := 1 << n
	ret := make([][]string, size)

	for bit := 0; bit < size; bit++ {
		subset := []string{}
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) != 0 {
				subset = append(subset, strs[i])
			}
		}
		ret[bit] = subset
	}
	return ret
}
