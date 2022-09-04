package main

import (
	"math/rand"
	"time"

	"github/fatih/color"

	"github.com/fatih/color"
)

// Number of product in a Batch
const maxProduct int = 10


// Product Successfully Made, Product Failed, Total Product produced
var sucProduct, failProduct, totalP int

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


func main() {

	// Seed to generate a random number
	rand.Seed(time.Now().UnixNano())

	// Print Starting text
	color.Cyan("We are Open for business\nPlease send in your Orders....")

	// Produce

	// Test Production

	// Ending Test
}
