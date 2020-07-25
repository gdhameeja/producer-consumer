package entities


const (
	numWidgetsLow = 1
	numWidgetsHigh = 3
)

// producer struct. The entity which produces the message.
type Producer struct {
	id string
	numWidgets int
}

func (p Producer) CreateWidget(channel chan Widget) {
}
