package letters_count_map

import (
	"reflect"
	"testing"
)

type Case struct {
	name string
	args args
	want map[string]int
}

type args struct {
	i string
}

func Test_count(t *testing.T) {

	case1 := Case{
		name: "Test case #1",
		args: args{
			i: "asdfffggf",
		},
		want: map[string]int{"a": 1, "s": 1, "d": 1, "f": 4, "g": 2},
	}
	tests := []Case{
		case1,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}
