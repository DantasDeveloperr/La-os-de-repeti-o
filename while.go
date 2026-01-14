// O while no Go não existe como uma construção separada, mas você pode usar o for para alcançar o mesmo efeito.
package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

//indicação de livro: "The Go Programming Language" by Alan A. A. Donovan and Brian W. Kernighan -> A linguagem de Programação Go
//Estrutura de repetição usando "for" como "while"
