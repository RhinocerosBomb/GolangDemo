package main

import (
	"fmt"
	"sync"
	"time"
)

type Cancelable struct {
	doSomething func()
}

func main() {
	c := Cancelable{doSomething: func() {
		fmt.Println("did something")
	}}

	cancel := c.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(2*time.Minute, func() {
		cancel()
		wg.Done()
	})

	wg.Wait()
	time.Sleep(10*time.Minute)
	cancel()
}

func (c *Cancelable) Start() (cancelFunc func()) {
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- time.After(2*time.Second):
				c.doSomething()
			case <- quit:
				return
			}
		}
	}()

	return func() {
		quit <- struct{}{}
	}
}
