package boot

import (
	"context"
	"fmt"
	"github.com/furyGo/SimpleSkeleton/app"
	"log"
)

type ExampleContainer struct {
	ctx     context.Context
	flowMap map[string]app.Flow
}

func (ctx *ExampleContainer) GetName() string {
	return "exampleContainer"
}

func (ctx *ExampleContainer) AddFlow(name string, flow func() app.Flow) {
	ctx.flowMap[name] = flow()
}

func (ctx *ExampleContainer) Receive(
	route string,
	message <-chan string) interface{} {
	flow := ctx.flowMap[route]
	if flow != nil {
		log.Println(
			fmt.Sprintf(
				"container %s contains flow %s, executing",
				ctx.GetName(),
				flow.GetName()))
		return flow.Execute(<-message)
	}
	return "404"
}
