package helloworld

const (
	SayHello     = "Hello, "
	HelloSpanish = "Hola, "
	HelloFrench  = "Bonjour, "
)

type Lang string

const (
	FR Lang = "French"
	ES Lang = "Spanish"
)

func Hello(name string, lang Lang) string {
	if name == "" {
		name = "World"
	}

	if lang == ES {
		return HelloSpanish + name
	}

	if lang == FR {
		return HelloFrench + name
	}
	return SayHello + name
}
