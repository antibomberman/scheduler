# Go Scheduler

Простая и удобная обертка над библиотекой [gocron](https://github.com/go-co-op/gocron) для планирования задач в Go.

## Установка

```bash
go get github.com/go-co-op/gocron/v2
```

## Быстрый старт

```go
package main

import (
	"fmt"
	"time"
	"your-project/scheduler"
)

func main() {
	// Создаем новый планировщик
	s, err := scheduler.New()
	if err != nil {
		panic(err)
	}
	defer s.Stop()

	// Планируем задачу
	s.EverySecond(sayHello, "World")

	// Ждем выполнения
	time.Sleep(5 * time.Second)
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
```

## API

### Создание планировщика

```go
s, err := scheduler.New()
if err != nil {
log.Fatal(err)
}
defer s.Stop() // Остановка планировщика при завершении
```

### Повторяющиеся задачи

#### Временные интервалы
```go
// Каждую секунду
s.EverySecond(myFunc)

// Каждые N секунд
s.EverySeconds(5, myFunc)

// Каждую минуту
s.EveryMinute(myFunc)

// Каждые N минут
s.EveryMinutes(10, myFunc)

// Каждый час
s.EveryHour(myFunc)

// Каждые N часов
s.EveryHours(6, myFunc)

// Каждый день
s.EveryDay(myFunc)

// Каждую неделю
s.EveryWeek(myFunc)

// Каждый месяц
s.EveryMonth(myFunc)

// Каждый год
s.EveryYear(myFunc)
```

#### Запуск в определенное время
```go
// Каждый день в 14:30
s.DailyAt(14, 30, myFunc)

// Каждую неделю в воскресенье (0) в 09:00
s.WeeklyAt(0, 9, 0, myFunc)

// Каждый месяц 15 числа в 12:00
s.MonthlyAt(15, 12, 0, myFunc)

// Каждый год 1 января в 00:00
s.YearlyAt(1, 1, 0, 0, myFunc)
```

#### Повторяющиеся задачи с интервалом
```go
// Каждые 5 минут
job, err := s.Duration(5*time.Minute, myFunc)
```

### Одноразовые задачи (выполняются только один раз)

```go
// Через 10 секунд
job, err := s.After(10*time.Second, myFunc)

// Через 1 секунду
job, err := s.AfterSecound(myFunc)

// Через N секунд
job, err := s.AfterSecounds(30, myFunc)

// Через 1 минуту
job, err := s.AfterMinute(myFunc)

// Через N минут
job, err := s.AfterMinutes(15, myFunc)

// Через 1 час
job, err := s.AfterHour(myFunc)

// Через N часов
job, err := s.AfterHours(2, myFunc)

// Через день
job, err := s.AfterDay(myFunc)

// Через N дней
job, err := s.AfterDays(7, myFunc)

// Через неделю
job, err := s.AfterWeek(myFunc)

// Через месяц
job, err := s.AfterMonth(myFunc)

// Через год
job, err := s.AfterYear(myFunc)
```

### Кастомные cron задачи

```go
// Используя cron выражение
err := s.NewJob("0 */2 * * *", myFunc) // Каждые 2 часа
```

## Передача параметров

Все методы поддерживают передачу параметров в функцию:

```go
func processData(id int, name string, config map[string]interface{}) {
fmt.Printf("Processing: %d, %s, %v\n", id, name, config)
}

config := map[string]interface{}{"timeout": 30}
s.EveryMinute(processData, 123, "test", config)
```

## Примеры использования

### Простой пример
```go
func main() {
s, _ := scheduler.New()
defer s.Stop()

// Логирование каждые 10 секунд
s.EverySeconds(10, func() {
fmt.Println("Heartbeat:", time.Now())
})

// Уведомление один раз через 5 минут
s.AfterMinutes(5, func() {
fmt.Println("Напоминание выполнено!")
})

// Ждем выполнения
select {}
}
```

### Сложный пример с параметрами
```go
type Config struct {
DatabaseURL string
Timeout     time.Duration
}

func backupDatabase(cfg Config) {
fmt.Printf("Backing up database: %s\n", cfg.DatabaseURL)
// Логика резервного копирования
}

func sendReport(email string, reportType string) {
fmt.Printf("Sending %s report to %s\n", reportType, email)
// Логика отправки отчета
}

func main() {
s, _ := scheduler.New()
defer s.Stop()

cfg := Config{
DatabaseURL: "postgresql://localhost/mydb",
Timeout:     30 * time.Second,
}

// Резервное копирование каждый день в 2:00
s.DailyAt(2, 0, backupDatabase, cfg)

// Еженедельный отчет в понедельник в 9:00
s.WeeklyAt(1, 9, 0, sendReport, "admin@company.com", "weekly")

// Уведомление через час
s.AfterHour(func() {
fmt.Println("Система работает уже час!")
})

select {}
}
```

## Форматы времени в cron

При использовании `NewJob()` поддерживается формат cron с секундами:

```
* * * * * *
│ │ │ │ │ │
│ │ │ │ │ └─ день недели (0-6, 0=воскресенье)
│ │ │ │ └─── месяц (1-12)
│ │ │ └───── день месяца (1-31)
│ │ └─────── час (0-23)
│ └───────── минута (0-59)
└─────────── секунда (0-59)
```

Примеры:
- `"* * * * * *"` - каждую секунду
- `"0 */5 * * * *"` - каждые 5 минут
- `"0 0 9 * * 1-5"` - каждый рабочий день в 9:00

## Управление задачами

```go
// Создание задачи с возможностью управления
job, err := s.After(time.Hour, myFunc)
if err != nil {
log.Fatal(err)
}

// Получение информации о задаче
fmt.Println("Next run:", job.NextRun())
fmt.Println("Last run:", job.LastRun())
```

## Обработка ошибок

```go
s, err := scheduler.New()
if err != nil {
log.Fatal("Failed to create scheduler:", err)
}

err = s.EveryMinute(myFunc)
if err != nil {
log.Printf("Failed to schedule job: %v", err)
}
```

## Лицензия

MIT