package utils

import "log"

func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func LogError(err error) error {
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
