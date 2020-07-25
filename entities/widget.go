package entities

// the widget struct. Basically resembles the message producers can create and
// consumers can process.
type Widget struct {
	id string
	message string
	producedBy Producer
}
