package checkups

import "fmt"

type ChainCheckup interface {
	Next(name string)
}

type Decorator func(ChainCheckup) ChainCheckup

type DecorateFunc func(name string)

func (d DecorateFunc) Next(name string) {
	d(name)
}

type IDCheckup struct {
	NextChain ChainCheckup
}

func (c *IDCheckup) Next(name string) {
	fmt.Printf("Checking the ID of %s\n", name)
	if c.NextChain != nil {
		c.NextChain.Next(name)
	}
}

type HealthCheckup struct {
	NextChain ChainCheckup
}

func (c *HealthCheckup) Next(name string) {
	fmt.Printf("Checking the ID of %s\n", name)
	if c.NextChain != nil {
		c.NextChain.Next(name)
	}
}

func Decorate(c ChainCheckup, decorators ...Decorator) ChainCheckup {
	decorated := c
	for _, decorator := range decorators {
		decorated = decorator(decorated)
	}
	return decorated
}

func IDCheckupFunc() Decorator {
	return func(chain ChainCheckup) ChainCheckup {
		return DecorateFunc(func(name string) {
			fmt.Printf("Checking the ID of %s\n", name)
			chain.Next(name)
		})
	}
}

func HealthCheckupFunc() Decorator {
	return func(chain ChainCheckup) ChainCheckup {
		return DecorateFunc(func(name string) {
			fmt.Printf("Checking the Health of %s\n", name)
			chain.Next(name)
		})
	}
}

func OralCheckupFunc(checker string) Decorator {
	return func(chain ChainCheckup) ChainCheckup {
		return DecorateFunc(func(name string) {
			fmt.Printf("Checking the mouth of %s with %s\n", name, checker)
			chain.Next(name)
		})
	}
}
