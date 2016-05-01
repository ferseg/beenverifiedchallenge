package utils

import(
	"encoding/json"
)


// Checks if an error occurs
func CheckError(err error) {
	if err != nil {
        panic(err)
    }
}

func ConvertToJSON(mapToConvert interface{}) ([]byte) {
	result, err := json.MarshalIndent(mapToConvert, "", " ")
	CheckError(err)
	return result
}