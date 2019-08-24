package main

func main() {
	src := make(chan int) // Create a channel that only passes ints
	go Generate(src)      // Launch Generate goroutine
	for i := 0; i < 100; i++ {
		prime := <-src
		println(prime)
		dst := make(chan int)
		go Filter(src, dst, prime)
		src = dst
	}
}

// Generate generate prime numbers
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'
	}
}

// Filter filter prime numbers
func Filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // loop over values received
		if i%prime != 0 {
			dst <- i // Send 'i' to the channel 'dst'
		}
	}
}
