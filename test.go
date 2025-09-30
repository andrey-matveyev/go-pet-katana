package main

import (
	"fmt"
)

// findAllCompositionsFixedK находит и печатает все слайсы длиной K,
// сумма элементов которых равна targetSum (N), с заданными ограничениями.
func findAllCompositionsFixedK(targetSum, k int) {
	if k < 2 {
		fmt.Println("Ошибка: Длина слайса K должна быть >= 2.")
		return
	}

	// Инициализируем слайс нужной длины
	slice := make([]int, k)

	// Начинаем рекурсивный поиск с индекса 0 и полной суммы
	findRecursiveFixedK(targetSum, 0, slice)
}

// Рекурсивная функция для поиска композиций
func findRecursiveFixedK(remainingSum, index int, currentSlice []int) {
	k := len(currentSlice)

	// --- 1. Базовый случай (Успех) ---
	if index == k {
		// Проверяем, что вся сумма была распределена
		if remainingSum == 0 {
			// Нашли правильное разбиение! Печатаем копию.
			result := make([]int, k)
			copy(result, currentSlice)
			fmt.Println(result)
		}
		return
	}

	// --- 2. Определение ограничений (minVal и maxVal) для текущего элемента ---

	minVal := 0
	maxVal := remainingSum

	// a) Ограничения, зависящие от позиции:
	if index > 0 && index < k-1 {
		// Средние элементы должны быть >= 1
		minVal = 1
	}

	// b) Ограничения, зависящие от оставшейся суммы:

	// Если это последний элемент (k-1), его значение должно быть равно оставшейся сумме.
	if index == k-1 {
		// Крайний элемент >= 0, поэтому нам просто нужно убедиться, что оставшаяся сумма >= 0
		if remainingSum < 0 {
			// Если оставшаяся сумма < 0, это неверный путь
			return
		}
		minVal = remainingSum
		maxVal = remainingSum
	} else {
		// Если это не последний элемент, мы должны оставить достаточно суммы
		// для оставшихся элементов (k - 1 - index).

		// Рассчитываем минимальную необходимую сумму для оставшихся элементов.
		minSumForRemaining := 0
		for i := index + 1; i < k; i++ {
			// Все оставшиеся элементы, кроме последнего, должны быть >= 1
			if i > 0 && i < k-1 {
				minSumForRemaining += 1
			}
			// Последний элемент и, возможно, первый, могут быть 0
		}

		// Обновляем maxVal, чтобы гарантировать, что оставшаяся сумма >= minSumForRemaining
		maxVal = remainingSum - minSumForRemaining
	}

	// Если minVal > maxVal (из-за ограничений), перебор невозможен.
	if minVal > maxVal || maxVal < 0 {
		return
	}

	// --- 3. Рекурсивный перебор ---
	for val := minVal; val <= maxVal; val++ {
		currentSlice[index] = val

		// Рекурсивный вызов для следующего элемента
		findRecursiveFixedK(remainingSum-val, index+1, currentSlice)
	}
}

func main1() {
	N := 15 // Целевая сумма
	K := 5  // Фиксированная длина слайса

	fmt.Printf("Все композиции N=%d для длины K=%d (Первый/Последний>=0, Средние>=1):\n", N, K)
	fmt.Println("----------------------------------------------------------------------")
	findAllCompositionsFixedK(N, K)

	N2 := 5
	K2 := 3
	fmt.Printf("\nВсе композиции N=%d для длины K=%d:\n", N2, K2)
	fmt.Println("----------------------------------------------------------------------")
	findAllCompositionsFixedK(N2, K2)
	// Пример результатов для N=5, K=3 (Первый>=0, Средний>=1, Последний>=0):
	// [0 1 4]
	// [0 2 3]
	// [0 3 2]
	// [0 4 1]
	// [1 1 3]
	// [1 2 2]
	// [1 3 1]
	// [2 1 2]
	// [2 2 1]
	// [3 1 1]
	// [4 1 0]
}
