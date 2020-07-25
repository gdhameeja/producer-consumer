package entities

import "log"

// Interface between consumers and producers. Consumer request the proxy for
// widgets instead of directly requesting the producers.
type Proxy struct {
	// Maintains the state of each producer keeping track
	// of how many widgets a producer has produced.
	producerMap map[Producer]int

}

// create a new proxy and register all input producers with it
func InitProxy(producers... Producer) Proxy {
	proxy := new(Proxy)
	for _, producer := range producers {
		proxy.producerMap[producer] = 0
	}
	return *proxy
}

// get a single widget from appropriate producer
func (p Proxy) getWidget() (Widget, error) {
	return Widget{}, nil
}

// when a producer will create a widget, the widget count for the producer will be
// incremented using this so that proxy can maintain state of producers
func (p Proxy) incrementMessageForProducer(producer Producer) {
}

// the main entry point that each consumer will call to receive widgets.
// connId is the id of the consumer. It is just used for logging purposes.
func (p Proxy) WriteWidgets(numWidgets int, conId int, in chan <- Widget) {
	for i := 0; i < numWidgets; i++ {
		widget, err := p.getWidget()
		if err != nil {
			log.Print("Error while creating widget: ", err)
			return
		}
		in <- widget
		log.Print("Wdiget successfully writen to channel: ", conId)
	}
}
