package assert

type config struct {
	PanicOnFail bool
}

var Config config = config{
	PanicOnFail: true,
}
