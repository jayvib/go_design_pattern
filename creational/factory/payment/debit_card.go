package payment

import "fmt"

type DebitCardPM struct {}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("An amount of %g is paid using debit card.\n", amount)
}
