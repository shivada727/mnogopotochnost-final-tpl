package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"Нулевой размер", 0, 0},
		{"Размер 1", 1, 1},
		{"Размер 10", 10, 10},
		{"Размер 100", 100, 100},
		{"Отрицательный размер", -5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			if len(result) != tt.expected {
				t.Errorf("generateRandomElements(%d) = %d элементов, ожидалось %d",
					tt.size, len(result), tt.expected)
			}

			if len(result) > 0 {
				for i, num := range result {
					if num <= 0 {
						t.Errorf("Элемент %d равен %d, должен быть положительным", i, num)
					}
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"Пустой слайс", []int{}, 0},
		{"Один элемент", []int{42}, 42},
		{"Два элемента", []int{10, 20}, 20},
		{"Несколько элементов", []int{1, 5, 3, 9, 2}, 9},
		{"Отрицательные числа", []int{-5, -10, -1, -100}, -1},
		{"Все одинаковые", []int{7, 7, 7, 7}, 7},
		{"Максимум в начале", []int{100, 1, 2, 3}, 100},
		{"Максимум в конце", []int{1, 2, 3, 100}, 100},
		{"Максимум в середине", []int{1, 100, 2, 3}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, result, tt.expected)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"Пустой слайс", []int{}, 0},
		{"Один элемент", []int{42}, 42},
		{"Меньше элементов чем чанков", []int{1, 2, 3}, 3},
		{"Точно 8 элементов", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"Больше элементов чем чанков", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 16},
		{"Максимум в первом чанке", []int{100, 1, 1, 1, 1, 1, 1, 1}, 100},
		{"Максимум в последнем чанке", []int{1, 1, 1, 1, 1, 1, 1, 100}, 100},
		{"Максимум в среднем чанке", []int{1, 1, 1, 100, 1, 1, 1, 1}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.data)
			if result != tt.expected {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, result, tt.expected)
			}
		})
	}
}

func TestLargeDataConsistency(t *testing.T) {
	data := generateRandomElements(1000)

	maxSingle := maximum(data)
	maxMulti := maxChunks(data)

	if maxSingle != maxMulti {
		t.Errorf("Результаты не совпадают: single=%d, multi=%d", maxSingle, maxMulti)
	}
}

func BenchmarkMaximum(b *testing.B) {
	data := generateRandomElements(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maximum(data)
	}
}

func BenchmarkMaxChunks(b *testing.B) {
	data := generateRandomElements(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maxChunks(data)
	}
}
