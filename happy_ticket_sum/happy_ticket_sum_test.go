package happy_ticket_sum

import "testing"

type args struct {
	start string
	end   string
}

type Case struct {
	name string
	args args
	want int
}

func Test_happyTicketSum(t *testing.T) {
	case1 := Case{
		name: "Test case #1",
		args: args{
			start: "000000",
			end:   "999999",
		},
		want: 55252,
	}
	tests := []Case{
		case1,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := happyTicketSum(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("happyTicketSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
