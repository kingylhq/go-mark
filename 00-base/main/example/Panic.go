package example

import "os"

func Panic() {

	panic("a problem")

	_, err := os.Create("/Users/ouyangdadi/GolandProjects/go-mark/01-gin/")
	if err != nil {
		panic(err)
	}
}
