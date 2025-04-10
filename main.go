package main
 
 import "fmt"
 
 func main() {
	age := 45
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Menor que 30 anos")
	} else if age < 40 {
		fmt.Println("Menor que 40 anos")
	} else {
		fmt.Println("Não é menor que 40 anos")
	}

	names := []string{"Kaiser", "Balu", "Ivete", "Dante", "Arthur", "Joui"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continue pós a posição", index, "e valor", value)
			continue
		}
		if index > 2 {
			fmt.Println("Sair após", index)
			break
		}
		fmt.Println("Valor: ", value)
	}
 }