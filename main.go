package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func getRandoms(n int)[]int{
	r:=[]int{}
	for{
		n:=rand.Intn(9)
		if !contains(r,n){
			r=append(r,n)
		}
		if len(r)==4{
			break
		}
	}
	return r
}
func getBC(random []int,user []int)(int,int){
	b,c:=0,0
	for i:=0; i<len(user); i++ {
		if random[i]==user[i]{
			b=b+1
		} else{
			if contains(random,user[i]){
				c=c+1
			}
		}
	}
	return b,c
}
func isUnieqNumbers(a []int)bool{
	n:=[]int{}
	for i:=0; i<=9; i++{
		n=append(n,0)
	}
	for i:=0; i<len(a); i++ {
		c:=a[i]
		n[c]++
		if n[c]==2{
			return false
		}
	}
	return true
}
func main(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to bulls and cows game!")
	fmt.Println("Rules:")
	fmt.Println("You should guess number (4 digits)")
	fmt.Println("The number can be started from zero")
	for{
		//init game
		fmt.Println("New game started")
		n:=getRandoms(4)
		tries:=0
		//read loop
		for{
			text:=input("Enter the number: ")
			tries++;
			user:=[]int{}
			if text=="q"{
				quiet()
			}
			if len(text)!=4{
				fmt.Println("Incorrect format")
				fmt.Println("Number should consists of four unique digits")
				continue
			}
			for i:=0;i<4;i++{
				number,_:=strconv.ParseInt(string(text[i]),10,32)
				user=append(user,int(number))
			}
			if !isUnieqNumbers(user){
				fmt.Println("Incorrect format")
				fmt.Println("Number should consists of four UNIQUE digits")
				continue
			}
			b,c:=getBC(n,user)
			fmt.Printf("Bulls: %d\nCows: %d\n",b,c)
			if b==4{
				fmt.Println("You win!")
				fmt.Printf("Tries: %d\n",tries)
				fmt.Printf("Quit? (q)\nIf you want to start new game press Enter")
				q:=input(" ")
				if q=="q"{
					quiet()
				} else{
					break
				}
			}
		}
	}
}