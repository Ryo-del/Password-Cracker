package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var charset = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

const numWorkers = 8 // число потоков

func main() {
	var pass string
	fmt.Print("Введите пароль: ")
	fmt.Scan(&pass)

	start := time.Now()

	length := len(pass)
	total := pow(len(charset), length)

	var found atomic.Value
	var wg sync.WaitGroup
	var stopFlag int32 = 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Считываем файл с паролями
		file, err := os.Open("database.json")
		if err != nil {
			fmt.Println("Ошибка открытия файла c популярными паролями:", err)
			return
		}
		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Ошибка чтения файла c популярными паролями:", err)
			return
		}
		var poppass []string
		err = json.Unmarshal(bytes, &poppass)
		if err != nil {
			fmt.Println("Ошибка парсинга JSON:", err)
			return
		}

		for _, p := range poppass {
			if p == pass {
				found.Store(pass)
				atomic.StoreInt32(&stopFlag, 1)
				fmt.Println("Пароль найден в файле c популярными паролями:")
				fmt.Print("\a")
				return
			}
		}
	}()

	// Запускаем numWorkers горутин
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID, total int) {
			defer wg.Done()
			for i := workerID; i < total; i += numWorkers {
				if atomic.LoadInt32(&stopFlag) == 1 {
					return
				}
				guess := indexToString(i, length)
				if guess == pass {
					found.Store(guess)
					atomic.StoreInt32(&stopFlag, 1)
					return
				}
			}
		}(w, total)
	}

	wg.Wait()

	// Вывод результатов
	if value := found.Load(); value != nil {
		fmt.Println("Пароль найден:", value.(string))
		fmt.Print("\a")
		fmt.Println("Время подбора:", time.Since(start))
	} else {
		fmt.Println("Пароль не найден")
	}
}

// Преобразуем индекс в строку по charset
func indexToString(index int, length int) string {
	base := len(charset)
	result := make([]rune, length)
	for i := length - 1; i >= 0; i-- {
		result[i] = charset[index%base]
		index /= base
	}
	return string(result)
}

// Быстрое возведение в степень
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
