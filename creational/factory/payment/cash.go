package payment

import "fmt"

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("An amount of %g is paid using cash.\n", amount)
}
