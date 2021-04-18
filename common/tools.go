package common

import (
	"encoding/json"
	"log"
)

func FailOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Marshal 序列化消息
func Marshal(message interface{}) (buff []byte, err error) {
	switch message.(type) {
	case []byte:
		return message.([]byte), nil
	case string:
		buff = []byte(message.(string))
	default:
		buff, err = json.Marshal(message)
		return
	}
	return
}