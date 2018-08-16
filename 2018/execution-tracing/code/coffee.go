func callWork() {
	orderID := strconv.Itoa(rand.Int())
	ctx, task := trace.NewTask(context.Background(), "makeCappuccino")
	trace.Log(ctx, "orderID", orderID)
	milk := make(chan bool)
	espresso := make(chan bool)
	go func() {
		trace.WithRegion(ctx, "steamMilk", func() {
			trace.Log(ctx, "makingMilk", "starting steam with params")
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