package main

import (
	"fmt"
	"bufio"
	"os"
)

func contains(a []int,element int)bool{
	for i:=0;i<len(a);i++ {
		if a[i]==element {
			return true;
		}
	}
	return false;
}
func input(message string)string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	text, _ := reader.ReadString('\n')
	return text[:len(text)-1] //delete \n
}
func quiet(){
	fmt.Println("Bye bye")
	os.Exit(0)
}