package main

import (
	"fmt"
	"strconv"
	"strings"
)

var follow = map[string][]string{
	"P'":    {"EOF"},
	"P":     {"EOF"},
	"V":     {"leia", "escreva", "se", "facaAte", "fim"},
	"LV":    {"leia", "escreva", "se", "facaAte", "fim"},
	"D":     {"id", "varfim"},
	"L":     {"int", "real", "lit"},
	"TIPO":  {"PT_V"},
	"A":     {"EOF"},
	"ES":    {"leia", "escreva", "id", "se", "fim", "fimse", "facaAte", "fimFaca"},
	"ARG":   {"PT_V"},
	"CMD":   {"leia", "escreva", "id", "se", "fimse", "fimFaca"},
	"LD":    {"PT_V"},
	"OPRD":  {"leia", "escreva", "id", "se", "fim", "fimse", "facaAte", "fimFaca"},
	"COND":  {"leia", "escreva", "id", "se", "fim", "fimse", "facaAte", "fimFaca"},
	"CAB":   {"leia", "escreva", "id", "se", "fimse"},
	"EXP_R": {"FC_P"},
	"CP":    {"leia", "escreva", "id", "se", "fim", "fimse", "facaAte", "fimFaca"},
	"R":     {"leia", "escreva", "id", "se", "fim", "facaAte"},
	"CP_R":  {"leia", "escreva", "id", "se", "fim", "facaAte"},
}

var traducaoTipo = map[string]string{
	"num":     "numerico",
	"literal": "literal",
	"id":      "Variavel",
	"OPR":     "Operador Relacional",
	"RCB":     "<-",
	"OPM":     "Operador Matemático",
	"AB_P":    "(",
	"FC_P":    ")",
	"PT_V":    ";",
	"VIR":     ",",
}

func entradasEsperadasErro(estado string, token *Token) ([]string, []string) {
	esperados := []string{}
	redutiveis := []string{}
	for k, v := range tabelaSintaticaTerminais[estado] {
		if k == "Estado" {
			continue
		}
		if v != "" {
			esperados = append(esperados, k)
			if string(tabelaSintaticaTerminais[estado][k][0]) == "R" {
				redutiveis = append(redutiveis, k)
			}
		}
	}
	return esperados, redutiveis

}

func Parse(tokens []*Token) (err bool) {
	err = false
	errosSeguidos := 0
	inicializarTabelaSintaticas()
	inicializarCodigoObjeto()
	sintaticaStack := Stack{"0"}
	semanticaStack := StackToken{}
	indexTokenList := 0
	var tokenAntigo *Token
	var token *Token
	for {
		if errosSeguidos > 5 {
			break
		}
		if tokenAntigo == nil {
			token = tokens[indexTokenList]
		}

		if token!= nil && (token.classe == "comentario" || token.classe == "ERRO") {
			indexTokenList++
			continue
		}
		estadoPilha := sintaticaStack[len(sintaticaStack)-1]
		acao := tabelaSintaticaTerminais[estadoPilha][token.classe]

		if acao != "" {
			errosSeguidos = 0
			//fmt.Printf("AcaO:%s \n", acao)
			//SHIFT
			if string(acao[0]) == "S" {
				estadoEmpilhar := acao[1:]
				sintaticaStack.Push(estadoEmpilhar)
				semanticaStack.Push(token)

				if tokenAntigo != nil {
					token = tokenAntigo
					tokenAntigo = nil
				} else {
					indexTokenList++
				}
			}
			//REDUCE
			if string(acao[0]) == "R" {
				indiceRegra := acao[1:]
				beta, _ := strconv.Atoi(tabelaSintaticaRegras[indiceRegra]["TamanhoB"])
				for i := 0; i < beta; i++ {
					sintaticaStack.Pop()
				}

				estadoPilha = sintaticaStack[len(sintaticaStack)-1]
				//fmt.Printf("Novo estado:%s \n", estadoPilha)
				//GOTO
				p1 := tabelaSintaticaRegras[indiceRegra]["P1"]
				novoEstado := tabelaSintaticaNaoTerminais[estadoPilha][p1]
				sintaticaStack.Push(novoEstado)
				fmt.Printf("%s -> %s\n", tabelaSintaticaRegras[indiceRegra]["P1"], tabelaSintaticaRegras[indiceRegra]["P2"])

				// Semantico
				var tokensSemanticos []*Token
				for i := 0; i < beta; i++ {
					token, isOk := semanticaStack.Pop()
					if !isOk {
						err = true
						_ = fmt.Errorf("Erro ao desimpilhar\n")
					} else {
						tokensSemanticos = append(tokensSemanticos, token)
					}
				}
				tokensSemanticos = reverseArray(tokensSemanticos)
				// Ajustar a passagem do atributo de TIPO.tipo para o id.tipo
				if indiceRegra == "7" || indiceRegra == "8" && len(tokensSemanticos) == 1 {
					tokensSemanticos[0].tipo = semanticaStack[len(semanticaStack)-1].tipo
				}

				tokenSemanticoReduzido, isErro := analisadorSemantico(indiceRegra, p1, tokensSemanticos)
				err = err || isErro
				semanticaStack.Push(tokenSemanticoReduzido)


			}
			//ACCEPT
			if acao == "ACCEPT" {
				break
			}
		} else {
			errosSeguidos++
			err = true
			var erro = ""
			esperados, redutiveis := entradasEsperadasErro(estadoPilha, token)
			if token != nil {
				if token.classe == "PT_V" {
					indexTokenList++
					erro = "Remova o ';'"
				} else if token.classe == "FC_P" {
					indexTokenList++
					erro = "Remova o ')'"
				} else if len(esperados) == 1 {
					tokenAntigo = token
					t:= Token{"", esperados[0], "-", tokenAntigo.lin, tokenAntigo.col}
					token = &t
					erro = "Você esqueceu de adicionar o " + traducaoTipo[esperados[0]]
				} else if len(redutiveis) == 1 {
					tokenAntigo = token
					t:= Token{"", esperados[0], "-", tokenAntigo.lin, tokenAntigo.col}
					token = &t
				}

				if len(esperados) > 1 {
					erro = "Entrada " + token.lexema + " inesperada. São esperados: " + strings.Join(esperados, ",")
				}

				HandleError(erro, token.lin, token.col)

			}

			if len(redutiveis) > 1 {
				break
			}

		}
	}
	return err
}

func reverseArray(arr []*Token) []*Token{
	for i, j := 0, len(arr)-1; i<j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}