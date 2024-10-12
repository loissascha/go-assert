package assert

import "fmt"

// assert that value should be true. If it's not, a panic will be executed
func True(value bool, keyword string) {
	if !value {
		failed("assert.True", keyword)
	}
}

func failed(funcName string, keyword string) {
	if Config.PanicOnFail {
		panic(fmt.Sprintf("%s failed in %s", funcName, keyword))
	}
}
