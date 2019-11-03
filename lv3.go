package main

import "fmt"

func main(){
	var i int //yaochudeshu//
	var b int //shu//

	for b=2;b<10000;b++{
		for i=2;i<=b;i++ {
			if b==i{
				fmt.Println(b)
			}
			if b%i ==0{
				break
			}
		}
	}




}
