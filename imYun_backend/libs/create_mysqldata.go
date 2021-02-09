package libs

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomString2(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func GetRandomType(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	times := string(time.Now().Format("20060102150405"))
	times = string(times[len(times)-1])
	switch times {
	case "0","1":
		return fmt.Sprintf("%x"+".docx", randBytes)
	case "2", "3":
		return fmt.Sprintf("%x"+".elsx", randBytes)
	case "4", "5":
		return fmt.Sprintf("%x"+".pptx", randBytes)
	case "6", "7":
		return fmt.Sprintf("%x"+".doc", randBytes)
	case "8", "9":
		return fmt.Sprintf("%x"+".ppt", randBytes)
	}
	return fmt.Sprintf("%x"+".els", randBytes)
}

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max - min) + min
	return randNum
}