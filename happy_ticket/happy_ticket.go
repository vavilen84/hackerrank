package happy_ticket

import (
	"log"
	"strconv"
	"strings"
)

const ticketLength = 6

func isTicketHappy(i string) bool {
	if len(i) != ticketLength {
		log.Fatalln("Ticket must have 6 digits")
	}
	symbols := strings.Split(i, "")
	ints := make([]int, ticketLength)
	for k, v := range symbols {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("Ticket must contain integers only")
		}
		ints[k] = intValue
	}
	firstPart := ints[:ticketLength/2]
	lastPart := ints[ticketLength/2:]
	return sum(firstPart) == sum(lastPart)
}

func sum(i []int) int {
	res := 0
	for k := range i {
		res += i[k]
	}
	return res
}
