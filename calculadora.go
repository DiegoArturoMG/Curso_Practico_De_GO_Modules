package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Calc struct {
	e1 int
	e2 int
}

func (self Calc) Operate(operador string) int {
	operador1 := self.e1
	operador2 := self.e2

	var resultado int

	switch operador {
	case "+":
		resultado = operador1 + operador2
	case "-":
		resultado = operador1 - operador2
	case "*":
		resultado = operador1 * operador2
	case "/":
		if operador2 == 0{
			fmt.Println("Error: Division entre cero")
			return 0
		}
		resultado = operador1 / operador2
	default:
		fmt.Println("Operador no aceptado")
		resultado = 0
	}

	return resultado
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}