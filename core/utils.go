/*Package common contains small pieces of code that are used across challenges.*/
package core

import (
	"os"
	"path/filepath"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadPackageData(packageName string, fileName string) *os.File {
	path := filepath.Join(".", packageName, "data", fileName)
	file, err := os.Open(path)
	Check(err)
	return file
}
