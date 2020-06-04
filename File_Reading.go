package main

import( 
	"fmt"
	"io/ioutil"
)

func main(){
	data, err :=  ioutil.ReadFile("C:/Users/Admin/Desktop/Read.txt")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(len(data))
	fmt.Println(string(data))
}
