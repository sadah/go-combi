package main

import (
	"fmt"

	"github.com/sadah/go-combi"
)

func main() {
	func() {
		psin := combi.PowerSetIndex(3)
		fmt.Println(psin)
		// => [[] [0] [1] [0 1] [2] [0 2] [1 2] [0 1 2]]
	}()

	func() {
		ints := []int{1, 2, 3}
		psi := combi.PowerSetInts(ints)
		fmt.Println(psi)
		// => [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
	}()

	func() {
		strs := []string{"a", "b", "c"}
		pss := combi.PowerSetStrs(strs)
		fmt.Println(pss)
		// => [[] [a] [b] [a b] [c] [a c] [b c] [a b c]]
	}()

	func() {
		cin := combi.CombinationIndex(3, 2)
		fmt.Println(cin)
		// => [[0 1] [0 2] [1 2]]
	}()

	func() {
		ints := []int{1, 2, 3}
		ci := combi.CombinationInts(ints, 2)
		fmt.Println(ci)
		// => [[1 2] [1 3] [2 3]]
	}()

	func() {
		strs := []string{"a", "b", "c"}
		cs := combi.CombinationStrs(strs, 2)
		fmt.Println(cs)
		// => [[a b] [a c] [b c]]
	}()

	func() {
		c := combi.C(5, 2)
		fmt.Println(c)
		// => 10
	}()

	func() {
		p := combi.P(5, 2)
		fmt.Println(p)
		// => 20
	}()

	func() {
		factorial := combi.Factorial(5)
		fmt.Println(factorial)
		// => 120
	}()
}
