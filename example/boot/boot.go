package boot

import (
	"SimpleSkeleton/app"
	"context"
	"fmt"
)

func Boot(process app.App) {
	process.CreateContainer(func() app.Container {
		ctx := context.Background()
		con := &ExampleContainer{
			ctx:     ctx,
			flowMap: make(map[string]app.Flow),
		}
		SetRoutes(con)
		for k := range con.flowMap {
			fmt.Println(
				fmt.Sprintf(
					"i, route: %s, in container %s, "+
						"send request to /%s for response",
					k,
					con.GetName(),
					k))
		}
		return con
	})

	process.ListenTCP().Run()
}
