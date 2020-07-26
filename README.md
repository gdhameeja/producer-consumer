## This is a simulation of producer and consumers


# Producer
The entity which produces a message

# Consumer
The entity which consumes a message

# Widget
The message itself

# Proxy
The interface between consumers and producers. All consumers request the proxy for widgets instead of producers. This lets the Proxy have the control over
producers and the logic for that is now maintained in a single place.

# Pointers
- main.go is the entry point for running the project.
- All entities (Producers, Consumers, Widgets) reside in the entity package
