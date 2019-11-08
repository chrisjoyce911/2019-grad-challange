package score

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func Example_scoreGame() {

	lowestDiv := 4
	drawnNumbers := []int{7, 22, 24, 31, 33, 40}

	t := []lottoGame{
		{owner: "John", game: 1, numbers: []int{7, 9, 13, 24, 33, 40}},
		{owner: "John", game: 2, numbers: []int{16, 19, 22, 29, 31, 39}},
		{owner: "John", game: 3, numbers: []int{1, 7, 18, 22, 30, 36}},
		{owner: "Mary", game: 1, numbers: []int{2, 22, 13, 24, 32, 39}},
		{owner: "Mary", game: 2, numbers: []int{7, 22, 24, 31, 33, 40}},
		{owner: "Mary", game: 3, numbers: []int{3, 7, 18, 21, 37, 38}},
	}

	for _, game := range t {
		resault, wonDiv := scoreGame(game, drawnNumbers)
		if wonDiv >= lowestDiv {
			fmt.Println(resault)
		}
	}
	// Output:
	// John wins a division 5 on game #2 with matches [22 31] in game [16 19 22 29 31 39]
	// John wins a division 5 on game #3 with matches [7 22] in game [1 7 18 22 30 36]
	// Mary wins a division 5 on game #1 with matches [22 24] in game [2 22 13 24 32 39]
	// Mary wins a division 6 on game #3 with matches [7] in game [3 7 18 21 37 38]
}

func Benchmark_scoreGame(b *testing.B) {
	game := lottoGame{owner: "John", game: 1, numbers: []int{7, 9, 13, 24, 33, 40}}
	drawnNumbers := []int{7, 22, 24, 31, 33, 40}

	for n := 0; n < b.N; n++ {
		_, _ = scoreGame(game, drawnNumbers)
	}

}

func Test_scoreGame(t *testing.T) {

	w := []int{7, 22, 24, 31, 33, 40}

	tests := []struct {
		name string
		lg   lottoGame
		r    string
		d    int
		w    []int
	}{
		{
			name: "John 1",
			lg:   lottoGame{owner: "John", game: 1, numbers: []int{7, 9, 13, 24, 33, 40}},
			w:    w,
			d:    3,
			r:    "John wins a division 3 on game #1 with matches [7 24 33 40] in game [7 9 13 24 33 40]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, d := scoreGame(tt.lg, tt.w)
			assert.Equal(t, r, tt.r)
			assert.Equal(t, d, tt.d)
		})
	}
}

func Test_scoreEntery(t *testing.T) {
	type args struct {
		g []int
		w []int
	}
	tests := []struct {
		name string
		args args
		m    []int
		d    int
	}{
		{
			name: "div 4",
			args: args{
				g: []int{7, 9, 13, 24, 33, 40},
				w: []int{7, 22, 24, 31, 33, 40},
			},
			d: 4,
			m: []int{7, 24, 33, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, d := scoreEntery(tt.args.g, tt.args.w)
			assert.Equal(t, tt.m, m)
			assert.Equal(t, tt.d, d)
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "match",
			args: args{
				e: 7,
				s: []int{7, 9, 13, 24, 33, 40},
			},
			want: true,
		},
		{
			name: "no match",
			args: args{
				e: 6,
				s: []int{7, 9, 13, 24, 33, 40},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, contains(tt.args.s, tt.args.e), tt.want)
		})
	}
}
