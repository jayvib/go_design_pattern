package observer

type Observer interface {
	Notify(message string)
}

type Publisher struct {
	ObserversList []Observer
}

func (s *Publisher) AddObserver(o Observer) {}

func (s *Publisher) RemoveObserver(o Observer) {}

func (s *Publisher) NotifyObservers(m string) {}
