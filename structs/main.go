package main

import (
	"fmt"
	"time"
)

// order struct

type order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time // nonosecend precision
	customer
}

type customer struct {
	name        string
	phoneNumber int
}

// reciver type
// Struce doing deRefrencng autimaticealy
func (o *order) chnageStatue(status string) {
	o.status = status

}

// Return Amount
func (o order) getAmount() float32 {
	return o.amount

}

// constructor New is a convenction

func newOrder(id int, amount float32, status string) *order {
	myOrder := &order{
		id:     id,
		amount: amount,
		status: status,
	}
	return myOrder
}

func main() {

	// if i don't set any field, default value is zero value
	// int => 0 , float => 0, string => "", bool => false

	language := struct {
		name   string
		isGood bool
	}{"GoLang", true}

	fmt.Println(language)
	order1 := order{
		id:     1,
		amount: 200.2,
		status: "active",
	}
	order1.createdAt = time.Now()

	order1.chnageStatue("Conform")

	order2 := order{
		id:        2,
		amount:    23.9,
		status:    "deliverd",
		createdAt: time.Now(),
	}

	order3 := newOrder(3, 30, "active")

	newOrder := order{
		id:        4,
		amount:    30,
		status:    "active",
		createdAt: time.Now(),
		customer: customer{
			name:        "Ashish",
			phoneNumber: 123456789,
		},
	}

	fmt.Println(newOrder)
	fmt.Println(order3)

	fmt.Println(order1)
	fmt.Println(order2)
	fmt.Println(order1.getAmount())

}
