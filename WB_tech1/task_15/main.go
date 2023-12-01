package main

var justString string

func createHugeString(sz int) string {
	return string(make([]byte, sz))
}

func someFunc() {
	v := createHugeString(1 << 10) // Создаём строку размера 1024
	t := make([]byte, 100)         // Создаём отдельный буфер
	copy(t, v[:100])               // Копируем в буфер подстроку до 100-го символа, при этом мы не тащим за собой большую строку
	justString = string(t)
}

func main() {
	someFunc()
}
