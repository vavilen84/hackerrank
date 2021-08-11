package tic_tac_toe

import "testing"

type args struct {
	i [3][3]PlayerId
}

type Case struct {
	name string
	args args
	want Winner
}

func Test_findWinner(t *testing.T) {
	case1 := Case{
		name: "Test case #1 (vertical)",
		args: args{
			i: [3][3]PlayerId{
				{plusId, plusId, minusId},
				{plusId, notFilled, minusId},
				{notFilled, notFilled, minusId},
			},
		},
		want: minus,
	}
	case2 := Case{
		name: "Test case #1 (no winner)",
		args: args{
			i: [3][3]PlayerId{
				{plusId, minusId, plusId},
				{minusId, minusId, plusId},
				{notFilled, plusId, minusId},
			},
		},
		want: undefined,
	}
	case3 := Case{
		name: "Test case #1 (horizontal)",
		args: args{
			i: [3][3]PlayerId{
				{plusId, plusId, plusId},
				{minusId, minusId, notFilled},
				{minusId, notFilled, notFilled},
			},
		},
		want: plus,
	}
	tests := []Case{
		case1,
		case2,
		case3,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWinner(tt.args.i); got != tt.want {
				t.Errorf("findWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
