package libs

import (
	"math/rand"
	"strconv"
)

var statuses = []string{"paid", "pending", "draft"}

func GenerateRandomIdInvoices() string {
	return "#RT" + strconv.Itoa(rand.Intn(5000)+1000)
}

func GenerateRandomStatus() string {
	return statuses[rand.Intn(len(statuses))]
}

func GenerateRandomPrice() float64 {
	return rand.Float64() * 2000
}

func GenerateRandomName() string {
	names := []string{"Jensen Huang", "John Doe", "Jane Doe", "Alice Smith", "Bob Johnson"}
	return names[rand.Intn(len(names))]
}
