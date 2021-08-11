package happy_ticket

import (
	"log"
	"strconv"
	"strings"
)

const (
	TicketLength = 6
)

func IsTicketHappy(i string) bool {
	if len(i) != TicketLength {
		log.Fatalln("Ticket must have 6 digits")
	}
	symbols := strings.Split(i, "")
	ints := make([]int, TicketLength)
	for k, v := range symbols {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("Ticket must contain integers only")
		}
		ints[k] = intValue
	}
	firstPart := ints[:TicketLength/2]
	lastPart := ints[TicketLength/2:]
	return sum(firstPart) == sum(lastPart)
}

func sum(i []int) int {
	res := 0
	for k := range i {
		res += i[k]
	}
	return res
}
