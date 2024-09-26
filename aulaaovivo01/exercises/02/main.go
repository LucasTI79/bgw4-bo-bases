package main

import "fmt"

/*
Declare 3 variáveis especificando o tipo de dados; como valor, elas devem ter a temperatura, a umidade e a pressão do local onde você está.
Imprima os valores das variáveis no console.
Que tipo de dados você atribuiria às variáveis?
*/

func main() {
	var temperatura float64
	var umidade float64
	var pressaoAtmosferica float64

	temperatura = 28.2
	umidade = 47.5
	pressaoAtmosferica = 944.2

	// fmt.Println("temperatura", temperatura, "umidade", umidade, "pressão atmosférica", pressaoAtmosferica)

	texto := fmt.Sprintf("temperatura: %.1f, umidade: %.1f, pressão atmosférica: %.2f", temperatura, umidade, pressaoAtmosferica)
	println(texto)
}
