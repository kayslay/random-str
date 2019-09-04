package str

import (
	"math/rand"
	"testing"
	"time"
)

func TestWriteN_char(t *testing.T) {
	// use a predictable seed
	r = rand.New(rand.NewSource(1))
	type args struct {
		n      int
		format []rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"10 Capital letters",
			args{10, []rune{'A'}},
			"XVLBZGBAIC",
		},
		{
			"10 Small letters",
			args{10, []rune{'a'}},
			"mrajwwhthc",
		},
		{
			"10 digits",
			args{10, []rune{'d'}},
			"5688777805",
		},
		{
			"10 symbols",
			args{10, []rune{'s'}},
			"=~@:>=^;_|",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WriteN(tt.args.n, tt.args.format...); got != tt.want {
				t.Errorf("WriteN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteN_stringLength(t *testing.T) {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	type args struct {
		n      int
		format []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"return a string of length 10",
			args{10, []rune{'s'}},
			10,
		},
		{
			"return a string of length 12",
			args{12, []rune{}},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WriteN(tt.args.n, tt.args.format...); len(got) != tt.want {
				t.Errorf("WriteN() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestWriteFromFormat(t *testing.T) {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	type args struct {
		format string
	}
	tests := []struct {
		name string
		args args
		want [][2]int32
	}{
		// TODO: Add test cases.
		{
			"Char must exist in format range",
			args{"Aad*"},
			[][2]int32{[2]int32{65, 90}, [2]int32{97, 122}, [2]int32{48, 57}, [2]int32{32, 127}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range WriteFromFormat(tt.args.format) {
				if v < tt.want[i][0] || v > tt.want[i][1] {
					t.Errorf("Expected %s to be in range %d and %d", string(v), tt.want[i][0], tt.want[i][1])
				}
			}

		})
	}
}
