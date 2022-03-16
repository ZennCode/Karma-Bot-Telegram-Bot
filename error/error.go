package error

import "log"

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}