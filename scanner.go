package main

var tokens = map[string]string{
	"NUM":        "Num",
	"LITERAL":    "Literal",
	"ID":         "id",
	"COMENTARIO": "Comentário",
	"EOF":        "EOF",
	"OPR":        "OPR",
	"RCB":        "RCB",
	"OPM":        "OPM",
	"AB_P":       "AB_P",
	"FC_P":       "FC_P",
	"PT_V":       "PT_V",
	"ERRO":       "ERRO",
	"VIR":        "Vir",
	"S":          "/s",
}

var estados = map[string]int{
	"INICIO":                         0,
	"NUMERO_INTEIRO_FINAL":           1,
	"NUMERO_REAL_FINAL":              2,
	"NUMERO_CIENTIFICO_INCOMPLETO_1": 3,
	"NUMERO_CIENTIFICO_INCOMPLETO_2": 4,
	"NUMERO_CIENTIFICO_FINAL":        5,
	"LITERAL_INCOMPLETO":             6,
	"LITERAL_FINAL":                  7,
	"COMENTARIO_INCOMPLETO":          8,
	"COMENTARIO_FINAL":               9,
	"ID_FINAL":                       10,
	"OPR_FINAL_1":                    11,
	"RCB_FINAL":                      12,
	"OPR_FINAL_2":                    13,
	"OPR_FINAL_3":                    14,
	"OPM_FINAL":                      15,
	"AB_P_FINAL":                     16,
	"FC_P_FINAL":                     17,
	"PT_V_FINAL":                     18,
	"EOF_FINAL":                      19,
	"VIR_FINAL":                      20,
}

var classePorEstado = map[int]string{
	0:  "",
	1:  "NUM",
	2:  "NUM",
	3:  "",
	4:  "",
	5:  "NUM",
	6:  "",
	7:  "LITERAL",
	8:  "",
	9:  "COMENTARIO",
	10: "ID",
	11: "OPR",
	12: "RCB",
	13: "OPR",
	14: "OPR",
	15: "OPM",
	16: "AB_P",
	17: "FC_P",
	18: "PT_V",
	19: "EOF",
	20: "VIR",
}

var ignorar = []string{"\n", " ", "\t", ""}
var alfabeto = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}
var inteiros = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var operadoresAritmeticos = []string{"+", "-", "*", "/"}

func identificaPalavraReservada(entrada string, indexInicio int) (lexema string, indexFim int) {
	var novoIndexInicio int
	for novoIndexInicio = indexInicio; novoIndexInicio < len(entrada); novoIndexInicio++ {
		if !stringInSlice(string(entrada[novoIndexInicio]), ignorar) {
			break
		}
	}

	for _, reservada := range palavrasReservadas {
		if len(entrada) >= len(reservada)+novoIndexInicio &&
			reservada == entrada[novoIndexInicio:novoIndexInicio+len(reservada)] &&
			(novoIndexInicio+len(reservada) == len(entrada) || stringInSlice(entrada[novoIndexInicio:novoIndexInicio+len(reservada)], ignorar)) {
			lexema = reservada
			indexFim = novoIndexInicio + len(reservada)
			return
		}
	}
	indexFim = novoIndexInicio
	return
}

