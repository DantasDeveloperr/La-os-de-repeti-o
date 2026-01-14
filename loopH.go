package main

import "fmt"

func main() {
	for horas := 0; horas <= 24; horas++ {
		fmt.Println("São Horas:", horas)

		for minutos := 0; minutos < 60; minutos++ {
			fmt.Println("Minutos:", minutos)

			for segundos := 0; segundos < 60; segundos++ {
				fmt.Println("Segundos:", segundos)
			}
		}
	}

}

// Quando você for rodar este código, ele irá imprimir todas as combinações possíveis de horas, minutos e segundos em um formato de relógio de 24 horas. A estrutura de loops aninhados permite que cada unidade de tempo seja iterada completamente antes de passar para a próxima unidade.
// para parar a execução do código, você pode usar Ctrl + C no terminal onde o programa está sendo executado.
