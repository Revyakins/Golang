package linked_list

import (
	"fmt"
	"testing"
)

func ExampleList() {

	list := new(List)

	list.Insert(5)
	list.Insert(4)

	fmt.Print("Display: ")
	fmt.Println(list.Display())
	fmt.Print("Reverse: ")
	list.Reverse()
	fmt.Println(list.Display())

	//Output:
	//Display: 4 ->5 ->
	//Reverse: 5 ->4 ->
}

func TestList_Reverse(t *testing.T) {

	list := new(List)
	list.Insert(8)
	list.Insert(12)

	list.Reverse()

	if result := list.head.value; result != 8 {
		t.Error(
			"For", "Reverse",
			"expected", 8,
			"got", result,
		)
	}
}
