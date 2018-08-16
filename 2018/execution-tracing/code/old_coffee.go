func callWork() {
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
	} else {
		time.Sleep(50*time.Millisecond)
	}
}