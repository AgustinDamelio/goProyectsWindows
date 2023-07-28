package goproyectswindows

import "fmt"

func main() {

	nombres := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	nombres = append(nombres, "Grabiela")

	fmt.Println(nombres)
	fmt.Println()

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("La edad de Benjamin es: %v", employees["Benjamin"])
	fmt.Println()

	sum := 0
	for _, edad := range employees {
		if edad > 21 {
			sum++
		}
	}

	fmt.Printf("El total de mayores es: %v", sum)
	fmt.Println()

	delete(employees, "Pedro")
	fmt.Print(employees)

}
