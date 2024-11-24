//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Syed Muhammad Umar (C00278724@setu.ie)
// Description: This file is modified from previous version of the same file which was using a simple barrier. This code is modified to make our barier reusable. The logic is to create a barrier with two gates and synchronize them in such a way using channels that when one opens the other one closes and vice versa. Some variable and function names are changed to make it easy to explain.
// A simple reusable barrier implemented using mutex and channels
// Issues:
// None I hope
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Barrier is a reusable barrier struct.
type Barrier struct {
	currentCount int        // Tracks the number of goroutines that have reached the barrier
	totalCount   int        // Total number of goroutines required to release the barrier
	lock         sync.Mutex // Protects the currentCount
	beforeGate   chan bool  // First gate channel for synchronization
	afterGate    chan bool  // Second gate channel for synchronization
}

// NewBarrier initializes a reusable barrier for the specified number of goroutines.
func NewBarrier(totalCount int) *Barrier {
	b := &Barrier{
		totalCount: totalCount,
		beforeGate: make(chan bool, 1),
		afterGate:  make(chan bool, 1),
	}
	b.afterGate <- true // Initially, the after gate is closed, and the first gate is open
	return b
}

// WaitBefore blocks goroutines until all have reached the barrier (first stage).
func (b *Barrier) WaitBefore() {
	b.lock.Lock()
	b.currentCount++
	if b.currentCount == b.totalCount {
		<-b.afterGate        // Close the after gate
		b.beforeGate <- true // Open the before gate
	}
	b.lock.Unlock()
	<-b.beforeGate // Wait at the before gate
	b.beforeGate <- true
}

// WaitAfter blocks goroutines until all have left the barrier (second stage).
func (b *Barrier) WaitAfter() {
	b.lock.Lock()
	b.currentCount--
	if b.currentCount == 0 {
		<-b.beforeGate      // Close the before gate
		b.afterGate <- true // Open the after gate
	}
	b.lock.Unlock()
	<-b.afterGate // Wait at the after gate
	b.afterGate <- true
}

// doWork simulates work divided into two parts with a reusable barrier in between.
func doWork(workerID int, barrier *Barrier, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d: Part A completed\n", workerID)

	barrier.WaitBefore() // Wait at the first barrier
	fmt.Printf("Worker %d: Passed Barrier\n", workerID)

	time.Sleep(time.Second) // Simulate work after the barrier
	fmt.Printf("Worker %d: Part B completed\n", workerID)

	barrier.WaitAfter() // Wait at the second barrier
}

func main() {
	totalWorkers := 10
	var wg sync.WaitGroup
	wg.Add(totalWorkers)

	barrier := NewBarrier(totalWorkers) // Create a reusable barrier

	for i := 0; i < totalWorkers; i++ {
		go doWork(i, barrier, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers have finished.")
}
