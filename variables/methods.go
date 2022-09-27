package variables

import (
	"math/rand"
)

func (r *Reserve) RandomReserve(records [][]string) {
	var age int
	for {
		age = rand.Intn(64)
		if age > 18 {
			break
		}
	}
	r.Passenger = records[rand.Intn(199)][0]
	r.Age = age
}
func (r *Reserve)FormatTime(date string){

}