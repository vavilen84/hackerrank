package digits_sum

import (
	"log"
	"strconv"
	"strings"
)

func sum(i int) int {
	intsArr := convertIntToIntsArr(i)
	res := sumRecursive(intsArr)
	return res
}

func convertIntToIntsArr(i int) []int {
	stringInput := strconv.Itoa(i)
	stringInputArr := strings.Split(stringInput, "")
	intsArr := []int{}
	for _, v := range stringInputArr {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		intsArr = append(intsArr, intValue)
	}
	return intsArr
}

func sumIntsArr(i []int) []int {
	s := 0
	for _, v := range i {
		s += v
	}
	if s > 9 {
		return convertIntToIntsArr(s)
	}
	return []int{s}
}

func sumRecursive(i []int) int {
	s := sumIntsArr(i)
	if len(s) > 1 {
		return sumRecursive(s)
	}
	return s[0]
}
