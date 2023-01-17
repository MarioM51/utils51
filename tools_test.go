package toolkit

import (
	"fmt"
	"testing"
)

func Test_Tools_RandomString(t *testing.T) {
	var tools Tools

	const length = 10
	randomString1 := tools.RandomString(length)

	if len(randomString1) != length {
		t.Error(fmt.Sprintf("wrong length must be %v and it was %v \n", length, len(randomString1)))
	}

	randomString2 := tools.RandomString(length)
	if randomString1 == randomString2 {
		t.Error(fmt.Sprintf("Random strings are equals %s == %s \n", randomString1, randomString2))
	}

}
