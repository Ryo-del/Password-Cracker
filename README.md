# 🔐 Password Cracker (Go)

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://golang.org/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Multi-Threaded](https://img.shields.io/badge/Threads-8-important)](https://golang.org/doc/effective_go#goroutines)

Многопоточный инструмент для подбора паролей на Go. Проверяет пароли сначала в базе популярных вариантов, затем перебирает комбинации указанной длины с использованием горутин для максимальной производительности.

## 🚀 Особенности

- ⚡ **Мгновенная проверка** против 100+ популярных паролей (`database.json`)
- 🧩 **Поддержка кастомных символов** (цифры + латиница)
- 🛡 **Оптимизированный перебор** с распределением по потокам
- 📊 **Измерение времени** работы
- 🔊 **Звуковое оповещение** при успешном подборе

## 📦 Установка

### Вариант 1: Готовый бинарник (Windows/Linux)
```bash
# Скачать последний релиз
curl -L https://github.com/ваш-username/password-cracker/releases/latest/download/password-cracker -o password-cracker
chmod +x password-cracker
```
Вариант 2: Сборка из исходников
```bash

git clone https://github.com/ваш-username/password-cracker.git
cd password-cracker
go build -o password-cracker .
``` 
## 🖥 Использование
```bash

./password-cracker
```
Программа запросит ввод пароля для анализа. Пример работы:

Введите пароль: qwerty
[✓] Пароль найден в базе популярных (0.2ms)

## ⚙️ Настройка

Измените параметры в коде при необходимости:
```go

var charset = []rune("0123456789abcdefghijklmnopqrstuvwxyz") // Доступные символы
const numWorkers = 8                                         // Количество потоков
```

## 📊 Производительность
Длина пароля	Время (8 потоков)
3 символа	< 1 сек
4 символа	~5 сек
5 символов	~2 мин
## 🌍 Поддержка ОС

   Windows (x64)

   Linux (amd64/arm64)

   macOS (теоретически)

## ⚠️ Ответственность

   Внимание! Используйте только для:

      Тестирования собственных систем

      Восстановления доступа к своим данным

      Исследовательских целей

Запрещено применять для несанкционированного доступа!
## 📄 Лицензия

[License](https://github.com/Ryo-del/Password-Cracker/blob/main/LICENSE)

