package main
 
 import "fmt"
 
 func dadosPessoa(nome string, idade int) (int, string){
	if idade >= 18{
		return idade, "Maior de idade"
	}
	return idade, "Menor de idade"
 }

 func main() {
	nome := "Cesinha"
	idade := 16

	idadeRetornada, status := dadosPessoa(nome, idade)
	fmt.Printf("Nome: %s\nIdade: %d\nStatus: %s\n", nome, idadeRetornada, status)
 }