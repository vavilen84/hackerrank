package crossword_puzzle

import (
	"reflect"
	"testing"
)

func Test_fillCrossword(t *testing.T) {
	type args struct {
		crossword []string
		words     string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"test #1",
			args{
				crossword: []string{
					"+-++++++++",
					"+-++++++++",
					"+-++++++++",
					"+-----++++",
					"+-+++-++++",
					"+-+++-++++",
					"+++++-++++",
					"++------++",
					"+++++-++++",
					"+++++-++++",
				},
				words: "LONDON;DELHI;ICELAND;ANKARA",
			},
			[][]string{
				{
					"+", "L", "+", "+", "+", "+", "+", "+", "+", "+",
				},
				{
					"+", "O", "+", "+", "+", "+", "+", "+", "+", "+",
				},
				{
					"+", "N", "+", "+", "+", "+", "+", "+", "+", "+",
				},
				{
					"+", "D", "E", "L", "H", "I", "+", "+", "+", "+",
				},
				{
					"+", "O", "+", "+", "+", "C", "+", "+", "+", "+",
				},
				{
					"+", "N", "+", "+", "+", "E", "+", "+", "+", "+",
				},
				{
					"+", "+", "+", "+", "+", "L", "+", "+", "+", "+",
				},
				{
					"+", "+", "A", "N", "K", "A", "R", "A", "+", "+",
				},
				{
					"+", "+", "+", "+", "+", "N", "+", "+", "+", "+",
				},
				{
					"+", "+", "+", "+", "+", "D", "+", "+", "+", "+",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fillCrossword(tt.args.crossword, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillCrossword() = %v, want %v", got, tt.want)
			}
		})
	}
}
