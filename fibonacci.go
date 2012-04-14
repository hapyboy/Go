package main

import (
"fmt"
)

func main(){

	x := 0
	
	d := 1
	fmt.Printf("1\t")
	for i :=0;i<10;i++{
		
		d += x
		fmt.Printf("%v\t",d)
		x = d
		if i/10 == 0{
			fmt.Print('\n')
		}
	}
}


