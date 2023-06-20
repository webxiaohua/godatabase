package main

import "fmt"

type Animal struct {
	Name string
}

func (animal *Animal) Eat()  {
	fmt.Printf("Animal %s is eating.\n",animal.Name)
}

type Dog struct {
	Animal
}

func (dog *Dog) Eat()  {
	fmt.Printf("Dog %s is eating.\n",dog.Name)
}

type Cat struct {
	Animal
}
func (cat *Cat) Eat()  {
	fmt.Printf("Cat %s is eating.\n",cat.Name)
}

func Doit(animal interface{})  {
	t,ok := animal.(Animal)
	if ok{
		t.Eat()
	}else{
		fmt.Println("type convert error.")
	}
}

func main()  {
	myDog := &Dog{Animal{Name:"aaa"}}
	Doit(myDog)

	fmt.Printf("%s",myDog)
}