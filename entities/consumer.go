package entities

import (
	"time"
	"math/rand"
	"log"
	"sync"
)

const (
	MAX_NUM_WIDGETS = 10
	minConsumerSleep = 1
	maxConsumerSleep = 5
)

// consumer struct. The entity which processes/consumes the widget.
type Consumer struct {
	id int
	ch chan Widget
	proxy Proxy
	mu sync.Mutex
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
		log.Print("Consumer: Request initiated from consumer: ", c.id)
		c.RequestWidgets()
	}
}

func (c Consumer) RequestWidgets() {
	// TODO: randomSeconds should have a random value between 1 and 5
	if len(c.ch) >= MAX_NUM_WIDGETS {
		randomSeconds := rand.Intn((maxConsumerSleep - minConsumerSleep + 1) + minConsumerSleep)
		log.Print("Consumer: putting to sleep consumer: ", c.id)
		time.Sleep(time.Duration(randomSeconds) * time.Second)
		c.flush()
	}

	c.mu.Lock()
	c.proxy.WriteWidgets(c.id, c.ch)
	c.mu.Unlock()
}

// reset consumer to initial state (using widgets)
func (c Consumer) flush() {
	for i := 0; i < MAX_NUM_WIDGETS; i++ {
		<- c.ch
		log.Print("Consumer: Recieving from channel assigned to consumer: ", c.id)
	}
}
