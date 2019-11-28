package payment_test

import (
	"fmt"
	"github.com/jayvib/go_design_pattern/creational/factory/payment"
	"testing"
)

func TestSupermarketPayMethod(t *testing.T) {
	sm := &p.Supermarket{}

	msg, err := sm.Pay(12.349)
	if msg != "<nil>" {
		t.Error("once the payment method haven't set yet message should be <nil>")
	}

	sm.SetPaymentMethod(payment.Cash)

	amount := float32(10.23)
	expected := fmt.Sprintf("An amount of %g is paid using cash.\n", amount)

	msg, err = sm.Pay(amount)
	if err != nil {
		t.Error(err)
	}

	if msg != expected {
		t.Error("expected message doesn't matched.")
	}

}

func TestSupermarket_GetPaymentMethodFunc(t *testing.T) {
	amount := 10.23
	sm := &p.Supermarket{}
	pay, err := sm.GetPaymentMethodFunc(payment.Cash)
	if err != nil {
		t.Error(err)
	}

	msg := pay(float32(amount))

	expected := fmt.Sprintf("An amount of %g is paid using cash.\n", amount)
	if msg != expected {
		t.Error("expected msg wasn't matched")
	}

	fmt.Println(msg)
}
