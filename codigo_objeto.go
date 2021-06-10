package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var linCodigoObjeto Stack
var varTempCodigoObjeto Stack

func inicializarCodigoObjeto() {
	varTempCodigoObjeto = Stack{}
	linCodigoObjeto = Stack{}
}

func gravarCodigoObjeto(caminho string) {
	linhas := Stack{}
	inputs := Stack{}
	inputs.Push("#include<stdio.h>\n")
	inputs.Push("typedef char literal[256];\n")
	inputs.Push("void main(void)\n")
	inputs.Push("{\n")
	inputs.Push("/*----Variaveis temporarias----*/\n")
	for _, lin := range varTempCodigoObjeto {
		inputs.Push(lin)
	}
	inputs.Push("/*------------------------------*/\n")
	for _, lin := range linCodigoObjeto {
		inputs.Push(lin)
	}
	inputs.Push("}\n")

	acc :=""
	for _, input := range inputs {
		acc = acc+input
		if strings.Contains(acc, "\n") {
			linhas.Push(acc)
			acc = ""
		}
	}

	identacao:=0
	fonte := ""
	for _, lin := range linhas {
		if strings.Contains(lin, "}"){
			identacao--
		}
		if strings.Contains(lin, "\n") {
			for i := 0; i < identacao; i++ {
				fonte +="\t"
			}
		}
		fonte +=lin
		if strings.Contains(lin, "{"){
			identacao++
		}
	}

	fmt.Println(fonte)

	err := ioutil.WriteFile(caminho+".c", []byte(fonte), 0644)
	if err != nil {
		panic(err)
	}




}
