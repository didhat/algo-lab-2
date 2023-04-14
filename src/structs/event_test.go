package structs

import (
	"testing"
)

func insertXIntMap(x int, b *testing.B) {
	// Инициализируем Map и вставляем X элементов
	testmap := make(map[int]int, 0)
	// Сбрасываем таймер после инициализации Map
	b.ResetTimer()
	for i := 0; i < x; i++ {
		// Вставляем значение I в ключ I.
		testmap[i] = i
	}
}

func BenchmarkInsertIntMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntMap(100000, b)
	}
}

// BenchmarkInsertIntMap10000 тестирует скорость вставки 10000 целых чисел в карту.
func BenchmarkInsertIntMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntMap(10000, b)
	}
}

// BenchmarkInsertIntMap1000 тестирует скорость вставки 1000 целых чисел в карту.
func BenchmarkInsertIntMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntMap(1000, b)
	}
}

// BenchmarkInsertIntMap100 тестирует скорость вставки 100 целых чисел в карту.
func BenchmarkInsertIntMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntMap(100, b)
	}
}
