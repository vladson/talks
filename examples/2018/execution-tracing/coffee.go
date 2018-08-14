package main

import (
	"context"
	"fmt"
	"runtime/trace"
	"time"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"log"
	"strconv"
)

func main() {
	http.Handle("/coffee", http.HandlerFunc(helloHandler))

	log.Println(http.ListenAndServe("localhost:8181", http.DefaultServeMux),nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	callWork()
	w.Write([]byte("Your awesome cold brew!"))
}

func callWork() {

	orderID := strconv.Itoa(rand.Int())

	ctx, task := trace.NewTask(context.Background(), "makeCappuccino")
	trace.Log(ctx, "orderID", orderID)

	milk := make(chan bool)
	espresso := make(chan bool)

	go func() {
		trace.WithRegion(ctx, "steamMilk", func() {
			trace.Log(ctx, "makingMilk", "starting steam with params")
			if rand.Float32() > 0.6 {
				time.Sleep(2000*time.Millisecond)
				trace.Log(ctx, "makingMilk", "Boiler failure")
				fmt.Println("Bang")
			} else {
				time.Sleep(50*time.Millisecond)
			}
			trace.Log(ctx, "makingMilk", "steam finished with results")
		})
		milk <- true
	}()
	go func() {
		trace.WithRegion(ctx, "extractCoffee", extractCoffee)
		espresso <- true
	}()
	go func() {
		<-espresso
		<-milk
		trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
		defer task.End() // When assemble is done, the order is complete.
	}()
}

func extractCoffee() {
	time.Sleep(142*time.Millisecond)
}

func mixMilkCoffee() {
	time.Sleep(142*time.Millisecond)
}