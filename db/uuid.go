package db

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GetUUID() (string, error) {
	// generate 32 bits timestamp
	unix32bits := uint32(time.Now().UTC().Unix())

	buff := make([]byte, 8)

	numRead, err := rand.Read(buff)
	if numRead != len(buff) || err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x%x%x%x%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8])
	return uuid, nil
}
