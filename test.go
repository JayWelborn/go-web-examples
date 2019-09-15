package main
import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

    var ops uint64
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)

        go func() {
            for ops < 1000000000000 {
                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("ops:", ops)
}