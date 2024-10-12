package assert

type config struct {
	PanicOnFail bool
	PrintOnFail bool
}

var Config config = config{
	PanicOnFail: true,
	PrintOnFail: false,
}
