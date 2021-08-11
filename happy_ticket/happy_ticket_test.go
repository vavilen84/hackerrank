package happy_ticket

import "testing"

type args struct {
	i string
}

type Case struct {
	name string
	args args
	want bool
}

func Test_isTicketHappy(t *testing.T) {
	case1 := Case{
		name: "Test case #1",
		args: args{
			i: "001010",
		},
		want: true,
	}
	case2 := Case{
		name: "Test case #2",
		args: args{
			i: "001020",
		},
		want: false,
	}
	tests := []Case{
		case1,
		case2,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTicketHappy(tt.args.i); got != tt.want {
				t.Errorf("isTicketHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
