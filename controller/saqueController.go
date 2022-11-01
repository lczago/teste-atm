package controller

import "strconv"

type warn struct {
	Error string `json:"error"`
}

// Valores das notas reais em circulação.
var NotasReais = []int{200, 100, 50, 20, 10, 5, 2}

func Saque(valorParametro string) interface{} {
	valorSaque, err := validaValorSaque(valorParametro)
	if err != "" {
		return warn{err}
	}

	quantidadeNotas := map[int]int{2: 0, 5: 0, 10: 0, 20: 0, 50: 0, 100: 0, 200: 0} // "Valor da nota": "Quantidade"
	calculaQuantidadeNotas(0, valorSaque, &quantidadeNotas)
	return quantidadeNotas
}

// Método que utiliza recursividade para calcular a menor quantidade de cédulas possiveis para saque do valor informado.
func calculaQuantidadeNotas(indiceNota int, valorSaque int, quantidadeNota *map[int]int) {
	if valorSaque == 0 {
		return
	}

	valorNota := NotasReais[indiceNota]
	if valorSaque >= valorNota && valorNota+1 != valorSaque && valorNota+3 != valorSaque {
		quantidade := valorSaque / valorNota
		(*quantidadeNota)[valorNota] += quantidade
		valorSaque -= valorNota * quantidade
	}

	indiceNota++
	calculaQuantidadeNotas(indiceNota, valorSaque, quantidadeNota)
}

// Valida se o parâmetro informado é possivel de ser sacado, caso não seja retorna erro.
func validaValorSaque(valorParametro string) (int, string) {
	valor, err := strconv.Atoi(valorParametro)
	if err != nil {
		return 0, "Apenas são aceitos número inteiros!"
	}

	if valor < 20 {
		return 0, "Valor do saque deve ser igual ou superior a R$20"
	} else if valor%10 == 3 || valor%10 == 1 {
		return 0, "Valor do saque não pode terminar com valor 1 ou 3. Obs: Cedular de 1 real estão fora de circulação"
	}
	return valor, ""
}
