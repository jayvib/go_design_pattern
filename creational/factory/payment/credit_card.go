package payment

import "fmt"

type CreditCardPayment struct {}

func (c *CreditCardPayment) Pay(amount float32) string {
	return fmt.Sprintf("An amount of %f is paid using debit card (new)\n", amount)
}
