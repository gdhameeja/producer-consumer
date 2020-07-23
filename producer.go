package entities


const (
	numWidgetsLow = 1
	numWidgetsHigh = 3
)

// producer struct. The entity which produces the message.
type struct Producer {
	id string
	numWidgets int
}

// internal function used to create a single widget
func (p Producer) createWdiget() widget {

}

// function the consumer is supposed to call to get widgets
func (p Producer) provideWidget() []widget {

}
