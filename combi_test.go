// Package combi implements some combinatoric functions
package combi

import "testing"

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
