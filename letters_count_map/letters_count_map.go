package letters_count_map

/*
	Example:
	Input "asdfffggf"
	Output: a => 1; s => 1; d => 1; f => 4; g => 2
*/
func count(i string) map[string]int {
	result := make(map[string]int)
	for _, v := range i {
		if _, ok := result[string(v)]; ok {
			result[string(v)]++
		} else {
			result[string(v)] = 1
		}
	}
	return result
}
