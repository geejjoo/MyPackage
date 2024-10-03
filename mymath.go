package mymath

import "math"

// Sqrt - функция для вычисления квадратного корня из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Add - функция для сложения двух чисел
func Add(a, b float64) float64 {
	return a + b
}

// Subtract - функция для вычитания двух чисел
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply - функция для умножения двух чисел
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide - функция для деления двух чисел
func Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

func Abs(a float64) float64 {
	return math.Abs(a)
}

func Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
