package main

import "fmt"

func main(){
	var a,b,c int
	fmt.Scanf("%d %d %d",&a,&b,&c)
	if a > b && a > c{
	fmt.Println("A is biggest")
	}else if b > a && b >c{
	fmt.Println("B is biggest")
	}else{
	fmt.Println("C is biggest")
	}	
}