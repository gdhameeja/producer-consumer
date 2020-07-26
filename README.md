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

# RUN
`$ go run main.go`
# Output
Few lines of output (Note: not in order. Lines were deleted in between to accomodate the intresting parts of output")
$ go run main.go
{0}
```
2020/07/26 21:47:06 Consumer: Request initiated from consumer: 1
2020/07/26 21:47:06 Proxy: received request to provide widget to consumer: 1
2020/07/26 21:47:06 Proxy: number of widgets to be generated: 1
2020/07/26 21:47:06 Proxy: current number of widgets in channel for consumer: 1 0
2020/07/26 21:47:06 Producer: widget created by producer: 1
020/07/26 21:47:06 Wdiget successfully writen to channel: 2
2020/07/26 21:47:06 Producer: widget created by producer: 1
2020/07/26 21:47:06 Wdiget successfully writen to channel: 2
2020/07/26 21:47:06 Producer: widget created by producer: 1
2020/07/26 21:47:06 Wdiget successfully writen to channel: 2
2020/07/26 21:47:06 Consumer: Request initiated from consumer: 2
020/07/26 21:47:06 Producer: widget created by producer: 1
2020/07/26 21:47:06 Wdiget successfully writen to channel: 3
2020/07/26 21:47:06 Consumer: Request initiated from consumer: 3
2020/07/26 21:47:06 Consumer: putting to sleep consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Consumer: Recieving from channel assigned to consumer: 3
2020/07/26 21:47:06 Proxy: received request to provide widget to consumer: 3
```

