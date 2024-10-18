package mutex

// Mutex - собственная реализация мьютекса на основе канала
type Mutex struct {
    ch chan struct{}
}

// NewMutex создает новый мьютекс
func NewMutex() *Mutex {
    return &Mutex{
        ch: make(chan struct{}, 1), // Канал буферизован на 1 элемент
    }
}

// Lock блокирует мьютекс. Если канал пуст, горутина ждет
func (m *Mutex) Lock() {
    m.ch <- struct{}{} // Отправляем в канал, блокируя его
}

// Unlock разблокирует мьютекс, извлекая элемент из канала
func (m *Mutex) Unlock() {
    <-m.ch // Извлекаем элемент, разблокируя канал
}
