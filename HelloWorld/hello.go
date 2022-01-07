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
	return prefixHello(lang) + name
}

func prefixHello(lang Lang) (prefix string) {
	switch lang {
	case ES:
		prefix = HelloSpanish
	case FR:
		prefix = HelloFrench
	default:
		prefix = SayHello
	}
	return
}
