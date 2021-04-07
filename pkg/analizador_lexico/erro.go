package analizador_lexico

import "fmt"

type Error struct {
	erro string
	linha, coluna int
}

var erros = map[string]string{
	"ERRO1": "Número informado inválido",
	"ERRO2": "Você esqueceu de fechar as aspas no literal declarado",
	"ERRO3": "Você esqueceu de fechar chaves no seu comentário",
	"ERRO4": "Caracter inválido na linguagem",
}

func HandleError(erro string, linha, coluna int) {
	fmt.Printf("%s - %s, linha %d, coluna %d\n", erro, erros[erro], linha, coluna)
}
