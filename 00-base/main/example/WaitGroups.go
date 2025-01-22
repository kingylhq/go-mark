package example

import (
	"fmt"
	"sync"
	"time"
)

func worker3(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroup() {

	var wg sync.WaitGroup

	//wg.Add(5)
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker3(i)
		}()
	}

	wg.Wait()

}
