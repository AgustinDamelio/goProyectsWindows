package main

import "fmt"

// Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
// Tener un slice global de Product llamado Products, instanciado con valores.
// Dos métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
// Una función getById() a la cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
// Ejecutar al menos una vez cada método y función definidos desde main().

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{ID: 123456, Name: "Thiago", Price: 123456, Description: "gato", Category: "persona"},
	{ID: 123457, Name: "Nicolas", Price: 123456, Description: "gato", Category: "persona"},
	{ID: 123458, Name: "Agustin", Price: 123456, Description: "gato", Category: "persona"},
}

func (p Product) Save() {
	Products = append(Products, p)
}

func GetAll() {
	for _, p := range Products {
		fmt.Printf("ID: %d\n", p.ID)
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Price: %.2f\n", p.Price)
		fmt.Printf("Description: %s\n", p.Description)
		fmt.Printf("Category: %s\n", p.Category)
		fmt.Println("---------------------------")
	}
}

func getById(id int) Product {
	var product1 Product
	for _, Product := range Products {
		if Product.ID == id {
			product1 = Product
		}
	}
	fmt.Printf("ID: %d\n", product1.ID)
	fmt.Printf("Name: %s\n", product1.Name)
	fmt.Printf("Price: %.2f\n", product1.Price)
	fmt.Printf("Description: %s\n", product1.Description)
	fmt.Printf("Category: %s\n", product1.Category)
	fmt.Println("---------------------------")
	return product1
}

func main() {
	product1 := Product{
		ID:          123410,
		Name:        "pepe",
		Price:       123456,
		Description: "gato",
		Category:    "persona"}

	product1.Save()

	GetAll()

	getById(123410)
}
