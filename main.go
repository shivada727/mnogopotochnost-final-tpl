package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}
	return max
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize
		if i == CHUNKS-1 {
			endIndex = len(data)
		}
		chunk := data[startIndex:endIndex]
		go func(chunkIndex int, chunkData []int) {
			defer wg.Done()

			if len(chunkData) == 0 {
				maxValues[chunkIndex] = 0
				return
			}

			maxValues[chunkIndex] = maximum(chunkData)
		}(i, chunk)
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
