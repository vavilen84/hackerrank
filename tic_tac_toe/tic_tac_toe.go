package tic_tac_toe

const (
	plus      Winner = 1
	minus     Winner = 2
	undefined Winner = 3

	plusId    PlayerId = "+"
	minusId   PlayerId = "-"
	notFilled PlayerId = ""
)

type PlayerId string

type Winner int

func findWinner(i [3][3]PlayerId) Winner {
	winner := findWinnerHorizontal(i)
	if winner != undefined {
		return winner
	}
	winner = findWinnerVertical(i)
	if winner != undefined {
		return winner
	}
	return findWinnerObliquely(i)
}

func findWinnerObliquely(input [3][3]PlayerId) Winner {
	plusCounter := 0
	minusCounter := 0
	// left to right
	for i := 0; i < 3; i++ {
		switch input[i][i] {
		case plusId:
			plusCounter++
		case minusId:
			minusCounter++
		}
	}
	if plusCounter == 3 {
		return plus
	}
	if minusCounter == 3 {
		return minus
	}
	// left to right
	helpCounter := 0
	plusCounter = 0
	minusCounter = 0
	for i := 2; i >= 0; i-- {
		switch input[helpCounter][i] {
		case plusId:
			plusCounter++
		case minusId:
			minusCounter++
		}
		helpCounter++
	}
	if plusCounter == 3 {
		return plus
	}
	if minusCounter == 3 {
		return minus
	}
	return undefined
}

func findWinnerHorizontal(input [3][3]PlayerId) Winner {
	for i := 0; i < 3; i++ {
		plusCounter := 0
		minusCounter := 0
		for k, _ := range input[i] {
			switch input[i][k] {
			case plusId:
				plusCounter++
			case minusId:
				minusCounter++
			}
		}
		if plusCounter == 3 {
			return plus
		}
		if minusCounter == 3 {
			return minus
		}
	}
	return undefined
}

func findWinnerVertical(input [3][3]PlayerId) Winner {
	for columnKey := 0; columnKey < 3; columnKey++ {
		plusCounter := 0
		minusCounter := 0
		for rowKey := 0; rowKey < 3; rowKey++ {
			switch input[rowKey][columnKey] {
			case plusId:
				plusCounter++
			case minusId:
				minusCounter++
			}
		}
		if plusCounter == 3 {
			return plus
		}
		if minusCounter == 3 {
			return minus
		}
	}
	return undefined
}
