<h1 align="center">Hi there, I'm <a href="https://github.com/HeorhiiZemlianko" target="_blank">Heorhii</a> 
<img src="https://github.com/blackcater/blackcater/raw/main/images/Hi.gif" height="32"/></h1>
<h3 align="center">This repository shows one of the options for solving modular mathematics.</h3>
<img src="https://badges.frapsoft.com/os/v1/open-source.svg?v=103" >

## The goal of the work
Master modular arithmetic, algorithms for finding the greatest common divisor (GCD) and reciprocal element, finding prime numbers.

## Task statement
Create a program that provides arithmetic operations modulo some number and finding the inverse element in the corresponding group (G).

The software implementation should have the following capabilities:
- set the module m, according to which the calculations will be carried out;
- solve equations of the form **a mod m = x**;
- solve equations of the form **a^b mod m = x**;
- solve equations of the form **a*x ≡ b mod m**;
- generate a prime number in the range **A** to **B**.

## Notes
This program is implemented in the **Go** language, the following libraries were used: **`fmt`** and **`math`**

## Notation of variables
This table shows the short designation of the variable in the code and its description

| Name       | Type   | Description                      |
| ---------- | ------ | -------------------------------- |
| `a` | int | Any entered number |
| `m` | int | Modulus of number  |
| `b` | int | Degree of number   |
| `i` | int | The number of job samples at the very beginning of the program |
| `num1` | int | Lower limit for prime number generation |
| `num2` | int | Upper limit for prime number generation |

## Functions in the project code

- Calculation of GSD (NSD)
``` Go
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
```

- Calculation a*x ≡ b mod m
``` Go
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
```

- Calculation of the modulus to the power a^b mod m = x
``` Go
func binpow(a, b, m int) int {
	if b == 0 {
		return 1 % m
	} else if b%2 == 1 {
		return (binpow(a, b-1, m) * a) % m
	} else {
		return binpow((a*a)%m, b/2, m)
	}
}
```

- Prime number generation
``` Go
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
```
