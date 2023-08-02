package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result := calculateWagewage(50001)
	fmt.Println("Your wage is", result)

	scores := []int{9, 2, 5, 4, 7, 8}
	result2, err := calculateAverageScore(scores...)

	if err != nil {
		fmt.Println(errors.New("No puede haber notas negativas"))
	} else {
		fmt.Println("Tiene un promedio de", result2)
	}

}

func calculateWagewage(basicIncome int) (wage float64) {
	wage = float64(basicIncome)
	if basicIncome > 50000 {
		return math.Round(wage * 0.83)
	} else if basicIncome > 150000 {
		return math.Round(wage * 0.73)
	} else {
		return wage
	}
}

func calculateAverageScore(scores ...int) (averageScore int, err error) {
	scoresSum := 0
	for _, score := range scores {
		if score >= 0 {
			scoresSum += score
		} else {
			return 0, errors.New("error")
		}

	}
	averageScore = scoresSum / len(scores)
	return averageScore, nil
}
