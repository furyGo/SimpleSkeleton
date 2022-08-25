package boot

import "github.com/furyGo/SimpleSkeleton/app"

func SetRoutes(container app.Container) {

	container.AddFlow("flow", func() app.Flow {
		return NewExample()
	})
}
