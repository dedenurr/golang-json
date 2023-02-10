package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Zahra",
		MiddleName: "Assyifa",
		LastName: "Nurrahman",
	}

	encoder.Encode(customer)
	fmt.Println(customer)
}