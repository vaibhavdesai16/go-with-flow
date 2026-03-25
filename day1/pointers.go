package day1

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func updateAge(user *User) {
	user.Age = 30
}

func (u User) AgeIncrementVoid() {
	u.Age++
}

func (u *User) AgeIncrement() {
	u.Age++
}

func PlayingWithPointers() {
	u := User{Name: "Sachin", Age: 29}
	fmt.Printf("age of %s is %d", u.Name, u.Age)
	updateAge(&u)
	fmt.Printf("age of %s is %d \n", u.Name, u.Age)

	// value receiver
	//  AgeIncrement will receive copy of struct and orinal remains unchanged
	u1 := User{Name: "Virat", Age: 35}
	u1.AgeIncrementVoid()
	fmt.Println(u1.Age)

	u1.AgeIncrement()
	fmt.Println(u1.Age)

}
