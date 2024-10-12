package assert

import "fmt"

func True(value bool, keyword string) {
	if !value {
		failed("assert.True", keyword)
	}
}

func False(value bool, keyword string) {
	if value {
		failed("assert.False", keyword)
	}
}

func NotNil(value any, keyword string) {
	if value == nil {
		failed("assert.NotNil", keyword)
	}
}

func Nil(value any, keyword string) {
	if value != nil {
		failed("assert.Nil", keyword)
	}
}

func String(value string, expected string, keyword string) {
	if value != expected {
		failed(fmt.Sprintf("assert.String [%s]", expected), keyword)
	}
}

func failed(funcName string, keyword string) {
	if Config.PrintOnFail {
		fmt.Println(funcName, "failed in", keyword)
	}
	if Config.PanicOnFail {
		panic(fmt.Sprintf("%s failed in %s", funcName, keyword))
	}
}
