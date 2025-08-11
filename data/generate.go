package data

import (
	"fmt"
	"math/rand"
)

func Generate(dataType string) any {
	switch dataType {
	case TYPE_NAME:
		return generateName()
	case TYPE_DATE:
		return generateDate()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_PHONE:
		return generatePhone()

	}
	return ""

}

func generateName() string {
	nameLen := len(name)

	index := rand.Intn(nameLen) // Intn -> generate angka random 0 hingga (nameLen -1)
	return name[index]
}
func generateDate() string {
	return ""
}
func generateAddress() string {
	streetLen := len(SUBTYPE_STREET)
	cityLen := len(SUBTYPE_CITY)

	streetIdx := rand.Intn(streetLen)
	cityIdx := rand.Intn(cityLen)
	number := rand.Intn(100)

	return fmt.Sprintf("JL. %s No. %d, %s", address[SUBTYPE_STREET][streetIdx], number, address[SUBTYPE_CITY][cityIdx])
}
func generatePhone() string {
	return ""
}
