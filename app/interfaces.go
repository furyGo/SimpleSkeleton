package app

// App Set default configuration and bootstrap
type App interface {
	CreateContainer(container func() Container) Container
	ListenTCP() (app App)
	ListenCommands() (app App)
	Run()
}

// Container All flows in the container share the same context
type Container interface {
	GetName() string
	AddFlow(name string, flow func() Flow)
	Receive(route string, message <-chan string) interface{}
}

// Flow A flow controls data stream and channels' order
type Flow interface {
	GetName() string
	Execute(message string) interface{}
}
