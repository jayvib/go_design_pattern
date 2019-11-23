package main

var (
	addCh      = make(chan bool)
	getCountCh = make(chan chan int)
	quitCh     = make(chan bool)
)

func init() {
	var count int

	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return // end the goroutine.
			}
		}
	}(addCh, getCountCh, quitCh)
}

type singleton struct{}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh // this will block if no respond from the downstream operation.
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}
