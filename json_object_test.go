package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// untuk membuat json dibutuhkan sebuah struct => jadi buat terlebih dahulu structnya
type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age int
	Married bool
	Hobbies []string
	Addresses []Address
}

type Address struct{
	Street string
	Country string
	PostalCode string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Zahra",
		MiddleName: "Assyifa",
		LastName: "Nurrahman",
		Age : 6,
		Married: false,

	}
	//proses konversi tipe data di golang kedalam bentuk json
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

