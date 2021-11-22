// G. Бонусы для водителей
// Водители Яндекс.Такси узнали о раздаче бонусов и выстроились перед офисом. У каждого водителя есть рейтинг. Необходимо раздать водителям бонусы, соблюдая следующие условия:
// • Сумма бонуса кратна 500 рублям.
// • Каждый водитель должен получить как минимум 500 рублей.
// • Водитель с бóльшим рейтингом должен получить бóльшую сумму бонуса, чем его соседи слева или справа с меньшим рейтингом.
//
// Какое минимальное количество денег потребуется на бонусы?
//
// Формат ввода
// На первой строчке записано число N (1 <= N <= 20000), далее следует N строчек с рейтингами водителей Rn (0 <= Rn < 4096)
//
// Формат вывода
// Ответ должен содержать минимально необходимое количество денег для выплаты вознаграждений
//
// Пример 1
// Ввод	Вывод
// 4		5000
// 1
// 2
// 3
// 4
//
// Пример 2
// Ввод	Вывод
// 4		2000
// 5
// 5
// 5
// 5
//
// Пример 3
// Ввод	Вывод
// 4		3000
// 4
// 2
// 3
// 3
package main

import (
	"fmt"
)

const minBonus = 500

func main() {
	var total int

	fmt.Scanf("%d", &total)

	ratings := make([]int, 0, total)
	for i := 0; i < total; i++ {
		var value int
		fmt.Scanf("%d", &value)
		ratings = append(ratings, value)
	}

	// ratings := []int{1, 2, 3, 4}
	// ratings := []int{5, 5, 5, 5}
	// ratings := []int{4, 2, 3, 3}
	// ratings := []int{0}

	totalBonusSum := 0
	minRating := ratings[0]
	ratingScale := 1
	for _, rating := range ratings {
		// fmt.Println("rating", rating)
		// fmt.Println("minRating", minRating)
		// fmt.Println("ratingScale", ratingScale)

		var bonusSum int

		if rating < minRating {
			minRating = rating
			bonusSum += minBonus
		} else if rating > minRating {
			ratingScale++
			minRating = rating
			bonusSum += minBonus * ratingScale
		} else {
			bonusSum += minBonus * ratingScale
		}
		totalBonusSum += bonusSum

		// fmt.Println(bonusSum)
		// fmt.Println("============")
	}

	fmt.Println(totalBonusSum)
}
