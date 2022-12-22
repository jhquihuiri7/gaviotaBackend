package utils

import (
	uuid "github.com/satori/go.uuid"
)

func GenerateID() string {
	return uuid.NewV4().String()
}

func GenerateReserve() string {
	//letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	//reserve := ""
	//for {
	//	number1 := 1
	//	number2 := 1
	//	letter1 := ""
	//	letter2 := ""
	//	letter3 := ""
	//	reserve = ""
	//	for number1 <= 999999 {
	//		number1 = rand.Intn(9999999)
	//	}
	//	for number2 <= 9 {
	//		number2 = rand.Intn(99)
	//	}
	//	for i := 0; i < 3; i++ {
	//		letter1 += letters[rand.Intn(len(letters))]
	//	}
	//	letter2 = letters[rand.Intn(len(letters))]
	//	for i := 0; i < 4; i++ {
	//		letter3 += letters[rand.Intn(len(letters))]
	//	}
//
	//	reserve = fmt.Sprint(letter1, number1, letter2, number2, letter2)
	//	result1 := variables.ReservesGaviotaCollection.FindOne(context.TODO(),bson.D{{"reserve",reserve}})
	//	result2 := variables.ReservesGaviotaCollection.FindOne(context.TODO(),bson.D{{"reserve",reserve}})
	//	if result1.Err() == mongo.ErrNoDocuments && result2.Err() == mongo.ErrNoDocuments {
	//		break
	//	}
	//}
	return uuid.NewV4().String()
}
