package payment

import (
	"errors"
	"fmt"
)

type Method int

const (
	Cash      Method = 1
	DebitCard Method = 2
)

type PaymentMethodFunc func(float32) string

func CashPay(amount float32) string {
	return fmt.Sprintf("An amount of %g is paid using cash.\n", amount)
}

func DebitCardPay(amount float32) string {
	return fmt.Sprintf("An amount of %f is paid using debit card (new)\n", amount)
}

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m Method) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(CreditCardPayment), nil
	default:
		return nil, errors.New(fmt.Sprintf("payment method %d does not recognized", m))
	}
}

func GetPaymentMethodFunc(m Method) (PaymentMethodFunc, error) {
	switch m {
	case Cash:
		return CashPay, nil
	case DebitCard:
		return DebitCardPay, nil
	default:
		return nil, errors.New(fmt.Sprintf("payment method %d does not recognized", m))
	}
}
