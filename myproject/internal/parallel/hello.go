package parallel

import (
    "fmt"
    "myproject/pkg/mutex"
    "time"
)

// Say демонстрирует использование мьютекса
func Say() {
    m := mutex.NewMutex() // Создаем экземпляр нашего мьютекса

    m.Lock()
    fmt.Println("Main goroutine: locked")

    go func() {
        m.Lock()
        fmt.Println("Goroutine: locked")
        time.Sleep(2 * time.Second) // Имитируем работу
        m.Unlock()
        fmt.Println("Goroutine: unlocked")
    }()

    time.Sleep(1 * time.Second)
    fmt.Println("Main goroutine: trying to lock")
    m.Lock() // Ожидаем разблокировки от горутины
    fmt.Println("Main goroutine: unlocked")
    m.Unlock()
}
