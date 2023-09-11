package main
import "fmt"

func main(){
	total := sum(10,1);
	fmt.Println("Resultado", total);
}

func sum(a int, b int) int {
	return a + b;
}

func sub(a int, b int) int {
	return a - b;
}

func times(a int, b int) int {
	return a * b;
}

func division(a int, b int) int {
	return a / b;
}