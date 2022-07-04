package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	operador := "r"
	valores := strings.Split(operacion, operador)
	
	operador1, err1 := strconv.Atoi(valores[0])
	operador2, err2 := strconv.Atoi(valores[1])

	if err1 != nil || err2 != nil{
		fmt.Println("El dato ingresado es incorrecto")
	} 
	
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
	case "-":
		fmt.Println(operador1 - operador2)
	case "*":
		fmt.Println(operador1 * operador2)
	case "/":
		fmt.Println(operador1 / operador2)
	default:
		fmt.Println("El operador no esta soportado")
	}
	
	

}