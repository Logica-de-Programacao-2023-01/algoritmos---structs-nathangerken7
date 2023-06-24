package main

import (
	"fmt"
)

type Pessoa struct {
	nome  string
	idade int
	endereço
}

type endereço struct {
	rua    string
	numero string
	cidade string
	estado string
}

func main() {

	p := Pessoa{
		nome:  "Matheus",
		idade: 19,
		endereço: endereço{
			rua:    "Rua",
			numero: "Numero",
			cidade: "Cidade",
			estado: "Estado",
		},
	}
	fmt.Printf("Nome: %s \nIdade: %d \nRua: %s \nnumero: %s \ncidade: %s \nestado: %s", p.nome, p.idade, p.rua, p.numero, p.cidade, p.estado)

}
