package randomMatch

import "math/rand"

func RandomMath(minnumber int, maxnumber int)int {
	// value := rand.Int()%(minnumber+maxnumber-1) + maxnumber
	var value = rand.Int() % (maxnumber)
	return value
}
