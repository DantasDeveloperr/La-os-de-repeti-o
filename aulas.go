//Aula sobre arrays e fatias em Go
/*
package main

import "fmt"

func main() {

	array := [5]float64{7.5, 8.0, 9.6, 6.4, 5.5}
	fatia := array[1:3]
	fmt.Println(fatia)

}
*/

// Aula sobre mapas em Go
package main

import "fmt"

func main() {

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["O"] = "Oxigênio"

	h := elemento["H"]
	o := elemento["O"]
	resultado := fmt.Sprintf("A fórmula da água é %s%s", h, o)
	fmt.Println(resultado)

	// Exemplo de criação de mapa com literais
	//	meuMapa := map[string]int{
	//		"pedro": 30,
	//		"maria": 25,
	//	}
	//	fmt.Println(meuMapa["maria"])

	//	meuMapa := make(map[string]int)
	//	meuMapa["pedro"] = 30

	//	fmt.Println(meuMapa["pedro"])
}
