package entities

import (
	"log"
	"math/rand"
)

const (
	minNumWidgets = 1
	maxNumWidgets = 3
	maxWidgetsPerProducer = 10
)

// Interface between consumers and producers. Consumer request the proxy for
// widgets instead of directly requesting the producers.
type Proxy struct {
	// Maintains the state of each producer keeping track
	// of how many widgets a producer has produced.
	producerMap map[Producer]int
}

// create a new proxy and register all input producers with it
func InitProxy(producers... Producer) Proxy {
	proxy := Proxy{ producerMap: make(map[Producer]int) }
	for _, producer := range producers {
		proxy.producerMap[producer] = 0
	}
	return proxy
}

// get a single widget from appropriate producer
func (p Proxy) getWidget() Widget {
	producer := p.getProducer()
	widget := producer.ProvideWidget()
	p.producerMap[producer]++
	return widget
}

// decide which producer is going to be used according to their capacity
// goes over each producer in round robin manner
func (p Proxy) getProducer() Producer {
	for producer, numWidgets := range p.producerMap {
		if numWidgets < maxWidgetsPerProducer {
			return producer
		} else {
			// flush producer, this algorithm works for only 2 producers, since
			// if we flush while checking, the third producer will never get used
			p.producerMap[producer] = 0
		}
	}
	// this will never go into infinite loop because we're flushing producers as we go
	return p.getProducer()
}

// the main entry point that each consumer will call to receive widgets.
// connId is the id of the consumer. It is just used for logging purposes.
func (p Proxy) WriteWidgets(conId int, in chan <- Widget) {
	numWidgets := rand.Intn((maxNumWidgets - minNumWidgets) + minNumWidgets)
	for i := 0; i < numWidgets; i++ {
		widget := p.getWidget()
		in <- widget
		log.Print("Wdiget successfully writen to channel: ", conId)
	}
}
