package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	conteudo,_ := ioutil.ReadFile("test/mgol_files/simple.alg")
	fmt.Println(string(conteudo))
}