package utils

import (
	"github.com/amanzanero/advent-of-code-2020/day16/ticket"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"regexp"
	"strconv"
)

func ParseTicketFields(lines []string) map[string]*ticket.Field {
	tfs := make(map[string]*ticket.Field)
	fieldReg := regexp.MustCompile("(: | or )")
	rangeReg := regexp.MustCompile("-")

	for _, line := range lines {
		fieldSpit := fieldReg.Split(line, 3)
		r1, r2 := fieldSpit[1], fieldSpit[2]

		r1Split := rangeReg.Split(r1, 2)
		r2Split := rangeReg.Split(r2, 2)

		rangeLower1, errL1 := strconv.Atoi(r1Split[0])
		rangeUpper1, errU1 := strconv.Atoi(r1Split[1])
		rangeLower2, errL2 := strconv.Atoi(r2Split[0])
		rangeUpper2, errU2 := strconv.Atoi(r2Split[1])
		lib.Check(errL1)
		lib.Check(errL2)
		lib.Check(errU1)
		lib.Check(errU2)

		tfs[fieldSpit[0]] = ticket.NewTicketField(rangeLower1, rangeUpper1, rangeLower2, rangeUpper2)
	}
	return tfs
}

func ParseTicket(line string) *ticket.Ticket {
	tkt := ticket.NewTicket()
	commaReg := regexp.MustCompile(",")

	idSplit := commaReg.Split(line, -1)

	for _, id := range idSplit {
		parsedId, err := strconv.Atoi(id)
		lib.Check(err)
		tkt.AddTicketId(parsedId)
	}

	return tkt
}

func ParseTickets(lines []string) []*ticket.Ticket {
	tkts := make([]*ticket.Ticket, 0)
	for _, line := range lines {
		tkt := ParseTicket(line)
		tkts = append(tkts, tkt)
	}
	return tkts
}
