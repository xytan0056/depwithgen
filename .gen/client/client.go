package client

import "fmt" 

func Client(id string) string {
	fmt.Println("I'm generated client") 
	return "client@"+id 
}