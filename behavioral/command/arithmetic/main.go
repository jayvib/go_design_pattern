package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Calculator interface {
	Calculate(x, y int) int
}

type Appender interface {
	Append(container *Container)
}

type Container struct {
	x, y       int
	calculator Calculator
	resultChan chan int
}

type ContainerQueue struct {
	mu         sync.Mutex
	containers []*Container
}

func (c *ContainerQueue) Pop() *Container {
	c.mu.Lock()
	defer c.mu.Unlock()
	var cont *Container
	cont, c.containers = c.containers[len(c.containers)-1], c.containers[:len(c.containers)-2]
	return cont
}

func (c *ContainerQueue) Append(container *Container) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.containers = append(c.containers, container)
}

func (c *ContainerQueue) Do() {
	for _, container := range c.containers {
		fmt.Printf("%d and %d: %d\n", container.x, container.y, container.calculator.Calculate(container.x, container.y))
	}
}

func containerProducer(ctx context.Context, wg *sync.WaitGroup, appender Appender, fn func(int, int) int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println("Appending")
		time.Sleep(time.Duration(rand.Intn(600)) * time.Millisecond)
		select {
		case <-ctx.Done():
			return
		default:
		}
		newContainer := &Container{
			x:          rand.Int(),
			y:          rand.Int(),
			calculator: calculatorFunc(fn),
		}
		appender.Append(newContainer)
	}
}

type calculatorFunc func(x, y int) int

func (p calculatorFunc) Calculate(x, y int) int {
	return p(x, y)
}

func containerGenerator(ctx context.Context, fn func(int, int) int) <-chan *Container {
	containerChan := make(chan *Container)
	rand.Seed(time.Now().UnixNano())
	go func() {
		defer close(containerChan)
		for {
			n1 := rand.Intn(100)
			n2 := rand.Intn(500)
			resultChan := make(chan int)
			newContainer := &Container{
				x:          n1,
				y:          n2,
				calculator: calculatorFunc(fn),
				resultChan: resultChan,
			}
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			case containerChan <- newContainer:
				fmt.Println("Container sent")
			}

			res := <-resultChan
			fmt.Println("Waiting for the result...")
			fmt.Println("Received result:", res)
			fmt.Println("-------------------")
		}
	}()
	return containerChan
}

func process(ch <-chan *Container) {
	for c := range ch {
		time.Sleep(1 * time.Second)
		c.resultChan <- c.calculator.Calculate(c.x, c.y)
	}
}

func main() {
	var wg sync.WaitGroup
	prodCal := func(x, y int) int {
		return x * y
	}

	addCal := func(x, y int) int {
		return x + y
	}

	SubCal := func(x, y int) int {
		return x - y
	}

	cq := new(ContainerQueue)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()

	wg.Add(3)
	go containerProducer(ctx, &wg, cq, prodCal)
	go containerProducer(ctx, &wg, cq, addCal)
	go containerProducer(ctx, &wg, cq, SubCal)

	fmt.Println("Waiting the children to finish...")
	wg.Wait()

	cq.Do()
}
