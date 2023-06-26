package main

type(
	Customer struct {
	firstName string
	lastName string
	fullName string
}
)

func getCustomers()(customers []Customer){

	praveen := Customer{firstName: "Praveen", lastName: "Hosamani"}

	customers = append(customers, praveen,
		Customer{firstName: "Pragati", lastName: "Hosamani"},
	)

	return customers
}