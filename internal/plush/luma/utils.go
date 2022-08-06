package luma

import (
	"math/rand"
	"strings"
)

func RawrGenerator() string {
	return "Rawr" + strings.Repeat("r", rand.Intn(4))
}
