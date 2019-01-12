package main

import (
	"fmt"
)

type People interface {
	GetName() string
	GetAge() int32
	GetAddress() string
	Talk() string
}

type Human struct {
	Name    string
	Age     int32
	City    string
	Address string
}

func (h Human) GetName() string {
	return h.Name
}

func (h Human) GetAge() int32 {
	return h.Age
}

func (h Human) GetAddress() string {
	return fmt.Sprintf("%s %s\n", h.City, h.Address)

}

func (h Human) Talk() string {
	return fmt.Sprintf("Hello my name is %s ", h.Name)

}

func setHumanInfo(name string, age int32, city string, address string) *Human {
	return &Human{
		name,
		age,
		city,
		address,
	}
}

func PeopleTalk(humans []People) {
	for _, s := range humans {
		fmt.Println(s.Talk())
	}
}

type crowd []People

func main() {

	Sergey := setHumanInfo("Sergey", 26, "Kyiv", "Ahmatova")
	Artem := setHumanInfo("Tema", 29, "Kyiv", "test")

	cr := append(make(crowd, 0), Sergey, Artem)

	PeopleTalk(cr)
}
