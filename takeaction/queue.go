package takeaction

import (
	"fmt"
)

//Task 1: Write a simple queue
//Your first task is to write a simple first-in-first-out (FIFO) container, also called a queue.
//Below is some stub code to complete. Replace all TODO labels with your code.
//Tip: A slice comes handy as the basis for a queue. Do you remember how to append a value to a slice, and how to remove the first element of a slice?

// TODO: define the Queue type
type Queue []interface{}


// PutAny adds an element to the queue.
func (c *Queue) PutAny(elem interface{}) {
	//TODO: append elem to the queue
	*c = append(*c, elem)
}

// GetAny removes an element from the queue.
// If the queue is empty, GetAny returns an error.
func (c *Queue) GetAny() (interface{}, error) {
	//TODO: fetch the first element's value, and then remove the first element from the queue.
	//	If the queue is already empty, return the zero value of interface{} and an error.

	if len(*c) == 0 {
		return nil, fmt.Errorf("Queue empty nothing to return")
	}

	e := (*c)[0]
	*c = (*c)[1:]
	return e, nil
}

func TestDriveGenericQueue() {
	var q Queue //FIFO
	q.PutAny("apple")
	q.PutAny("peach")
	q.PutAny("orange")

	for i:=0; i<4; i++ {
		if c, err := q.GetAny(); err == nil {
			fmt.Println(c)
		}
	}
}



//Task 2: Write a wrapper to gain run-time type safety
//
//Imagine you need a queue of integers.
//Create a type IntQueue based on the generic Queue type, and implement Put and Get methods that append an int to the queue and return an int from the queue, respectively.

type IntQueue struct {
	q Queue
}

func (ic *IntQueue) Put(n int) {
	(*ic).q.PutAny(n)
}

func (ic *IntQueue) Get() (int, error) {
	v, err := (*ic).q.GetAny()

	if err != nil {
		return 0, err
	}

	if i, ok := v.(int); ok {
		return i, err
	}

	return 0, fmt.Errorf("Failed to type assert interface{} to int")
}

func TestDriveIntQueue() {
	ic := IntQueue{}
	ic.Put(7)
	ic.Put(42)

	for i := 0; i < 3; i++ {
		elem, err := ic.Get()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Got: %d (%[1]T)\n", elem)
	}
}