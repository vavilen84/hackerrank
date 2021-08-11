package happy_ticket_sum

import (
	"github.com/vavilen84/hackerrank/happy_ticket"
	"log"
	"strconv"
	"strings"
)

func happyTicketSum(start, end string) int {
	res := 0
	startIntVal, err := strconv.Atoi(start)
	if err != nil {
		log.Fatalln("start arg invalid")
	}
	endIntVal, err := strconv.Atoi(end)
	if err != nil {
		log.Fatalln("end arg invalid")
	}
	for i := startIntVal; i <= endIntVal; i++ {
		if happy_ticket.IsTicketHappy(convertToString(i)) {
			res++
		}
	}
	return res
}

func convertToString(i int) string {
	res := []string{"0", "0", "0", "0", "0", "0"}
	strValue := strconv.Itoa(i)
	sliceValue := strings.Split(strValue, "")
	copy(res[happy_ticket.TicketLength-len(sliceValue):], sliceValue)
	return strings.Join(res, "")
}
