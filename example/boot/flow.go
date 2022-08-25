package boot

import (
	"fmt"
)

type ExampleController struct {
}

func NewExample() *ExampleController {
	return &ExampleController{}
}

func (c *ExampleController) GetName() string {
	return "exampleController"
}

func (c *ExampleController) Execute(msg string) interface{} {
	fmt.Println(fmt.Sprintf("i, %s, executed.", c.GetName()))
	fmt.Println(fmt.Sprintf("message print: %s", msg))
	fmt.Println([]byte(fmt.Sprintf("message print: %s", msg)))
	return c
}
