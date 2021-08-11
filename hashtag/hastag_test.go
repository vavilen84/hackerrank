package hashtag

import "testing"

type args struct {
	i string
}

type Case struct {
	name string
	args args
	want string
}

func Test_create(t *testing.T) {
	case1 := Case{
		name: "Test case #1",
		args: args{
			i: "Hello World!",
		},
		want: "#helloworld",
	}
	tests := []Case{
		case1,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := create(tt.args.i); got != tt.want {
				t.Errorf("create() = %v, want %v", got, tt.want)
			}
		})
	}
}
