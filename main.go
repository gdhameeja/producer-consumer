// solution approach, producers have a proxy in front of them and all consumers ask for widgets from this proxy
// consumers also hand over the `channel` on which they want to receive widgets. The proxy maintains the status of each 
// producer and assigns consumer to an appropriate prodoucer. The producer then creates widgets and sends on the channel
// given to it by the proxy (which originally came from consumer).
// Q. The proxy can maintain the status of each producer checking how many widgets each producer holds with them (honouring the each producer can have 10 widgets only limit).
// Q. Should we have a queue between consumers and producer proxy to make the widgets transaction easier? In this case shuffling between producers will become much easier.
// We might not need channels as well, if we use the queue. But we won't to showcase our go skills, not ds skills so we want to use channels.
package main

import (
	"entities"

const (
	numProducers = 4
	numConsumers = 2
)


// main entry point of the program
func main() {
	// create the producers, consumers

	// launch the producers and consumers in action forever

}
