package checkups

import "testing"

func TestChain(t *testing.T) {
	idCheck := &IDCheckup{}
	healthCheck := &HealthCheckup{
		NextChain: idCheck,
	}
	healthCheck.Next("Jayson")
}

func TestChainDecorator(t *testing.T) {
	idCheck := &IDCheckup{}
	d := Decorate(idCheck, HealthCheckupFunc(), OralCheckupFunc("Colgate"))
	d.Next("jayson")
}