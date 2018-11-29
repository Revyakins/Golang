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

type Man struct {
	Human
}

func (g Man) setHumanInfo(name string, age int32, city string, address string) Man {
	return Man{
		Human{
			name,
			age,
			city,
			address,
		},
	}
}

func PeopleTalk(humans []People) {
	for _, s := range humans {
		fmt.Println(s.Talk())
	}
}

type crowd []People

func main() {

	sergey := new(Man).setHumanInfo("Sergey", 26, "Kyiv", "Ahmatova")
	tema := new(Man).setHumanInfo("Tema", 29, "Kyiv", "test")

	cr := make(crowd, 0)

	cr = append(cr, sergey, tema)

	PeopleTalk(cr)

}
