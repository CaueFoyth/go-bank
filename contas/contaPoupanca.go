package contas

import "github.com/FoythCaue/go-bank/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito efetuado com sucesso!", c.saldo
	} else {
		return "Erro ao efetuar o deposito", c.saldo
	}
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) bool {
	if valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
