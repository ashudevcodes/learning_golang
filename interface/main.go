package main

import (
	"fmt"
)

type payment struct {
	getway paymenter
}

type paymenter interface {
	pay(amount float32)
}

func (p payment) makePayment(amount float32) {
	p.getway.pay(amount)
}

type razepay struct{}

func (r razepay) pay(amount float32) {
	fmt.Println("Making Payment using Razer pay", amount)
}

type strip struct{}

func (s strip) pay(amount float32) {
	fmt.Println("Make Payment using Stripe", amount)

}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Make Payment using PayPal", amount)
}

func main() {
	razepayGw := razepay{}
	newPayment := payment{
		getway: razepayGw,
	}

	stripGw := strip{}
	newPayment2 := payment{
		getway: stripGw,
	}

	paypalGw := paypal{}
	newPayment3 := payment{
		getway: paypalGw,
	}

	newPayment3.makePayment(300)
	newPayment2.makePayment(200)
	newPayment.makePayment(100)

}
