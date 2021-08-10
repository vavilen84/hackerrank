package digits_sum

import "testing"

type args struct {
	i int
}

type Case struct {
	name string
	args args
	want int
}

func Test_sum(t *testing.T) {
	case1 := Case{
		name: "Test case #1",
		args: args{
			i: 9875,
		},
		want: 2,
	}
	tests := []Case{
		case1,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.i); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
