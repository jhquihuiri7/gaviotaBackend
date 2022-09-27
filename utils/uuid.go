package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

func GenerateID() string {
	return uuid.NewV4().String()
}

func GenerateReserve() string{
	number := 1
	letter := ""
	date := ""

	letters := []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}

	for number <= 9{
		number = rand.Intn(99)
	}
	for i:= 0; i <3; i++ {
		letter += letters[rand.Intn(len(letters))]
	}
	if time.Now().Day() < 10 {
		date = "0"+ fmt.Sprint(time.Now().Day())
	}else {
		date = fmt.Sprint(time.Now().Day())
	}

	return fmt.Sprint("GF",number,"G",date,letter)
}