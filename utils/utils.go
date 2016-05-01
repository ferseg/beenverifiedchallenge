package utils



// Checks if an error occurs
func CheckError(err error) {
	if err != nil {
        panic(err)
    }
}