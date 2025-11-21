package main

import (
	// "flag"
	"fmt"
)

// min               run all tests/requests recursively in order
// min -p            run tests/requests in parallel
// min -w            watch for files changes (run tests/requests on change)
// min -v            verbose (displays full output on failure AND success)
// min <file.json>   run the test/request of a single file
// min <folder>      run all tests/requests recursively from said directory

func main() {
	sum := 0
	for range 10 {
		sum += 1
		run_all_tests("gabe")
	}
	fmt.Printf("")
	fmt.Printf("")
	fmt.Println("Hello world!", sum, "I love you!")
}

// min
func run_all_tests(x string) {
	fmt.Println("Running all tests in order!")
}

// min -p
func run_all_tests_parallel() {
	fmt.Println("Running all tests in parallel!")
}
