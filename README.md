# MP3

Understanding the performance of approximate consensus and analysis performance

## To Run

### Send messages between different clients

Open your terminal, run

```
go run main.go
```

Then you will see a log appears (wait for a few seconds) something like this

```
Process 5 has value 0.414 at round 1 
Process 1 has value 0.514 at round 1 
Process 6 has value 0.407 at round 1 
Process 4 has value 0.414 at round 1 
Process 7 has value 0.407 at round 1 
Process 8 has value 0.407 at round 1 
```

When approximate consensus achieved, it will shows

```
2020/10/12 23:43:33 consensus achieved using 1m14.440437813s
```

## Structure and Design

### Application

In `/application/interface.go`, there are two struct:

```
type Process struct {
	Id    string
	Ip    string
	Port  string
	Ln    net.Listener
	Conns []net.Conn
	InitialV float64
}

type Message struct {
	S string
	R int
	M float64
}
```

Basicly, each process represents a node in approx. consensus algorithm. It has six given properties, the first three are very straight forward, and Ln represents the net.Listener when we created a listener for that process, and Conns stores the connection to all other processes. InitialV is the given first input.

`Message` has three components, `S` is the source process id, `R` is the round in the tuple, `M` is the value in the tuple.

### Network

The `/network/connection.go` contains several functions, `Initialize` function intialize the connection among all the processes. I reimplemented this for MP1. (Since I didn't understand how to do it several weeks ago.) `ApproximateConsensus` function simulates the algorithm we learned in the class. It will terminate when the value of two rounds are within 0.001.

The `/network/msg.go` is similar to the MP1. I modified a little bit based on Colin's group's idea.

### Main
The main function uses `/util/utils.go` to fetch data from the config file. Then it will call the `initialize` function to initialize the connection, and then each process will has its own goroutine for receiving messages. They will update value each round and stop when consensus has achieved. I implemented this feature by adding the control channel, when nonfailure node has achieved consensus, the main process will stop.

### Bottleneck and acknowledgement
I encountered some problems with coding, and I used Colin's group's design to solve the problem since I couldn't figure out what was wrong and how to debug the error. It was something related to gorountine.

I also encountered some problems with `gob`. It gave `gob: type not registered for interface:` error, and I couldn't figure this out. I asked Roger Pan for some help, and found out that I shouldn't pack net.Listner into `gob` because it cannot do so. 

I didn't implement the random faliure node feature. I am not sure how to do so.

## Plot analysis

I put several plots in the `/plots` directory, and the data can be found in `/datas` directory. I did this manually. I learned the shell script by reading the tutorial, but wasn't sure how to do so.

The first analysis is the input size. Given same max and min delay time and similar inputs. From the plots, the consensus achieved time increases, but the rounds is almost the same. It is reasonable since each commincation costs some time, but since we are taking the average number, they converges to a number quickly.

The second analysis is the failure node. Given other conditions same, I set 9 processes. The time achived consensus and rounds look not vary a lot. I am not sure whether this result makes sense, because I didn't implement the failure node feature.

The third analysis is about mindelay time. Given other conditions same, I set 7 processes, max delay is 3000 and failure node is 1. The min delay time ranges from 0 to 2500. We can see that rounds still don't change much, but the time achieved consensus increased as minDelay gets longer.

The final analysis is about maxdelay time. Given other conditions same, I set 7 processes, min delay to be 1000 and failure node is 1. The max delay time ranges from 3000 to 5000. From the plot, we can see that rounds don't change much, but the time achieved consensus increased.

## Conclusion

Overall, the input size and delay time affect a lot toward the consensus time. I guess it should go higher very quickly because when there are more nodes, each node implies a lot more connections. If delay time is also high, then the time will increases significantly.

The rounds, however, doesn't seem to change a lot. In my tests, it achieves consensus at around round 3-4. I guess this is becasue the input is limited to real number between 0 and 1, and since we are taking the average, the consensus will be achieved after a small amount of rounds.

In the future, if I would do similar tests, I think it is necessary to learn how to use shell script to do the test. I didn't find an example of how to use shell script to do the test. In this MP, I manually change my config file and run the program, then recording them to a csv file. Then I use python script to plot the graph.

## Source
* [Channels and Go Routines](https://www.justindfuller.com/2020/01/go-things-i-love-channels-and-goroutines/)
* [Create a TCP and UDP Client and Server using Go](https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/)
* [Go Routines](https://golangbot.com/goroutines/)
* [Net Package](https://golang.org/pkg/net/)
* [How to stop a Go Routine](https://stackoverflow.com/questions/6807590/how-to-stop-a-goroutine/6807784#6807784)
* [Colin's Group](https://github.com/jzulewski/MP3)


## Authors

* **Gary Ge** - *initial work*
