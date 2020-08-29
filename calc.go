package main

func main(){

	fmt.Println(add(2,3))
	fmt.Println(sub(2,3))

}

func add(num1, num2 int) int {
	return num1 + num2
}

func sub(num1, num2 int) int{
	return num1 - num2
}
