package entities

// the widget struct. Basically resembles the message producers can create and
// consumers can process.
type struct Widget {
	id string
	message string
	producedBy producer
}
