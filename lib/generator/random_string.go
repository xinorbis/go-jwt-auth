package generator

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func GenerateUniqueHash() string {
	data := strconv.FormatInt(time.Now().UnixNano()+100, 10)
	hasher := sha256.New()
	hasher.Write([]byte(data))
	result := fmt.Sprintf("%x", hasher.Sum(nil))

	fmt.Println("data:", data)

	return result
}
