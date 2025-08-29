package mutex

import (
	"fmt"
	"sync"
)

func Increment() {
	var (
		counter int
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)

	// Jalankan 5 goroutine yang masing-masing menambahkan counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			mutex.Lock()
			counter++
			fmt.Printf("Goroutine %d menambahkan counter menjadi %d\n", id, counter)
			mutex.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
