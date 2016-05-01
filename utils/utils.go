// Author: Fernando Segovia Salgado
// Creation: 04/30/2016
// Last Update: 05/01/2016
// Description: Utilitary functions
// file: utils.go
// package: utils
// dependencies:
//      - "encoding/json"
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

// Converts a interface to JSON
func ConvertToJSON(mapToConvert interface{}) ([]byte) {
    result, err := json.MarshalIndent(mapToConvert, "", " ")
    CheckError(err)
    return result
}