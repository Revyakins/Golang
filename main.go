package main

import "fmt"


func main() {

var wait = make(chan struct{} )

 go func() {
 
 fmt.Println("Hello world")
 wait <- struct {}{} 
 }()

 <- wait
}
