package entities

import "time"

const MAX_NUM_WIDGETS = 10

// consumer struct. The entity which processes/consumes the widget.
type Consumer struct {
	id int
	ch chan Widget
	proxy Proxy
	widgets []Widget
}

// constructor function to provide a Consumer with buffered channel
// proxy is supposed to be shared between all instances of Consumers
func ConstructConsumer(id int, proxy Proxy) Consumer {
	return Consumer{
		id: id,
		ch: make(chan Widget, MAX_NUM_WIDGETS),
		proxy: proxy,
	}
}

// start the consumer to work, called by main
func (c Consumer) Start() {
	for {
		c.RequestWidgets()
	}
}

func (c Consumer) RequestWidgets() {
	// TODO: randomSeconds should have a random value between 1 and 5
	var randomSeconds time.Duration = 5  
	if len(c.widgets) >= 10 {
		time.Sleep(randomSeconds * time.Second)
		c.flush()
	}
	// TODO: randomWidgetsNum should be a random number 1 and 3
	randomWidgetsNum := 3
	c.proxy.WriteWidgets(randomWidgetsNum, c.id, c.ch)
	for widget := range c.ch {
		c.widgets = append(c.widgets, widget)
	}
}

// reset consumer to initial state (using widgets)
func (c Consumer) flush() {
	c.widgets = nil
}
