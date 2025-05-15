package main

import "fmt"

type Jogador struct {
	nome string 
	vida int
	nivel int
}

func exibeDados(j Jogador) {
	fmt.Println("Nome do jogador: ", j.nome)
	fmt.Println("Vida do Jogador: ", j.vida)
	fmt.Println("Nível do Jogador: ", j.nivel)
}

func main() {
	j1 := Jogador{"Astrubal", 100, 1}
	exibeDados(j1)
}