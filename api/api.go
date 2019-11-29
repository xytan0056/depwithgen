package api

import (
	"fmt"
	c "github.com/xytan0056/depwithgen/.gen/client"
)

func API() {
	fmt.Println("I'm depwithgen " + c.Client("x"))
}