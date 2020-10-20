package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type chopStick struct {
	sync.Mutex
}

type philo struct {
	leftCS, rightCS *chopStick
	id              int
	eats            int
	finishEat       chan int
	startEat        chan int
	reqEat          chan int
}

func (p philo) eat() {
	for {
		// check goroutine how many time is eating
		if p.eats == 3 {
			p.finishEat <- p.id
			gw.Done()
			return
		}
		// send start eating request to host goroutine
		p.reqEat <- p.id
		// block on startEat chanel if a value is (1) start eating else wait in a line.
		if eat := <-p.startEat; eat == 1 {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("[#] starting to eat %d\n", p.id)
			// random time to simulate eating time.
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(p.id*1000)))
			fmt.Printf("[#] finishing eating %d\n", p.id)
			p.eats++
			p.finishEat <- p.id
			p.leftCS.Unlock()
			p.rightCS.Unlock()
		}
	}
}

var gw sync.WaitGroup

func init() {
	// Random seed
	rand.Seed(time.Now().UTC().UnixNano())
}
func main() {
	// initialization section
	gw.Add(5)
	CSticks := make([]*chopStick, 5)
	philos := make([]*philo, 5)
	finishEat := make(chan int)
	startEat := make(chan int)
	requestEat := make(chan int)
	quitConf := make(chan int)
	quit := make(chan int)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(chopStick)
	}
	for i := 0; i < 5; i++ {
		philos[i] = &philo{
			leftCS:    CSticks[i],
			rightCS:   CSticks[(i+1)%5],
			id:        i + 1,
			eats:      0,
			finishEat: finishEat,
			startEat:  startEat,
			reqEat:    requestEat,
		}
	}

	// generate goroutines section
	go host(requestEat, startEat, finishEat, quit, quitConf)
	for _, v := range philos {
		go v.eat()
	}
	// block here to let philos finish thier eating.
	gw.Wait()
	// send quit sign to host function to finish it's work gracefully
	quit <- 1
	<-quitConf
}

// host function use as a manager to recieve request eating from goroutines and check
// if there is a room in it's list and than send sign to goroutine to start eating or
// waiting based on host list. also is wait for finish eating and start eating sign from.
// goroutines
// and check if there are no more than 2 goroutine run concurrently
func host(reqEat, startEat, finishEat, quit, quitConf chan int) {
	clientsID := make(map[int]string)
	defer func() {
		quitConf <- 1
	}()
	for {
		select {
		case id := <-reqEat:
			if len(clientsID) < 2 {
				clientsID[id] = "Eating"
				startEat <- 1
			} else {
				startEat <- 0
			}
		case id := <-finishEat:
			delete(clientsID, id)
		case <-quit:
			return
		}
	}

}
