package main

import "fmt"

type Human struct {
	name string
	sex  string
	age  int8
}

func (h *Human) Say() {
	fmt.Printf("Hi, i am %s. I am %d years old. My gender is %s.\n", h.name, h.age, h.sex)
}

type Action struct {
	Human
	sick bool
}

func (a *Action) TakePill() {
	if a.sick {
		fmt.Printf("Taking pills, %s not sick anymore\n", a.name)
		a.sick = false
	} else {
		fmt.Printf("My name is %s and I am completely healthy.\n", a.name)
	}
}

func main() {
	man := &Action{
		Human: Human{
			name: "Really Man Name",
			sex:  "male",
			age:  37,
		},
		sick: true,
	}
	uma := &Action{
		Human: Human{
			name: "Oguri Cap",
			sex:  "female",
			age:  15,
		},
		sick: false,
	}
	narrator := &Action{
		Human: Human{
			name: "Tyler Durden",
			sex:  "unknown",
			age:  26,
		},
		sick: true,
	}

	man.Say()
	man.TakePill()
	uma.Say()
	uma.TakePill()
	narrator.Say()
	narrator.TakePill()
	narrator.TakePill()
}
