package crossword_puzzle

import (
	"log"
	"strings"
)

// https://www.hackerrank.com/challenges/crossword-puzzle/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=recursion-backtracking

const (
	fillable    = "-"
	notFillable = "+"
)

func fillCrossword(crossword []string, words string) []string {
	height := len(crossword)
	width := len(crossword[0])

	// parse words input string to slice
	wordsList := strings.Split(words, ";")

	// Ex.: 3 => {abc, def}, 4 => {abcd}, etc ...
	wordsLengthMap := make(map[int][]string)
	for k, _ := range wordsList {
		length := len(wordsList[k])
		if _, ok := wordsLengthMap[length]; ok {
			wordsLengthMap[length] = append(wordsLengthMap[length], wordsList[k])
		} else {
			wordsLengthMap[length] = make([]string, 0)
			wordsLengthMap[length] = append(wordsLengthMap[length], wordsList[k])
		}
	}

	// Exclude from wordsLengthMap words with unique length
	uniqueLengthWords := make(map[int]string)
	for length, values := range wordsLengthMap {
		if len(values) == 1 {
			uniqueLengthWords[length] = values[0]
			delete(wordsLengthMap, length)
		}
	}

	// get a new crossword matrix which we will fill
	crosswordMtx := getCrosswordMtx(height, crossword)

	// the main idea to fill crosswordMtx is:
	// 1. fill words with unique length
	// 2. then fill other words
	// flows - all available filling options
	flows := 4

	// take copy because we need their modifications on each flow
	uniqueLengthWordsCopy, wordsLengthMapCopy := copyWordsData(uniqueLengthWords, wordsLengthMap)

	for i := 0; i <= flows; i++ {

		if len(uniqueLengthWordsCopy) == 0 && len(wordsLengthMapCopy) == 0 {
			// if all words maps are empty - we are done
			break
		} else {
			// previous flow was unsuccessful, we need new crosswordMtx and maps with words
			crosswordMtx = getCrosswordMtx(height, crossword)
			uniqueLengthWordsCopy, wordsLengthMapCopy = copyWordsData(uniqueLengthWords, wordsLengthMap)
		}

		switch i {
		case 0:
			fillFlowA(crosswordMtx, width, height, uniqueLengthWordsCopy, wordsLengthMapCopy)
		case 1:
			fillFlowB(crosswordMtx, width, height, uniqueLengthWordsCopy, wordsLengthMapCopy)
		case 2:
			fillFlowC(crosswordMtx, width, height, uniqueLengthWordsCopy, wordsLengthMapCopy)
		case 3:
			fillFlowD(crosswordMtx, width, height, uniqueLengthWordsCopy, wordsLengthMapCopy)
		case 4:
			log.Fatal("crossword data is invalid")
		}
	}

	// convert [][]string => []string
	res := make([]string, height)
	for k, _ := range crosswordMtx {
		for k1, _ := range crosswordMtx[k] {
			res[k] += crosswordMtx[k][k1]
		}
	}

	return res
}

func getCrosswordMtx(height int, crossword []string) [][]string {
	crosswordMtx := make([][]string, height)
	for k, _ := range crossword {
		crosswordMtx[k] = strings.Split(crossword[k], "")
	}
	return crosswordMtx
}

func copyWordsData(
	uniqueLengthWords map[int]string,
	wordsLengthMap map[int][]string,
) (uniqueLengthWordsCopy map[int]string, wordsLengthMapCopy map[int][]string) {
	uniqueLengthWordsCopy = make(map[int]string)
	for k, v := range uniqueLengthWords {
		uniqueLengthWordsCopy[k] = v
	}

	wordsLengthMapCopy = make(map[int][]string)
	for k, v := range wordsLengthMap {
		wordsLengthMapCopy[k] = v
	}

	return
}

func fillFlowA(crosswordMtx [][]string, width, height int, uniqueLengthWords map[int]string, wordsLengthMap map[int][]string) {
	fillUniqueLengthWordsHorizontal(crosswordMtx, width, uniqueLengthWords)
	fillUniqueLengthWordsVertical(crosswordMtx, width, height, uniqueLengthWords)
	fillRestHorizontalWords(crosswordMtx, height, wordsLengthMap)
	fillVerticalWords(crosswordMtx, width, height, wordsLengthMap)
}

func fillFlowB(crosswordMtx [][]string, width, height int, uniqueLengthWords map[int]string, wordsLengthMap map[int][]string) {
	fillUniqueLengthWordsVertical(crosswordMtx, width, height, uniqueLengthWords)
	fillUniqueLengthWordsHorizontal(crosswordMtx, width, uniqueLengthWords)
	fillRestHorizontalWords(crosswordMtx, height, wordsLengthMap)
	fillVerticalWords(crosswordMtx, width, height, wordsLengthMap)
}

func fillFlowC(crosswordMtx [][]string, width, height int, uniqueLengthWords map[int]string, wordsLengthMap map[int][]string) {
	fillUniqueLengthWordsHorizontal(crosswordMtx, width, uniqueLengthWords)
	fillUniqueLengthWordsVertical(crosswordMtx, width, height, uniqueLengthWords)
	fillVerticalWords(crosswordMtx, width, height, wordsLengthMap)
	fillRestHorizontalWords(crosswordMtx, height, wordsLengthMap)
}

