package ticket

type Ticket struct {
	Fields []int
}

func NewTicket() *Ticket {
	return &Ticket{
		make([]int, 0),
	}
}

func (t *Ticket) AddTicketId(id int) {
	t.Fields = append(t.Fields, id)
}
