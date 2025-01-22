package example

import (
	"bufio"
	"fmt"
	"os"
)

func check2(e error) {
	if e != nil {
		panic(e)
	}
}

func WritingFiles() {

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/Users/ouyangdadi/GolandProjects/go-mark/01-gin/defer.txt", d1, 0644)
	check2(err)

	f, err := os.Create("/Users/ouyangdadi/GolandProjects/go-mark/01-gin/defer.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
