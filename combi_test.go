// Package combi implements some combinatoric functions
package combi

import (
	"reflect"
	"testing"
)

func TestP(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 P 0", args{0, 0}, 1},
		{"5 P 0", args{5, 0}, 1},
		{"5 P 1", args{5, 1}, 5},
		{"5 P 2", args{5, 2}, 20},
		{"5 P 3", args{5, 3}, 60},
		{"5 P 4", args{5, 4}, 120},
		{"5 P 5", args{5, 5}, 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("P() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestC(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 C 0", args{0, 0}, 1},
		{"5 C 0", args{5, 0}, 1},
		{"5 C 1", args{5, 1}, 5},
		{"5 C 2", args{5, 2}, 10},
		{"5 C 3", args{5, 3}, 10},
		{"5 C 4", args{5, 4}, 5},
		{"5 C 5", args{5, 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("C() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorical(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0!", args{0}, 1},
		{"1!", args{1}, 1},
		{"2!", args{2}, 2},
		{"3!", args{3}, 6},
		{"4!", args{4}, 24},
		{"5!", args{5}, 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorical(tt.args.n); got != tt.want {
				t.Errorf("Factorical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPowerSetIndex(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Power Set Index 0", args{0}, [][]int{{}}},
		{"Power Set Index 1", args{1}, [][]int{{}, {0}}},
		{"Power Set Index 2", args{2}, [][]int{{}, {0}, {1}, {0, 1}}},
		{"Power Set Index 3", args{3}, [][]int{{}, {0}, {1}, {0, 1}, {2}, {0, 2}, {1, 2}, {0, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowerSetIndex(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PowerSetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPowerSetInts(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Power Set Ints 0", args{[]int{}}, [][]int{{}}},
		{"Power Set Ints 1", args{[]int{1}}, [][]int{{}, {1}}},
		{"Power Set Ints 2", args{[]int{1, 2}}, [][]int{{}, {1}, {2}, {1, 2}}},
		{"Power Set Ints 3", args{[]int{1, 2, 3}}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowerSetInts(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PowerSetInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPowerSetStrs(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"Power Set Strs 0", args{[]string{}}, [][]string{{}}},
		{"Power Set Strs 1", args{[]string{"a"}}, [][]string{{}, {"a"}}},
		{"Power Set Strs 2", args{[]string{"a", "b"}}, [][]string{{}, {"a"}, {"b"}, {"a", "b"}}},
		{"Power Set Strs 3", args{[]string{"a", "b", "c"}}, [][]string{{}, {"a"}, {"b"}, {"a", "b"}, {"c"}, {"a", "c"}, {"b", "c"}, {"a", "b", "c"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowerSetStrs(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PowerSetStrs() = %v, want %v", got, tt.want)
			}
		})
	}
}
