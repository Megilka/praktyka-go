package main

import "fmt"

type calculator struct {
	a float64
	b float64
}

func (c calculator) add() float64 {
	return c.a + c.b
}

func (c calculator) subtract() float64 {
	return c.a - c.b
}
func (c calculator) multiply() float64 {
	return c.a * c.b
}

func (c calculator) divide() float64 {
	return c.a / c.b
}

func main() {
	//result := calculator{12, 5}.add()
	var a, b float64
	fmt.Scan(&a, &b)

	c := calculator{a: a, b: b}

	fmt.Printf("Suma: %v", c.add())
	fmt.Printf("Roznica: %v", c.subtract())
	fmt.Printf("Multiply: %v", c.multiply())
	fmt.Printf("divide: %v", c.divide())
}
