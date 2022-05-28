package main

import (
	"fmt"
	"math"
)

func main() {
	/*1. Напишите программу для вычисления площади прямоугольника. Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.*/
	var a, b int
	fmt.Println("Введите длины сторон прямоугольника")
	fmt.Scanln(&a, &b)
	fmt.Println(squareOfRectangle(a, b))
	/*2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.*/

	var squareOfCircle float64
	fmt.Println("Введите площадь круга")
	fmt.Scanln(&squareOfCircle)
	fmt.Println(parametersOfCircle(squareOfCircle))

	/*	3. С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количеству сотен, десятков и единиц в этом числе. Например, для введенного 345 должно вывестись 3 сотни, 4 десятка, 5 единиц. Нужно сделать это через математические операции.*/
	var n int
	fmt.Println("Введите трёхзначное число")
	fmt.Scanln(&n)
	fmt.Println(threeDigit(n))
}

func parametersOfCircle(squareOfCircle float64) (float64, float64) {
	diameter := math.Sqrt(4.00 * squareOfCircle / math.Pi)
	lenOfCircle := math.Pi * diameter
	return diameter, lenOfCircle
}

func squareOfRectangle(a, b int) int {
	result := a * b
	return result
}

func threeDigit(n int) string {

	var strHundreds, strDozens, strUnits string

	intHundreds := n / 100
	intDozens := (n / 10) % 10
	intUnits := n % 100 % 10

	if intHundreds == 1 {
		strHundreds = "сотня"
	} else if intHundreds > 1 && intHundreds < 5 {
		strHundreds = "сотни"
	} else {
		strHundreds = "сотен"
	}

	if intDozens == 1 {
		strDozens = "десяток"
	} else if intDozens > 1 && intDozens < 5 {
		strDozens = "десятка"
	} else {
		strDozens = "десятков"
	}

	if intUnits == 1 {
		strUnits = "единица"
	} else if intUnits > 1 && intUnits < 5 {
		strUnits = "единицы"
	} else {
		strUnits = "единиц"
	}
	result := fmt.Sprint(intHundreds, " ", strHundreds, " ", intDozens, " ", strDozens, " ", intUnits, " ", strUnits)
	return result
}
