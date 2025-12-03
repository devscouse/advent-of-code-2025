/*Package common contains small pieces of code that are used across challenges.*/
package common

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
