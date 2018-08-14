package main

import (
	"fmt"
	"time"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"log"
)

func main() {
	http.Handle("/coffee", http.HandlerFunc(oldCoffeeHandler))

	log.Println(http.ListenAndServe("localhost:8180", http.DefaultServeMux),nil)
}

func oldCoffeeHandler(w http.ResponseWriter, r *http.Request) {
	oldCallWork()
	w.Write([]byte("Your awesome cold brew!"))
}

func oldCallWork() {

	milk := make(chan bool)
	espresso := make(chan bool)

	go func() {
		steamMilk()
		milk <- true
	}()
	go func() {
		extractCoffee1()
		espresso <- true
	}()
	go func() {
		<-espresso
		<-milk
		mixMilkCoffee1()
	}()
}

func steamMilk() {
	if rand.Float32() > 0.6 {
		time.Sleep(2000*time.Millisecond)
		fmt.Println("Bang")
	} else {
		time.Sleep(50*time.Millisecond)
	}
}


func extractCoffee1() {
	time.Sleep(142*time.Millisecond)
}

func mixMilkCoffee1() {
	time.Sleep(142*time.Millisecond)
}