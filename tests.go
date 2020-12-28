package main

import (
	"fmt"
)
func assertEqual(a interface{}, b interface{})bool{
	return a==b
}
//BCtest tests getBC
func BCtest(x []int,y []int)int{
	b,c:=getBC(x,y)
	if b==3 && c==0{
		return 1
	}
	return 0
}
func testUnieq(a[] int)int{
	if isUnieqNumbers(a){
		return 1
	}
	return 0
}
func testAll(){
	//count:=2
	success:=0
	x:=[]int{1,2,3,4}
	y:=[]int{7,2,3,4}
	z:=[]int{1,1,5,4}
	success+=BCtest(x,y)
	success+=testUnieq(z)
	fmt.Println(success)
}