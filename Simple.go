package main

import (
	"fmt"
	"math"
)

func main() {
	var a, m, b, i, num1, num2 int

	//Выбор задачи
	fmt.Print("\n1. Решение модульных выражений\n")
	fmt.Print("2. Генерация простого числа\n")
	fmt.Print("\nВведите номер задачи: ")
	fmt.Scan(&i)
	if i == 1 {
		//Инфа
		fmt.Print("\nВыражения которые тут вычисляются:\n")
		fmt.Print("1. a mod m = x\n")
		fmt.Print("2. a^b mod m = x\n")
		fmt.Print("3. a*x ≡ b mod m\n")
		//Ввод данных
		fmt.Print("\nВведите число: ")
		fmt.Scan(&a)
		fmt.Print("Введите значение модуля: ")
		fmt.Scan(&m)
		fmt.Print("Введите значение степени, для решения a^b mod m = x: ")
		fmt.Scan(&b)
		// Вывод результатов
		fmt.Println("\nРешение уравнения вида a mod m = x, ответ: x =", a%m)
		fmt.Println("Решение уравнения вида a^b mod m = x, ответ: x =", binpow(a, b, m))
		fmt.Println("Решение уравнения вида a*x ≡ b mod m, ответ: x =", modInverses(a, b, m))
	} else if i == 2 {
		fmt.Print("\nВведите начальное значение диапозона: ")
		fmt.Scan(&num1)
		fmt.Print("Введите конечное значение диапозона: ")
		fmt.Scan(&num2)
		printPrimeNumbers(num1, num2)
		fmt.Print("\n")
	} else {
		fmt.Println("\nВведите корректный номер")
	}
}

// Расчет GSD (НСД)
func gcdExtended(a int, b int, m int, x *int, y *int) int {
	if a == 0 {
		*x = 0
		*y = b
		return m
	}
	var x1, y1, gcd int // To store results of recursive call
	gcd = gcdExtended(m%a, b, a, &x1, &y1)
	*x = y1 - (m/a)*x1
	*y = x1
	return gcd
}

// Расчет  a*x ≡ b mod m
func modInverses(a, b, m int) int {
	var x, y, g, res int
	g = gcdExtended(a, b, m, &x, &y)
	if g != 1 {
		fmt.Println("Инверсии не существует")
	} else {
		res = (x%m + m) % m
	}
	return res
}

// Расчет модуля в степени a^b mod m = x
func binpow(a, b, m int) int {
	if b == 0 {
		return 1 % m
	} else if b%2 == 1 {
		return (binpow(a, b-1, m) * a) % m
	} else {
		return binpow((a*a)%m, b/2, m)
	}
}

// Генерация простых чисел
func printPrimeNumbers(num1, num2 int) {
	if num1 < 2 || num2 < 2 {
		fmt.Println("Числа должны быть больше 2.")
		return
	}
	fmt.Printf("\nПростые числа в диапозоне от %d до %d:\n\n", num1, num2)
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num1)
		}
		num1++
	}
	fmt.Println()
}
