# go-combi

[![Go](https://github.com/sadah/go-combi/workflows/Go/badge.svg)](https://github.com/sadah/go-combi/actions)
[![codecov](https://codecov.io/gh/sadah/go-combi/branch/master/graph/badge.svg)](https://codecov.io/gh/sadah/go-combi)
[![Go Report Card](https://goreportcard.com/badge/github.com/sadah/go-combi)](https://goreportcard.com/report/github.com/sadah/go-combi)

combi implements some combinatoric functions

# SYNOPSIS

You can see a sample code in cmd directory.

```go
package main

import (
	"fmt"

	"github.com/sadah/go-combi"
)

func main() {
    ints := []int{1, 2, 3}
    psi := combi.PowerSetInts(ints)
    fmt.Println(psi)
    // => [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}
```

# DESCRIPTION

TBD

# API

## Power Set

### Power Set Index

```go
psi := combi.PowerSetIndex(5)
fmt.Println(psi)
// => [[] [0] [1] [0 1] [2] [0 2] [1 2] [0 1 2]]
```

### Power Set Ins

```go
ints := []int{1, 2, 3}
psi := combi.PowerSetInts(ints)
fmt.Println(psi)
// => [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
```

### Power Set Strs

```go
strs := []string{"a", "b", "c"}
pss := combi.PowerSetStrs(strs)
fmt.Println(pss)
// => [[] [a] [b] [a b] [c] [a c] [b c] [a b c]]
```

## Combination

### Combination Index

```go
cin := combi.CombinationIndex(3, 2)
fmt.Println(cin)
// => [[0 1] [0 2] [1 2]]
```

### Combination Ints

```go
ints := []int{1, 2, 3}
ci := combi.CombinationInts(ints, 2)
fmt.Println(ci)
// => [[1 2] [1 3] [2 3]]
```

### Combination Strs

```go
strs := []string{"a", "b", "c"}
cs := combi.CombinationStrs(strs, 2)
fmt.Println(cs)
// => [[a b] [a c] [b c]]
```

## Permutation

TBD

## Arithmetic Functions

### C

C returns C(n,k)

C(n,k) = n!/((n-k)!k!)

```go
c := combi.C(5, 2)
fmt.Println(c)
// => 10
```

### P

P calculates P(n, k)

P(n, k) = n! / k!

```go
p := combi.P(5, 2)
fmt.Println(p)
// => 20
```

### Factorical

Factorial returns n!

```go
factorial := combi.Factorial(5)
fmt.Println(factorial)
// => 120
```
