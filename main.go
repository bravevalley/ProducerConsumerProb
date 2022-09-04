package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Number of product in a Batch
// const maxProduct int = 10


// // Product Successfully Made, Product Failed, Total Product produced
// var sucProduct, failProduct, totalP int

// Product Stat
type orderStat struct {
	batchNo int
	Memo    string
	suc     bool
}

// Go concurrency channel to communicate
// between Main routine and Goroutine
type serviceHost struct {
	order chan orderStat
	quit chan chan error
}

// Sending Order to Factory
func production(output *serviceHost) {
	var orderNo int

	produce(orderNo)

}

func produce(orderNumber int) {
	// Increment the order Number

	orderNumber++

	// Take a moment to produce
	processingTime := rand.Intn(7) * 1

	// Inform client we have received order
	color.White("We have recieved the Order #%d\nYour order will take %d seconds to be processed, kindly exercise patience", orderNumber, processingTime)


	// Begin producing based on Probably
	time.Sleep(time.Duration(processingTime)  * time.Second)

	fmt.Println("I am still working...")

	wg.Done()

}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// Seed to generate a random number
	rand.Seed(time.Now().UnixNano())

	// Print Starting text
	color.Cyan("We are Open for business\nPlease send in your Orders....")

	// Init Service host
	factory := serviceHost{
		order: make(chan orderStat),
		quit: make(chan chan error),
	}

	// Produce
	go production(&factory)

	// Test Production

	// Ending Test
	wg.Wait()
}
