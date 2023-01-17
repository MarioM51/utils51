package toolkit

import "crypto/rand"

const randomStringSource = "qwertyuiopasdfghjklzxcvbnm-1234567890_QAZWXEDCRFVTGBYHNUJMIKOLP+"

type Tools struct {
}

func (t *Tools) RandomString(length int) string {
	result := make([]rune, length)
	source := []rune(randomStringSource)
	for i := range result {
		randomGen, _ := rand.Prime(rand.Reader, len(source))
		randomNum := randomGen.Uint64()
		limit := uint64(len(source))
		result[i] = source[randomNum%limit]
	}

	return string(result)
}
