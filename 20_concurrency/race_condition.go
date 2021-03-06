package main

import (
	"fmt"
	"runtime"
	"sync"
)


var counter int
var wg sync.WaitGroup


func incrementer(){
	wg.Add(1)
	i := counter
	runtime.Gosched()
	i++
	counter = i
	fmt.Println("counter: ",counter,"NumOfGoRountine: ",runtime.NumGoroutine())
	wg.Done()
}

func main(){
	counter = 0
	for j := 0; j<100;j++{
		go incrementer()

	}
	wg.Wait()
	fmt.Println("Final counter value: ",counter)

}
