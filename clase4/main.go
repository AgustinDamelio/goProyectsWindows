package main

import "fmt"

type Product interface {
	Price() (totalCost float32)
}

type SmallProduct struct {
	cost float32
}

func (sp SmallProduct) Price() (totalCost float32) {
	return sp.cost
}

type MediumProduct struct {
	cost float32
}

func (mp MediumProduct) Price() (totalCost float32) {
	return mp.cost * 1.03
}

type BigProduct struct {
	cost float32
}

func (bp BigProduct) Price() (totalCost float32) {
	return bp.cost*1.06 + 2500
}

func factory(productType string, cost float32) Product {
	switch productType {
	case "big":
		return BigProduct{cost: cost}
	case "medium":
		return MediumProduct{cost: cost}
	case "small":
		return SmallProduct{cost: cost}
	default:
		return nil
	}
}

func main() {
	product1 := factory("big", 2000)
	fmt.Println(product1.Price())
	product2 := factory("medium", 1500)
	fmt.Println(product2.Price())
	product3 := factory("big", 1000)
	fmt.Println(product3.Price())
}
