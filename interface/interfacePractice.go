package main

// interfaces are needed when there is some common functionality exists between different objects for example suzuki and honda both manufactures bikes (hence it is common between them)
import "fmt"

type Animal interface {
	Says(string) string
	Legs(int) string
}

type Rabbit struct {
	eyes int
	ears int
}
type Python struct {
	eyes int
	tail int
}

func (r Rabbit) Says(sound string) string {
	return fmt.Sprintf("It makes a %v", sound)
}
func (r Rabbit) Legs(totalLegs int) string {
	return fmt.Sprintf("It has total %v legs", totalLegs)
}

func CommonFeatureForAllAnimals(animal Animal) string {
	return "It is a Breed of animals"
}

func main() {
	duffy := Rabbit{eyes: 2, ears: 2}
	fmt.Println(duffy.Says("Rabbit Sound")) // It makes a Rabbit Sound

	fmt.Println(CommonFeatureForAllAnimals(duffy)) // It is a Breed of animals :- we can see that this function only accept animal interface then why it has accepted our Rabbit struct (duffy) , so this is because Rabbit struct has implementd all the functions of Animal interface hence it has access to this function

	boa := Python{eyes: 2, tail: 1}
	_ = boa
	// fmt.Println(CommonFeatureForAllAnimals(boa)) // if we un comment this line then we we will get an error because python struct have not implements the functions of the animal interface

}
