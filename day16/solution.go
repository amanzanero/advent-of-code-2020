package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/day16/ticket"
	"github.com/amanzanero/advent-of-code-2020/day16/utils"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"strings"
)

func main() {
	parsedLines := lib.ParseLines("day16/input.txt")

	s1 := lib.Elapsed("-- took: ")
	// find first blank
	endTicketFields := 0
	for i, line := range parsedLines {
		if line == "" {
			endTicketFields = i
			break
		}
	}
	parsedTicketFields := utils.ParseTicketFields(parsedLines[:endTicketFields])

	// find my ticket
	myTicketLine := 0
	for i := endTicketFields + 1; i < len(parsedLines); i++ {
		if parsedLines[i] == "your ticket:" {
			myTicketLine = i + 1
			break
		}
	}

	myTicket := utils.ParseTicket(parsedLines[myTicketLine])
	nearbyTickets := utils.ParseTickets(parsedLines[myTicketLine+3:])

	invalidTickets := 0
	validNearbyTickets := make([]*ticket.Ticket, 0)
	for _, nbyTkt := range nearbyTickets {
		isValidTicket, invalidTicketField := IsValidTicket(parsedTicketFields, nbyTkt)
		if !isValidTicket {
			invalidTickets += invalidTicketField
		} else {
			validNearbyTickets = append(validNearbyTickets, nbyTkt)
		}
	}
	fmt.Printf("Invalid tickets: %d\n", invalidTickets)
	s1()

	s2 := lib.Elapsed("-- took: ")
	// create a 2d matrix for our graph representation
	// keys are
	bipartiteGraph := make(map[string][]bool)
	for fieldName, ticketField := range parsedTicketFields {
		bipartiteGraph[fieldName] = make([]bool, len(parsedTicketFields))
		for i := 0; i < len(parsedTicketFields); i++ {
			bipartiteGraph[fieldName][i] = true
		}

		for _, tkt := range validNearbyTickets {
			for i, tktId := range tkt.Fields {
				bipartiteGraph[fieldName][i] = bipartiteGraph[fieldName][i] && ticketField.IsValidField(tktId)
			}
		}
	}

	// this map will contain the positions for each ticket class
	matchR := make([]string, len(bipartiteGraph))

	var bipartiteMatch func(string, []bool) bool
	bipartiteMatch = func(ticketClass string, seen []bool) bool {
		for v := range bipartiteGraph[ticketClass] {
			if bipartiteGraph[ticketClass][v] && seen[v] == false {
				seen[v] = true

				if matchR[v] == "" || bipartiteMatch(matchR[v], seen) {
					matchR[v] = ticketClass
					return true
				}
			}
		}
		return false
	}

	// make our maximum bipartite match
	for ticketClass := range parsedTicketFields {
		seen := make([]bool, len(bipartiteGraph))
		bipartiteMatch(ticketClass, seen)
	}

	// now just take the ones that start with departure
	result := 1
	for i, match := range matchR {
		if strings.HasPrefix(match, "departure") {
			result *= myTicket.Fields[i]
		}
	}
	fmt.Printf("Matched fields that start with 'departure' product: %d\n", result)

	s2()
}

func IsValidTicket(tfs map[string]*ticket.Field, ticket *ticket.Ticket) (bool, int) {
	for _, ticketId := range ticket.Fields {
		isValid := false
		for _, tf := range tfs {
			if tf.IsValidField(ticketId) {
				isValid = true
				break
			}
		}
		if !isValid {
			return false, ticketId
		}
	}
	return true, -1
}
