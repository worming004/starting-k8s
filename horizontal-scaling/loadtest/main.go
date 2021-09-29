package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type worker struct {
	address string
}

func (w worker) work() {
	for {
		_, err := http.Get(w.address)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

var address = flag.String("address", "http://20.82.159.140:8000/fib/45", "address to call")
var nbr = flag.Int("number", 10, "number of worker to create")

func main() {
	flag.Parse()
	for i := 0; i < *nbr; i++ {
		w := worker{address: *address}
		go w.work()
	}

	ossig := make(chan os.Signal, 1)
	signal.Notify(ossig, syscall.SIGINT, syscall.SIGTERM)
	v := <-ossig
	fmt.Println(v)
}
