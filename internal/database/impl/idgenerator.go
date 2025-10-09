package impl

import "math/rand/v2"

func GenerateId() uint64 {
	return rand.Uint64()
}
