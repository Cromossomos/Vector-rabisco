package main

import "fmt"

func main(){

nomes := []string{"Enzo Gabriel", "Thiago", "Vini", "Isabella"}

fmt.Println("Dois primeiros:", nomes[:2])
fmt.Println("Dois ultimos nomes:", nomes[len(nomes)-2:])
fmt.Println("Nomes do meio:", nomes[len(nomes)/2:len(nomes)/2+1])

}