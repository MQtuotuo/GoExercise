package main

import (
	"fmt"
	"time"
)

func callWebService(value int) {
	fmt.Println("WebService started:", value)
	time.Sleep(3 * time.Second)
	fmt.Println("WebService finished:", value)
}

var webservice = make(chan int)

func callWebService2() int {
	fmt.Println("Calling webservice")
	time.Sleep(5 * time.Second)
	fmt.Println("Webservice Finished")
	return 10
}

func showToUser() {
	fmt.Println("Showing info to User..")
}

func main() {
	//callWebService(1)
	//callWebService(2)
	//callWebService(3)

	//go callWebService(1)
	//go callWebService(2)
	//go callWebService(3)
	//time.Sleep(10 * time.Second)

	callWebService2()
	showToUser()

	//channel <-value
	result := <-webservice //receiving the value from another goroutine through the channel

	fmt.Println("Execution finished with the result:", result)
	time.Sleep(8 * time.Second)

}
