// solution approach, producers have a proxy in front of them and all consumers ask for widgets from this proxy
// consumers also hand over the `channel` on which they want to receive widgets. The proxy maintains the status of each 
// producer and assigns consumer to an appropriate prodoucer. The producer then creates widgets and sends on the channel
// given to it by the proxy (which originally came from consumer).
// Q. The proxy can maintain the status of each producer checking how many widgets each producer holds with them (honouring the each producer can have 10 widgets only limit).
// Q. Should we have a queue between consumers and producer proxy to make the widgets transaction easier? In this case shuffling between producers will become much easier.
// We might not need channels as well, if we use the queue. But we won't to showcase our go skills, not ds skills so we want to use channels.
// Also the proxy can flush the producer when a producer reaches it's max limit of number of widgets it can produce. flush just means to reset the widgetCounter to 0
// NOTE: remember to use buffered channels as transport. If we use unbuffered channels, the producer will get blocked, until, the consumer accepts the widget, and won't be
// able to cater to other consumers. Whereeas buffered channels will only block if buffered channel is full and producer tries to send something or buffered channel is empty
// and consumer tries to receive something.
// maybe have a channel per consumer so we don't get channels mixed up between consumers
// producer does not need to run inside a goroutine because the producer function is not conitnous, go routines are meant to run functions not *entities*
package main

import (
	"gdhameeja/producerconsumer/entities"
)

const (
	numProducers = 4
	numConsumers = 2
)


// main entry point of the program
func main() {
	// create the producers 
	producer1 := new(entities.Producer)
	producer2 := new(entities.Producer)

	// create Proxy
	proxy := entities.InitProxy(*producer1, *producer2)

	// create consumers
	consumer1 := entities.ConstructConsumer(1, proxy)
	consumer2 := entities.ConstructConsumer(2, proxy)
	consumer3 := entities.ConstructConsumer(3, proxy)

	// exitChannel is the channel which blocks the main goroutine from exiting, since it is never written to
	exitChannel := make(chan bool)	
	// launch the consumers in action forever
	launchConsumers(exitChannel, consumer1, consumer2, consumer3)
	<- exitChannel
	close(exitChannel)
}


func launchConsumers(exit chan <- bool, consumers... entities.Consumer) {
	for _, consumer := range consumers {
		go consumer.Start()
	}
}
