package entities

import "log"

const (
	numWidgetsLow = 1
	numWidgetsHigh = 3
)

// producer struct. The entity which produces the message.
type Producer struct {
	Id int
}

func (p Producer) ProvideWidget() Widget {
	log.Print("Widget created by producer: ", p.Id)
	return Widget {}
}
