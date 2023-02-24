package util

import (
	"fmt"

	"github.com/google/uuid"
)

func GetUuid() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return u.String(), nil
}
