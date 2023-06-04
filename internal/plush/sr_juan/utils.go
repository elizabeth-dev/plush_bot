package sr_juan

import (
	"math/rand"
	"strconv"
)

func AgeGenerator() string {
	age := rand.Intn(8-1) + 1

	if age == 1 {
		return "Tengo " + strconv.Itoa(age) + " añito"
	} else {
		return "Tengo " + strconv.Itoa(age) + " añitos"
	}
}
