package main

import (
	"fmt"
	"time"
)

// counter is variable increament by each goroutines.
var counter int

/*
race conditon is unsynchronized access to a shared resource and attempt to
read and write to that resource at the same time by more goroutines.

race condition occurs when Each goroutine overwrites the work of the other.
This happens when the goroutine swap is taking place.
Each goroutine makes its own copy of the variable and
then is swapped out for the other goroutine.
When the goroutine is given time to execute again,
the value of the variable has changed,
but the goroutine doesn’t update its copy.
Instead it continues to change the copy it has and set
the value back to the variable,
replacing the work the other goroutine performed
*/
func main() {
	loops := 0
	for loops < 2 {
		go func() {
			counter++
		}()
		loops++
	}
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("[X] %v\n", counter)
}
