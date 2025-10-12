package main

import (
	"fmt"
)

// Composition --> Has-A relationship
// Inheritance -> Is-A relationship
// Car -> is composed of many parts (Engine, Doors)

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" { // Handle empty address
		return "No address provided"
	}
	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address // embedded
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Printf("Customer ID: %d\n", c.CustomerID)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Email: %s\n", c.Email)
	fmt.Println("Billing Address:", c.BillingAddress.FullAddress())   // Accessing method of composed type
	fmt.Println("Shipping Address:", c.ShippingAddress.FullAddress()) // Accessing method of composed type
}

func main() {

	fmt.Println("---- Composition -------")

	cust1 := Customer{
		CustomerID: 1001,
		Name:       "Gadget Corp",
		Email:      "sales@gadgetcorp.com",
		BillingAddress: Address{
			Street:  "123 Tech Road",
			City:    "Innovateville",
			State:   "CA",
			ZipCode: "90210",
		},
		ShippingAddress: Address{ // Different shipping address
			Street:  "456 Factory Lane",
			City:    "Manufacturicity",
			State:   "NV",
			ZipCode: "89101",
		},
	}

	cust1.PrintDetails()

	fmt.Println("------- customer with same billing and shipping address -------")
	mainAddress := Address{
		Street:  "789 Main St",
		City:    "Anytown",
		State:   "TX",
		ZipCode: "75001",
	}
	cust2 := Customer{
		CustomerID:      1002,
		Name:            "John Doe",
		Email:           "john.doe@email.com",
		BillingAddress:  mainAddress,
		ShippingAddress: mainAddress, // Reusing the Address instance
	}

	cust2.PrintDetails()
}