func fillFlowD(crosswordMtx [][]string, width, height int, uniqueLengthWords map[int]string, wordsLengthMap map[int][]string) {
	fillUniqueLengthWordsVertical(crosswordMtx, width, height, uniqueLengthWords)
	fillUniqueLengthWordsHorizontal(crosswordMtx, width, uniqueLengthWords)
	fillVerticalWords(crosswordMtx, width, height, wordsLengthMap)
	fillRestHorizontalWords(crosswordMtx, height, wordsLengthMap)
}

func fillUniqueLengthWordsVertical(crosswordMtx [][]string, width, height int, uniqueLengthWords map[int]string) {
	for vertical := 0; vertical < height; vertical++ {
		// counter - 'fillable' symbols. Ex.: if we have 5 'fillable' symbols in sequence -
		// we try to fill word which consists of 5 letters (if exists).
		counter := 0
		for horizontal := 0; horizontal < width; horizontal++ {
			// handle last idx
			if horizontal == width-1 {
				if crosswordMtx[horizontal][vertical] != notFillable && crosswordMtx[horizontal-1][vertical] != notFillable {
					counter++
				}
				if _, ok := uniqueLengthWords[counter]; ok {
					c := 0
					str := strings.Split(uniqueLengthWords[counter], "")
					for h := horizontal - counter + 1; h < width; h++ {
						crosswordMtx[h][vertical] = str[c]
						c++
					}
					delete(uniqueLengthWords, counter)
				}
				break
			} else {
				if crosswordMtx[horizontal][vertical] != notFillable {
					counter++
					if crosswordMtx[horizontal+1][vertical] != notFillable {
						continue
					} else {
						if _, ok := uniqueLengthWords[counter]; ok {
							c := 0
							str := strings.Split(uniqueLengthWords[counter], "")
							strLen := len(str)
							for h := horizontal - counter + 1; h < width && c <= strLen-1; h++ {
								crosswordMtx[h][vertical] = str[c]
								c++
							}
							delete(uniqueLengthWords, counter)
						}
						counter = 0
					}
				}
			}
		}
	}
}

func isSymbolFillable(symbol string) bool {
	return string(symbol) == fillable
}

func isPreviousSymbolFillable(crosswordMtx [][]string, rowKey, columnKey int) bool {
	return crosswordMtx[rowKey][columnKey-1] == fillable
}

func isLastHorizontalIdx(columnKey, width int) bool {
	return columnKey == width-1
}

func fillWordHorizontal(crosswordMtx [][]string, rowKey, columnKey, counter int, uniqueLengthWords map[int]string) {
	// put word to crossword
	copy(crosswordMtx[rowKey][columnKey-counter+1:], strings.Split(uniqueLengthWords[counter], ""))
	// delete from word map
	delete(uniqueLengthWords, counter)
}

func isNextSymbolFillable(crosswordMtx [][]string, rowKey, columnKey int) bool {
	return crosswordMtx[rowKey][columnKey+1] == fillable
}

func fillUniqueLengthWordsHorizontal(crosswordMtx [][]string, width int, uniqueLengthWords map[int]string) {
	for rowKey, _ := range crosswordMtx {
		// counter - 'fillable' symbols. Ex.: if we have 5 'fillable' symbols in sequence -
		// we try to fill word which consists of 5 letters (if exists).
		counter := 0
		for columnKey, symbol := range crosswordMtx[rowKey] {
			if isLastHorizontalIdx(columnKey, width) {
				if isSymbolFillable(symbol) && isPreviousSymbolFillable(crosswordMtx, rowKey, columnKey) {
					counter++
				}
				if _, ok := uniqueLengthWords[counter]; ok {
					fillWordHorizontal(crosswordMtx, rowKey, columnKey, counter, uniqueLengthWords)
				}
				break
			} else {
				if isSymbolFillable(symbol) {
					counter++
					if isNextSymbolFillable(crosswordMtx, rowKey, columnKey) {
						continue
					} else {
						fillWordHorizontal(crosswordMtx, rowKey, columnKey, counter, uniqueLengthWords)
						// end of sequence, so we need to reset counter
						counter = 0
					}
				}
			}
		}
	}
}

func isLastVerticalIdx(columnKey, height int) bool {
	return columnKey == height-1
}

