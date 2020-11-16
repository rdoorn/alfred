package integration

type Plugin interface {
	Name() string
	Version() string
	Subscribe(e <-chan Event) error
	Publish(*Event) error
	Shutdown()
}
