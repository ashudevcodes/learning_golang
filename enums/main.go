package main

import "fmt"

// Custon types
type MyType string

// enums is an emurated types but it not in go

type OrderStatus int

type UserRole string

const (
	Admin UserRole = "admin"
	User  UserRole = "user"
)

// iota increment value
const (
	Recived OrderStatus = iota
	Confermed
	Prepared
	Deliverd
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing Order status to", status)
}

func chnageUserRole(role UserRole) {
	fmt.Println("Role is", role)
}

func main() {

	changeOrderStatus(Recived)
	chnageUserRole(Admin)
}
