package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Number of product in a Batch
const maxProduct int = 10

// Product Successfully Made, Product Failed, Total Product produced
var sucProduct, failProduct int // totalP int

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
	quit  chan chan error
}

// Stop production signal
func (o *serviceHost) Close() error {
	ch := make(chan error)
	
	o.quit <- ch
	return <-ch
}

// Sending Order to Factory
func production(incomingOrder *serviceHost) {

	// Create a variable to keep track of the incoming order
	var orderNo int

	// Send the incoming order to the factory for procuction
	// We get the output of the production
	for {

		output := produce(orderNo)

		select {

		// Send the product back to client(MAIN FUNC) via a channel
		case incomingOrder.order <- *output:

		case quitCh := <-incomingOrder.quit:
			close(incomingOrder.order)
			close(quitCh)
			return
		}
	}

}

func produce(orderNumber int) *orderStat {

	var exitMesg string
	success := false

	// Increment the order Number
	orderNumber++

	// Take a moment to produce
	processingTime := rand.Intn(7) + 1

	// Inform client we have received order
	color.White("We have recieved the Order #%d\nYour order will take %d seconds to be processed, kindly exercise patience", orderNumber, processingTime)

	// Begin producing based on Probably
	time.Sleep(time.Duration(processingTime) * time.Second)

	temp := rand.Intn(12) + 1

	switch {
	case temp < 2:
		failProduct++
		exitMesg = fmt.Sprintln("Apologies, we are unable to produce as the temperature is too low for the machines to work")
	case temp < 5:
		failProduct++
		exitMesg = fmt.Sprintln("We are experience some idle time here.")
	case temp < 9:
		sucProduct++
		exitMesg = fmt.Sprintln("Successfully made the product.")
		success = true
	case temp <= 12:
		failProduct++
		exitMesg = fmt.Sprintln("Too hot for worker to work.")

	}

	p := orderStat{
		batchNo: orderNumber,
		Memo:    exitMesg,
		suc:     success,
	}

	return &p
}

// Function to call when the temperature is too low to work
//accepts an argument of the orderNo

func main() {

	// Seed to generate a random number
	rand.Seed(time.Now().UnixNano())

	// Print Starting text
	color.Cyan("We are Open for business\nPlease send in your Orders....")

	// Init Service host
	factory := serviceHost{
		order: make(chan orderStat),
		quit:  make(chan chan error),
	}

	// Produce
	go production(&factory)

	// Test Production
	for i := range factory.order {
		if i.batchNo <= maxProduct {
			if i.suc {
				color.Green(i.Memo)
				color.Green("Order #%d was produced successfully", i.batchNo)
			} else {
				color.Red(i.Memo)
				color.Red("What a terrible day")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := factory.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	// Ending Test

}
