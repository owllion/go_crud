package handler

import (
	"math/rand"
)

func generateInt64ID() int64 {
	return rand.Int63()
}