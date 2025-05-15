package main

import "fmt"

type car struct {
	modelo string 
	ano int
	veloMax int
}

func fichaTecnica(c car) {
	fmt.Println("Modelo do carro: ", c.modelo)
	fmt.Println("Ano do carro: ", c.ano)
	fmt.Println("Velocidade maxima: ", c.veloMax)
}

func main() {
	c1 := car{"Onix AT Turbo", 2025, 187}
	fichaTecnica(c1)
}