func fillRestHorizontalWords(crosswordMtx [][]string, height int, wordsLengthMap map[int][]string) {
	for rowKey, _ := range crosswordMtx {
		// counter - 'fillable' symbols. Ex.: if we have 5 'fillable' symbols in sequence -
		// we try to fill word which consists of 5 letters (if exists).
		counter := 0
		for columnKey, symbol := range crosswordMtx[rowKey] {
			// handle last idx
			if isLastVerticalIdx(columnKey, height) {
				if isSymbolFillable(symbol) && isPreviousSymbolFillable(crosswordMtx, rowKey, columnKey) {
					counter++
				}
				if _, ok := wordsLengthMap[counter]; ok {
					wordsSlice := wordsLengthMap[counter]
					for wordsSliceKey, _ := range wordsSlice {
						mismatch := false
						for symbolKey, _ := range wordsSlice[wordsSliceKey] {
							key := columnKey - counter + 1 + symbolKey
							if crosswordMtx[rowKey][key] != fillable &&
								crosswordMtx[rowKey][key] != notFillable &&
								crosswordMtx[rowKey][key] != string(wordsSlice[wordsSliceKey][symbolKey]) {
								mismatch = true
								break
							}
						}
						if mismatch {
							continue
						}

						copy(crosswordMtx[rowKey][columnKey-counter+1:], strings.Split(wordsSlice[wordsSliceKey], ""))
						wordsLengthMap[counter] = append(wordsLengthMap[counter][:wordsSliceKey], wordsLengthMap[counter][wordsSliceKey+1:]...)
					}
				}
				break
			} else {
				if symbol != "+" {
					counter++
					next := crosswordMtx[rowKey][columnKey+1]
					if next != "+" {
						continue
					} else {
						if _, ok := wordsLengthMap[counter]; ok {
							wordsSlice := wordsLengthMap[counter]
							for wordsSliceKey, _ := range wordsSlice {
								mismatch := false
								for symbolKey, _ := range wordsSlice[wordsSliceKey] {
									key := columnKey - counter + 1 + symbolKey
									if crosswordMtx[rowKey][key] != fillable &&
										crosswordMtx[rowKey][key] != notFillable &&
										crosswordMtx[rowKey][key] != string(wordsSlice[wordsSliceKey][symbolKey]) {
										mismatch = true
										break
									}
								}
								if mismatch {
									continue
								}
								copy(crosswordMtx[rowKey][columnKey-counter+1:], strings.Split(wordsSlice[wordsSliceKey], ""))
								wordsLengthMap[counter] = append(wordsLengthMap[counter][:wordsSliceKey], wordsLengthMap[counter][wordsSliceKey+1:]...)
							}
						}
						counter = 0
					}
				}
			}
		}
	}
}

func fillVerticalWords(crosswordMtx [][]string, width, height int, wordsLengthMap map[int][]string) {
	for vertical := 0; vertical < height; vertical++ {
		// counter - 'fillable' symbols. Ex.: if we have 5 'fillable' symbols in sequence -
		// we try to fill word which consists of 5 letters (if exists).
		counter := 0
		for horizontal := 0; horizontal < width; horizontal++ {
			// handle last idx
			if horizontal == width-1 {
				if crosswordMtx[horizontal][vertical] != notFillable && crosswordMtx[horizontal-1][vertical] != notFillable {
					counter++
				}
				if _, ok := wordsLengthMap[counter]; ok {
					wordsSlice := wordsLengthMap[counter]
					for wordsSliceKey, _ := range wordsSlice {
						mismatch := false
						for symbolKey, _ := range wordsSlice[wordsSliceKey] {
							key := horizontal - counter + 1 + symbolKey
							if crosswordMtx[key][vertical] != fillable &&
								crosswordMtx[key][vertical] != notFillable &&
								crosswordMtx[key][vertical] != string(wordsSlice[wordsSliceKey][symbolKey]) {
								mismatch = true
								break
							}
						}
						if mismatch {
							continue
						}
						for symbolKey, _ := range wordsSlice[wordsSliceKey] {
							key := horizontal - counter + 1 + symbolKey
							crosswordMtx[key][vertical] = string(wordsSlice[wordsSliceKey][symbolKey])
						}
						wordsLengthMap[counter] = append(wordsLengthMap[counter][:wordsSliceKey], wordsLengthMap[counter][wordsSliceKey+1:]...)
					}
				}
				break
			} else {
				if crosswordMtx[horizontal][vertical] != notFillable {
					counter++
					if crosswordMtx[horizontal+1][vertical] != notFillable {
						continue
					} else {
						if _, ok := wordsLengthMap[counter]; ok {
							wordsSlice := wordsLengthMap[counter]
							for wordsSliceKey, _ := range wordsSlice {
								mismatch := false
								for symbolKey, _ := range wordsSlice[wordsSliceKey] {
									key := horizontal - counter + 1 + symbolKey
									if crosswordMtx[key][vertical] != fillable &&
										crosswordMtx[key][vertical] != notFillable &&
										crosswordMtx[key][vertical] != string(wordsSlice[wordsSliceKey][symbolKey]) {
										mismatch = true
										break
									}
								}
								if mismatch {
									continue
								}
								for symbolKey, _ := range wordsSlice[wordsSliceKey] {
									key := horizontal - counter + 1 + symbolKey
									crosswordMtx[key][vertical] = string(wordsSlice[wordsSliceKey][symbolKey])
								}
								wordsLengthMap[counter] = append(wordsLengthMap[counter][:wordsSliceKey], wordsLengthMap[counter][wordsSliceKey+1:]...)
							}
						}
						counter = 0
					}
				}
			}
		}
	}
}
