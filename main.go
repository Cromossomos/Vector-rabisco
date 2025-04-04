package main

import (
"fmt" 
"strings"
"sort"
)
func main(){

	greeting := "Opa meu amigos"
	fmt.Println(strings.Contains(greeting, "amigos"))
	fmt.Println(strings.ReplaceAll(greeting, "Opa", "Oi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "meu"))
	fmt.Println(strings.Split(greeting, "amigos"))
	ages := []int {50, 80, 10}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 50)
	fmt.Println(index)
	names := []string{"Caroline", "Maicon", "Diego"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Caroline"))
	
}