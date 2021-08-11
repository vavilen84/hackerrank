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

	//case1 := struct {
	//	name string
	//	args args
	//	want []string
	//}{
	//	"test #1",
	//	args{
	//		crossword: []string{
	//			"+-++++++++",
	//			"+-++++++++",
	//			"+-++++++++",
	//			"+-----++++",
	//			"+-+++-++++",
	//			"+-+++-++++",
	//			"+++++-++++",
	//			"++------++",
	//			"+++++-++++",
	//			"+++++-++++",
	//		},
	//		words: "LONDON;DELHI;ICELAND;ANKARA",
	//	},
	//	[]string{
	//		"+L++++++++",
	//		"+O++++++++",
	//		"+N++++++++",
	//		"+DELHI++++",
	//		"+O+++C++++",
	//		"+N+++E++++",
	//		"+++++L++++",
	//		"++ANKARA++",
	//		"+++++N++++",
	//		"+++++D++++",
	//	},
	//}

	case2 := struct {
		name string
		args args
		want []string
	}{
		"test #2",
		args{
			crossword: []string{
				"+-++++++++",
				"+-++++++++",
				"+-------++",
				"+-++++++++",
				"+-++++++++",
				"+------+++",
				"+-+++-++++",
				"+++++-++++",
				"+++++-++++",
				"++++++++++",
			},
			words: "AGRA;NORWAY;ENGLAND;GWALIOR",
		},
		[]string{
			"+E++++++++",
			"+N++++++++",
			"+GWALIOR++",
			"+L++++++++",
			"+A++++++++",
			"+NORWAY+++",
			"+D+++G++++",
			"+++++R++++",
			"+++++A++++",
			"++++++++++",
		},
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		//case1,
		case2,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fillCrossword(tt.args.crossword, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillCrossword() = %v, want %v", got, tt.want)
			}
		})
	}
}
