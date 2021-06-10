package main

import (
	"fmt"
	"strings"
)

func analisadorSemantico(regra, p1 string, tokens []*Token) (*Token, bool) {

	nToken := Token{p1, p1, "", tokens[0].lin, tokens[0].col}

	if regra == "5" {
		//LV→ varfim;
		linCodigoObjeto.Push("\n")
		linCodigoObjeto.Push("\n")
		linCodigoObjeto.Push("\n")
	} else if regra == "6" {
		// D→ TIPO L;
		linCodigoObjeto.Push(";\n")
	} else if regra == "7" || regra == "8" {
		// L→ id, L || L→ id
		nToken.tipo = tokens[0].tipo
		linCodigoObjeto.Push(tokens[0].lexema)
	} else if regra == "9" {
		// TIPO→ int
		tokens[0].tipo = "int"
		nToken.tipo = tokens[0].tipo
		linCodigoObjeto.Push(nToken.tipo + " ")
	} else if regra == "10" {
		// TIPO→ real
		tokens[0].tipo = "double"
		nToken.tipo = tokens[0].tipo
		linCodigoObjeto.Push(nToken.tipo + " ")
	} else if regra == "11" {
		// TIPO→ lit
		tokens[0].tipo = "literal"
		nToken.tipo = tokens[0].tipo
		linCodigoObjeto.Push(nToken.tipo + " ")
	} else if regra == "13" {
		// ES→ leia id;
		id := tokens[1]
		if id.tipo == "" {
			HandleError("Variável não encontrada", id.lin, id.col)
			return nil, true
		}
		var cmd = map[string]string{
			"literal": "scanf(\"%s\"," + id.lexema + ")\n",
			"int":     "scanf(\"%d\",&" + id.lexema + ")\n",
			"double":  "scanf(\"%lf\",&" + id.lexema + ")\n",
		}
		linCodigoObjeto.Push(cmd[id.tipo])

	} else if regra == "14" {
		// ES→ escreva ARG; - INCREMENTADO
		arg := tokens[1]
		if arg != nil {
			if arg.classe == "id" {
				var cmd = map[string]string{
					"literal": "printf(\"%s\"," + arg.lexema + ")\n",
					"int":     "printf(\"%d\"," + arg.lexema + ")\n",
					"double":  "printf(\"%lf\"," + arg.lexema + ")\n",
				}
				linCodigoObjeto.Push(cmd[arg.tipo])
			} else {
				linCodigoObjeto.Push(fmt.Sprintf("printf(%s);\n", tokens[1].lexema))
			}
		}

	} else if regra == "15" {
		// ARG→ literal
		tokens[0].tipo = "literal"
		nToken.set(tokens[0])
	} else if regra == "16" {
		//  ARG→ num
		tokens[0].tipo = getTipoNumero(tokens[0])
		nToken.set(tokens[0])
	} else if regra == "17" {
		// ARG→ id
		id := tokens[0]
		if id.tipo == "" {
			HandleError("Erro: Variável não declarada", id.lin, id.col)
			return nil, true
		}
		nToken.set(tokens[0])
	} else if regra == "19" {
		// CMD→ id rcb LD;
		id := tokens[0]
		if id.tipo == "" {
			HandleError("Erro: Variável não declarada", id.lin, id.col)
			return nil, true
		}
		ld := tokens[2]
		if ld.tipo == "" {
			HandleError("Erro: Variável não declarada", ld.lin, ld.col)
			return nil, true
		}
		if ld.tipo != id.tipo {
			HandleError("Erro: Tipos diferentes para atribuição", id.lin, id.col)
			return nil, true
		}
		linCodigoObjeto.Push(fmt.Sprintf("%s=%s;\n", id.lexema, ld.lexema))
	} else if regra == "20" {
		// LD→ OPRD opm OPRD
		oprd1 := tokens[0]
		opm := tokens[1]
		oprd2 := tokens[2]

		if oprd1.tipo == oprd2.tipo && oprd1.tipo != "literal" {
			nToken.lexema = fmt.Sprintf("T%d", len(varTempCodigoObjeto))
			nToken.tipo = oprd1.tipo
			varTempCodigoObjeto.Push(fmt.Sprintf("%s %s;\n", nToken.tipo, nToken.lexema))
			linCodigoObjeto.Push(fmt.Sprintf("%s=%s%s%s;\n", nToken.lexema, oprd1.lexema, opm.lexema, oprd2.lexema))
		} else {
			HandleError("Erro: Operandos com tipos incompatíveis", oprd1.lin, oprd1.col)
			return nil, true
		}

	} else if regra == "21" {
		// LD→ OPRD
		nToken.set(tokens[0])

	} else if regra == "22" {
		// OPRD→ id
		id := tokens[0]
		if id.tipo == "" {
			HandleError("Erro: Variável não declarada", id.lin, id.col)
			return nil, true
		}
		nToken.set(id)

	} else if regra == "23" {
		// OPRD→ num
		tokens[0].tipo = getTipoNumero(tokens[0])
		nToken.set(tokens[0])
	} else if regra == "25" {
		// COND→ CAB CP
		linCodigoObjeto.Push("}\n")

	} else if regra == "26" {
		// CAB→ se (EXP_R) então
		linCodigoObjeto.Push(fmt.Sprintf("if(%s)\n", tokens[2].lexema))
		linCodigoObjeto.Push("{\n")
	} else if regra == "27" {
		// EXP_R→ OPRD opr OPRD
		oprd1 := tokens[0]
		opr := tokens[1]
		oprd2 := tokens[2]

		if oprd1.tipo == oprd2.tipo && oprd1.tipo != "literal" {
			nToken.lexema = fmt.Sprintf("T%d", len(varTempCodigoObjeto))
			nToken.tipo = oprd1.tipo
			varTempCodigoObjeto.Push(fmt.Sprintf("%s %s;\n", nToken.tipo, nToken.lexema))
			linCodigoObjeto.Push(fmt.Sprintf("%s=%s%s%s;\n", nToken.lexema, oprd1.lexema, opr.lexema, oprd2.lexema))
		} else {
			HandleError("Erro: Operandos com tipos incompatíveis", oprd1.lin, oprd1.col)
			return nil, true
		}

	} else if regra == "33" {
		// TODO IMPRIMINDO ERRADO :(
		// R → facaAte (EXP_R) CP_R
		linCodigoObjeto.Push(fmt.Sprintf("while(%s)\n", tokens[2].lexema))
		linCodigoObjeto.Push("{\n")
	} else if regra == "34" {
		// CP_R→ COND CP_R
		linCodigoObjeto.Push("}\n")
	}
	return &nToken, false
}

func getTipoNumero(input *Token) string {
	if strings.Contains(input.lexema, ".") {
		return "double"
	} else {
		return "int"
	}
}
