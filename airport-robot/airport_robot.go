package airportrobot

// General
type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}

// Italian
type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(visitorName string) string {
	return "Ciao " + visitorName + "!"
}

// Portuguese
type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (p Portuguese) Greet(visitorName string) string {
	return "Olá " + visitorName + "!"
}
