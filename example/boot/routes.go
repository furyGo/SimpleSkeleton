package boot

import "SimpleSkeleton/app"

func SetRoutes(container app.Container) {

	container.AddFlow("flow", func() app.Flow {
		return NewExample()
	})
}
