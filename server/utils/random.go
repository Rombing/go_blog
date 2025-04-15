package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func GenerateVerificationCode(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%0*d", length, r.Intn(int(math.Pow10(length))))
}
