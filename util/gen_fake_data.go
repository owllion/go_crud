package util

import (
	"fmt"
	"math/rand"
)

type Product struct {
	Name  string
	Qty   int
	Price float64
}

func GenFakeData() []Product {
	products := []Product{}
	for i := 1; i <= 10; i++ {
		p := Product{
			Name: fmt.Sprintf("Product %d", i),
			Qty: rand.Intn(50) +1 ,
			Price: rand.Float64() * 100,
		}
		products = append(products, p)
	}

	return products
}