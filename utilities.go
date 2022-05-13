package sample_licensed_project

import "math/rand"

func GetText() string {
	if rand.Int()%2 == 0 {
		return "Here is your text"
	}

	return "Or another text"
}
