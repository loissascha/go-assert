package assert

import "fmt"

// assert that value should be true. If it's not, a panic will be executed
func True(value bool, keyword string) {
	if !value {
		panic(fmt.Sprintf("assert.True failed in %s", keyword))
	}
}
