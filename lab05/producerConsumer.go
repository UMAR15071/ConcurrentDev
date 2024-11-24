//name:				Syed Muhammad Umar
//Student Number 	C00278724

package main

import (
	"fmt"
	"sync"
	"time"
)

const Timeout = 100 // Define a timeout value for the consumer (in milliseconds)

func main() {
	// Number of producers and consumers are adjustable
	multiConsumerProducer(2, 2)
}

// multiConsumerProducer creates multiple producers and consumers
func multiConsumerProducer(producerSize, consumerSize int) {
	// Create a channel for string messages
	ch := make(chan string)
	var wg sync.WaitGroup

	// Start multiple producers
	for i := 0; i < producerSize; i++ {
		wg.Add(1) // Increment the wait group counter for each producer
		go producer(i, ch, &wg)
	}

	// Start multiple consumers
	for i := 0; i < consumerSize; i++ {
		wg.Add(1) // Increment the wait group counter for each consumer
		go consumer(i, ch, &wg)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()
	close(ch) // Close the channel when all tasks are done
}

// producer function simulates the production of items
func producer(index int, ch chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		// Send a message to the channel
		ch <- fmt.Sprintf("Producer %v send %v", index, i)
	}
	wg.Done() // Mark this producer as done
}

// consumer function simulates consumption of items from the channel
func consumer(index int, ch chan string, wg *sync.WaitGroup) {
	for {
		select {
		case msg, ok := <-ch:
			// Check if the channel is closed and all messages are consumed
			if !ok {
				wg.Done() // If the channel is closed, mark this consumer as done
				return
			}
			// Process the message from the channel
			fmt.Printf("Consumer %v Received: %s\n", index, msg)

		case <-time.After(Timeout * time.Millisecond):
			// If the consumer is idle for the given timeout, stop waiting
			wg.Done() // Decrement the wait group counter, allowing the program to terminate gracefully
			return
		}
	}
}
