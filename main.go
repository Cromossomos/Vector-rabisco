package main
 
 import "fmt"
 
 func main() {
 var numeros[5]int
 var soma int

 fmt.Print("digite um numero inteiros, ")

 for i := 0; i < 5; i++ {
	fmt.Printf("Número %d: ", i+1)
	fmt.Scan(&numeros[i])
	soma += numeros[i]

 }

fmt.Printf("a soma dos numeros é %d", soma)

  }