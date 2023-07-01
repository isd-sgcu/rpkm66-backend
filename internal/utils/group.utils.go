package utils

import "time"

func GenToken(id string) string {
	return Hash([]byte(time.Now().String() + id))
}
