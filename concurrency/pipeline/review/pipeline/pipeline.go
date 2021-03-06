package pipeline

func LaunchPipeline(amount int) int {
	return <-sum(power(generator(amount)))
}

func generator(max int) <-chan int {
	outChInt := make(chan int, 100)
	go func() {
		defer close(outChInt)
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
	}()
	return outChInt
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		defer close(out)
		for i := range in {
			i *= 2
			out <- i
		}
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		defer close(out)
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
	}()
	return out
}
