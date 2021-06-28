package crossword_puzzle

import "strings"

// https://www.hackerrank.com/challenges/crossword-puzzle/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=recursion-backtracking

func fillCrossword(crossword []string, words string) [][]string {
	mtxSideLength := len(crossword)
	wordsList := strings.Split(words, ";")
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
	uniqueLengthWords := make(map[int]string)
	for length, values := range wordsLengthMap {
		if len(values) == 1 {
			uniqueLengthWords[length] = values[0]
			delete(wordsLengthMap, length)
		}
	}

	crosswordMtx := make([][]string, mtxSideLength)
	for k, _ := range crossword {
		crosswordMtx[k] = strings.Split(crossword[k], "")
	}

	// search minuses horizontal
	for lineKey, _ := range crosswordMtx {
		counter := 0
		for k, symbol := range crosswordMtx[lineKey] {
			// handle last idx
			if k == mtxSideLength-1 {
				if symbol == "-" && crosswordMtx[lineKey][k-1] == "-" {
					counter++
				}
				if _, ok := uniqueLengthWords[counter]; ok {
					copy(crosswordMtx[lineKey][k-counter+1:], strings.Split(uniqueLengthWords[counter], ""))
					delete(uniqueLengthWords, counter)
				}
				break
			} else {
				if symbol == "-" {
					counter++
					if crosswordMtx[lineKey][k+1] == "-" {
						continue
					} else {
						if _, ok := uniqueLengthWords[counter]; ok {
							copy(crosswordMtx[lineKey][k-counter+1:], strings.Split(uniqueLengthWords[counter], ""))
							delete(uniqueLengthWords, counter)
						}
						counter = 0
					}
				}
			}
		}
	}

	// search minuses vertical
	for vertical := 0; vertical < mtxSideLength; vertical++ {
		counter := 0
		for horizontal := 0; horizontal < mtxSideLength; horizontal++ {
			// handle last idx
			if horizontal == mtxSideLength-1 {
				if crosswordMtx[horizontal][vertical] != "+" && crosswordMtx[horizontal-1][vertical] != "+" {
					counter++
				}
				if _, ok := uniqueLengthWords[counter]; ok {
					c := 0
					str := strings.Split(uniqueLengthWords[counter], "")
					for h := horizontal - counter + 1; h < mtxSideLength; h++ {
						crosswordMtx[h][vertical] = str[c]
						c++
					}
					delete(uniqueLengthWords, counter)
				}
				break
			} else {
				if crosswordMtx[horizontal][vertical] != "+" {
					counter++
					if crosswordMtx[horizontal+1][vertical] != "+" {
						continue
					} else {
						if _, ok := uniqueLengthWords[counter]; ok {
							c := 0
							str := strings.Split(uniqueLengthWords[counter], "")
							for h := horizontal - counter + 1; h < mtxSideLength; h++ {
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

	// fill rest horizontal words
	for lineKey, _ := range crosswordMtx {
		counter := 0
		for k, symbol := range crosswordMtx[lineKey] {
			// handle last idx
			if k == mtxSideLength-1 {
				if symbol == "-" && crosswordMtx[lineKey][k-1] == "-" {
					counter++
				}
				if _, ok := wordsLengthMap[counter]; ok {
					wordsSlice := wordsLengthMap[counter]
					for wordsSliceKey, _ := range wordsSlice {

						mismatch := false
						for symbolKey, _ := range wordsSlice[wordsSliceKey] {
							key := k - counter + 1 + symbolKey
							if crosswordMtx[lineKey][key] != "-" &&
								crosswordMtx[lineKey][key] != "+" &&
								crosswordMtx[lineKey][key] != string(wordsSlice[wordsSliceKey][symbolKey]) {
								mismatch = true
								break
							}
						}
						if mismatch {
							continue
						}

						copy(crosswordMtx[lineKey][k-counter+1:], strings.Split(wordsSlice[wordsSliceKey], ""))
						wordsLengthMap[counter] = append(wordsLengthMap[counter][:wordsSliceKey], wordsLengthMap[counter][wordsSliceKey+1:]...)
					}
				}
				break
			} else {
				if symbol != "+" {
					counter++
					next := crosswordMtx[lineKey][k+1]
					if next != "+" {
						continue
					} else {
						if _, ok := wordsLengthMap[counter]; ok {
							wordsSlice := wordsLengthMap[counter]
							for wordsSliceKey, _ := range wordsSlice {
								mismatch := false
								for symbolKey, _ := range wordsSlice[wordsSliceKey] {
									key := k - counter + 1 + symbolKey
									if crosswordMtx[lineKey][key] != "-" &&
										crosswordMtx[lineKey][key] != "+" &&
										crosswordMtx[lineKey][key] != string(wordsSlice[wordsSliceKey][symbolKey]) {
										mismatch = true
										break
									}
								}
								if mismatch {
									continue
								}
								copy(crosswordMtx[lineKey][k-counter+1:], strings.Split(wordsSlice[wordsSliceKey], ""))
								wordsLengthMap[counter] = append(wordsLengthMap[counter][:wordsSliceKey], wordsLengthMap[counter][wordsSliceKey+1:]...)
							}
						}
						counter = 0
					}
				}
			}
		}
	}

	// fill rest vertical words
	for vertical := 0; vertical < mtxSideLength; vertical++ {
		counter := 0
		for horizontal := 0; horizontal < mtxSideLength; horizontal++ {
			// handle last idx
			if horizontal == mtxSideLength-1 {
				if crosswordMtx[horizontal][vertical] != "+" && crosswordMtx[horizontal-1][vertical] != "+" {
					counter++
				}
				if _, ok := wordsLengthMap[counter]; ok {
					wordsSlice := wordsLengthMap[counter]
					for wordsSliceKey, _ := range wordsSlice {
						mismatch := false
						for symbolKey, _ := range wordsSlice[wordsSliceKey] {
							key := horizontal - counter + 1 + symbolKey
							if crosswordMtx[key][vertical] != "-" &&
								crosswordMtx[key][vertical] != "+" &&
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
				if crosswordMtx[horizontal][vertical] != "+" {
					counter++
					if crosswordMtx[horizontal+1][vertical] != "+" {
						continue
					} else {
						if _, ok := wordsLengthMap[counter]; ok {
							wordsSlice := wordsLengthMap[counter]
							for wordsSliceKey, _ := range wordsSlice {
								mismatch := false
								for symbolKey, _ := range wordsSlice[wordsSliceKey] {
									key := horizontal - counter + 1 + symbolKey
									if crosswordMtx[key][vertical] != "-" &&
										crosswordMtx[key][vertical] != "+" &&
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

	return crosswordMtx
}
