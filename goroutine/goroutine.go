package main

import (
	"fmt"
	"sync"
	"time"
)

func say(word string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}
func hi() {
	go say("hello")
	say("world")
}
func main() {
	//hi()
	//buffer()
	bomb()
}

func write(channel chan string) {
	channel <- "first"
	channel <- "second"
	//time.Sleep(5000 * time.Millisecond)
	channel <- "third"
}

func buffer() {
	var channel chan string = make(chan string, 6)
	/*
		channel <- "first"
			fmt.Println("send first")
			channel <- "second"
			fmt.Println("send second")
			channel <- "third"
			fmt.Println("send third")
		//go write(channel)

		a := <-channel
		fmt.Println(a)
		fmt.Println(<-channel)
		fmt.Println("waiting")
		fmt.Println(<-channel)
	*/
	go write(channel)
	go write(channel)

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			write(channel)
		}()
	}
	wg.Wait()

	close(channel)
	for i := range channel {
		fmt.Println(i)
	}
	fmt.Println("done")
}

func bomb() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	//var ronly <-chan time.Time
	//var wronly chan<- time.Time

	//ronly <-time.Tick(100 * time.Millisecond)
	//fmt.Printf(<-wronly)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
