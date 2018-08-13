package main

import (
	"context"
	"fmt"
	"runtime/trace"
	"os"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	orderID := "123e4567-e89b-12d3-a456-426655440000"

	ctx, task := trace.NewTask(context.Background(), "makeCappuccino")
	trace.Log(ctx, "orderID", orderID)

	milk := make(chan bool)
	espresso := make(chan bool)

	go func() {
		trace.WithRegion(ctx, "steamMilk", steamMilk)
		milk <- true
	}()
	go func() {
		trace.WithRegion(ctx, "extractCoffee", extractCoffee)
		espresso <- true
	}()
	go func() {
		defer task.End() // When assemble is done, the order is complete.
		<-espresso
		<-milk
		trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
	}()
}

func steamMilk() {
	trace.Log(context.Background(), "BeforeSteam", "starting steam with params")
	fmt.Println("Pshshshsh")
	trace.Log(context.Background(), "AfterSteam", "steam finished with results")
}

func extractCoffee() {
	fmt.Println("Bubblbe-buble zzzzzz")
}

func mixMilkCoffee() {
	fmt.Println("wirple")
}