func identificaTokenSemPalavraReservada(entrada string, indexInicio int, isUltimaLinha bool) (token Token, indexFim int) {
	var estado = estados["INICIO"]
	var lexema = ""
	var totalString = len(entrada)

	if indexInicio == totalString {
		indexFim = indexInicio
		if isUltimaLinha {
			token = Token{tokens["EOF"], tokens["EOF"], ""}
		} else {
			token = Token{}
		}
		return
	}

	for indexFim = indexInicio; indexFim < totalString; indexFim++ {
		var value = string(entrada[indexFim])
		if estado == estados["INICIO"] {
			if stringInSlice(value, ignorar) {
				estado = estados["INICIO"]
			} else if value == ";" {
				estado = estados["PT_V_FINAL"]
				lexema = lexema + value
			} else if value == "," {
				estado = estados["VIR_FINAL"]
				lexema = lexema + value
			} else if value == ")" {
				estado = estados["FC_P_FINAL"]
				lexema = lexema + value
			} else if value == "(" {
				estado = estados["AB_P_FINAL"]
				lexema = lexema + value
			} else if stringInSlice(value, operadoresAritmeticos) {
				estado = estados["OPM_FINAL"]
				lexema = lexema + value
			} else if value == "=" {
				estado = estados["OPR_FINAL_3"]
				lexema = lexema + value
			} else if value == ">" {
				estado = estados["OPR_FINAL_2"]
				lexema = lexema + value
			} else if value == "<" {
				estado = estados["OPR_FINAL_1"]
				lexema = lexema + value
			} else if stringInSlice(value, alfabeto) {
				estado = estados["ID_FINAL"]
				lexema = lexema + value
			} else if value == "{" {
				estado = estados["COMENTARIO_INCOMPLETO"]
				lexema = lexema + value
			} else if stringInSlice(value, inteiros) {
				estado = estados["NUMERO_INTEIRO_FINAL"]
				lexema = lexema + value
			} else if value == "\"" {
				estado = estados["LITERAL_INCOMPLETO"]
				lexema = lexema + value
			} else {
				estado = -1
				lexema = lexema + value
			}
		} else if estado == estados["NUMERO_INTEIRO_FINAL"] {
			if stringInSlice(value, inteiros) {
				lexema = lexema + value
			} else if value == "." {
				estado = estados["NUMERO_REAL_FINAL"]
				lexema = lexema + value
			} else if value == "E" || value == "e" {
				estado = estados["NUMERO_CIENTIFICO_INCOMPLETO_1"]
				lexema = lexema + value
			} else {
				break
			}
		} else if estado == estados["NUMERO_REAL_FINAL"] {
			if stringInSlice(value, inteiros) {
				lexema = lexema + value
			} else if value == "E" || value == "e" {
				estado = estados["NUMERO_CIENTIFICO_INCOMPLETO_1"]
				lexema = lexema + value
			} else {
				break
			}
		} else if estado == estados["NUMERO_CIENTIFICO_INCOMPLETO_1"] {
			if value == "+" || value == "-" {
				estado = estados["NUMERO_CIENTIFICO_INCOMPLETO_2"]
				lexema = lexema + value
			} else if stringInSlice(value, inteiros) {
				estado = estados["NUMERO_CIENTIFICO_FINAL"]
				lexema = lexema + value
			} else {
				break
			}
		} else if estado == estados["NUMERO_CIENTIFICO_INCOMPLETO_2"] {
			if stringInSlice(value, inteiros) {
				estado = estados["NUMERO_CIENTIFICO_FINAL"]
				lexema = lexema + value
			} else {
				break
			}
		} else if estado == estados["NUMERO_CIENTIFICO_FINAL"] {
			if stringInSlice(value, inteiros) {
				lexema = lexema + value
			} else {
				break
			}
		} else if estado == estados["LITERAL_INCOMPLETO"] {
			if value == "\"" {
				estado = estados["LITERAL_FINAL"]
				lexema = lexema + value
			} else {
				lexema = lexema + value
			}
		} else if estado == estados["COMENTARIO_INCOMPLETO"] {
			if value == "}" {
				estado = estados["COMENTARIO_FINAL"]
				lexema = lexema + value

			} else {
				lexema = lexema + value
			}
		} else if estado == estados["ID_FINAL"] {
			if stringInSlice(value, inteiros) || stringInSlice(value, alfabeto) {
				lexema = lexema + value
			} else {
				break
			}
		} else if estado == estados["OPR_FINAL_1"] {
			if value == "-" {
				estado = estados["RCB_FINAL"]
				lexema = lexema + value
			} else if value == "=" || value == ">" {
				estado = estados["OPR_FINAL_3"]
				lexema = lexema + value
			} else {
				break
			}
		} else if estado == estados["OPR_FINAL_2"] {
			if value == "=" {
				estado = estados["OPR_FINAL_3"]
				lexema = lexema + value
			} else {
				break
			}
		} else {
			break
		}

	}

	if classePorEstado[estado] != "" {
		token = Token{lexema, tokens[classePorEstado[estado]], ""}
	} else {
		token = Token{lexema, "ERRO4", ""}
	}

	if estado == estados["NUMERO_CIENTIFICO_INCOMPLETO_1"] || estado == estados["NUMERO_CIENTIFICO_INCOMPLETO_2"] {
		token = Token{lexema, "ERRO1", ""}
	}

	if estado == estados["LITERAL_INCOMPLETO"] {
		token = Token{lexema, "ERRO2", ""}
	}

	if estado == estados["COMENTARIO_INCOMPLETO"] {
		token = Token{lexema, "ERRO3", ""}
	}

	return
}

func Scanner(entrada string, indexInicio int, isUltimaLinha bool) (token Token, indexFim int) {
	lexema, indexFim := identificaPalavraReservada(entrada, indexInicio)
	if lexema != "" {
		token = tabelaSimbolos[lexema]
		return
	}

	token, indexFim = identificaTokenSemPalavraReservada(entrada, indexInicio, isUltimaLinha)
	if token.classe == "id" {
		if tokenSalvo, ok := tabelaSimbolos[token.lexema]; ok {
			token = tokenSalvo
		} else {
			tabelaSimbolos[token.lexema] = token
		}
		return
	}
	return
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
