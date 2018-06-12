package payment

import (
	"github.com/go_design_pattern/creational/factory/payment"
	"github.com/pkg/errors"
)

type Supermarket struct {
	paymentMethod payment.PaymentMethod // this is an example of strategy design pattern
}

func (sm *Supermarket) SetPaymentMethod(m payment.Method) error {
	pm, err := payment.GetPaymentMethod(m)
	if err != nil {
		return err
	}
	sm.paymentMethod = pm
	return nil
}

func (sm *Supermarket) Pay(amount float32) (string, error) {
	if sm.paymentMethod == nil {
		return "<nil>", errors.New("payment method haven't set")
	}
	return sm.paymentMethod.Pay(amount), nil
}

func (sm *Supermarket) GetPaymentMethodFunc(m payment.Method) (payment.PaymentMethodFunc, error) {
	return payment.GetPaymentMethodFunc(m)
}
