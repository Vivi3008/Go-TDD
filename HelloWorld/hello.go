package helloworld

const (
	SayHello = "Hello "
)

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return SayHello + name
}
