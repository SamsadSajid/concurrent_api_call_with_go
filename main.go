package main

import "fmt"

func main(){
	router := Controller()

	fmt.Println(router)

	router.Run(":8000")
}
