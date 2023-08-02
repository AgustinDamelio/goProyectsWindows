package main

import (
	"errors"
	"fmt"
)

func main() {

	// Ejercicio 1 - Calcular salario
	// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
	// Categoría C: su salario es de $1.000 por hora.
	// Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
	// Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
	// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

	fmt.Println(calculateWage(60, "C"))
	fmt.Println(calculateWage(60, "B"))
	fmt.Println(calculateWage(60, "A"))
	fmt.Println(calculateWage(60, "ZZZ"))

}

func calculateWage(minutes int, category string) (wage float64, err error) {
	switch category {
	case "C":
		wage = categoryC(minutes)
	case "B":
		wage = categoryB(minutes)
	case "A":
		wage = categoryA(minutes)
	default:
		return 0, errors.New("Category non existent")
	}
	return wage, nil
}

func categoryC(minutes int) (wage float64) {
	hours := float64(minutes / 60)
	wage = hours * 1000
	return
}

func categoryB(minutes int) (wage float64) {
	hours := float64(minutes / 60)
	wage = hours * 1500 * 1.2
	return
}

func categoryA(minutes int) (wage float64) {
	hours := float64(minutes / 60)
	wage = hours * 3000 * 1.5
	return
}
