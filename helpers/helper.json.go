package helpers

import (
	"encoding/json"

	"github.com/sirupsen/logrus"

	schemas "github.com/nutwreck/admin-pos-service/schemes"
)

func Strigify(payload interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

func Parse(payload []byte) schemas.SchemeResponses {
	var jsonResponse schemas.SchemeResponses
	err := json.Unmarshal(payload, &jsonResponse)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return jsonResponse
}
