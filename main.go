package main

import (
	"fmt"

	"github.com/FoythCaue/go-bank/clientes"
	"github.com/FoythCaue/go-bank/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.456.789-00",
		Profissao: "Desenvolvedor",
	}
	contaDoBruno := contas.ContaCorrente{
		Titular:       clienteBruno,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	contaDoBruno.Depositar(1000)
	PagarBoleto(&contaDoBruno, 500)
	fmt.Println(contaDoBruno.ObterSaldo())
}
