package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type Scheduler struct {
	scheduler gocron.Scheduler
}

func New() (*Scheduler, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	s.Start()

	return &Scheduler{scheduler: s}, nil
}
func (s *Scheduler) Stop() error {
	return s.scheduler.Shutdown()
}

func (s *Scheduler) NewJob(cron string, fn any, params ...any) error {
	_, err := s.scheduler.NewJob(
		gocron.CronJob(cron, true),
		gocron.NewTask(fn, params...),
	)
	return err
}

// запуск каждую секунду
func (s *Scheduler) EverySecond(fn any, params ...any) error {
	return s.NewJob("* * * * * *", fn, params...)
}

// Каждый n секунд
func (s *Scheduler) EverySeconds(n uint, fn any, params ...any) error {
	if n < 1 {
		return fmt.Errorf("invalid second interval: %d", n)
	}
	return s.NewJob(fmt.Sprintf("*/%d * * * * *", n), fn, params...)
}

// запуск каждую минуту
func (s *Scheduler) EveryMinute(fn any, params ...any) error {
	return s.NewJob("0 * * * * *", fn, params...)
}
func (s *Scheduler) EveryMinutes(n uint, fn any, params ...any) error {
	if n < 1 {
		return fmt.Errorf("invalid minute interval: %d", n)
	}
	spec := fmt.Sprintf("0 */%d * * * *", n)
	return s.NewJob(spec, fn, params...)
}

// запуск каждый час
func (s *Scheduler) EveryHour(fn any, params ...any) error {
	return s.NewJob("0 0 * * * *", fn, params...)
}
func (s *Scheduler) EveryHours(n uint, fn any, params ...any) error {
	if n <= 0 {
		return fmt.Errorf("invalid hour interval: %d", n)
	}
	spec := fmt.Sprintf("0 0 */%d * * *", n) // каждые n часов на нулевой минуте
	return s.NewJob(spec, fn, params...)
}

// запуск каждый день
func (s *Scheduler) EveryDay(fn any, params ...any) error {
	return s.NewJob("0 0 0 * * *", fn, params...)
}

// запуск каждую неделю
func (s *Scheduler) EveryWeek(fn any, params ...any) error {
	return s.NewJob("0 0 0 * * 0", fn, params...)
}

// запуск каждый месяц
func (s *Scheduler) EveryMonth(fn any, params ...any) error {
	return s.NewJob("0 0 0 1 * *", fn, params...)
}

// запуск каждый год
func (s *Scheduler) EveryYear(fn any, params ...any) error {
	return s.NewJob("0 0 0 1 1 *", fn, params...)
}

// запуск каждый день в определенное время
func (s *Scheduler) DailyAt(hour, minute int, fn any, params ...any) error {
	return s.NewJob(fmt.Sprintf("%d %d * * *", minute, hour), fn, params...)
}

// запуск каждую неделю в определенное время
func (s *Scheduler) WeeklyAt(day int, hour, minute int, fn any, params ...any) error {
	return s.NewJob(fmt.Sprintf("%d %d * * %d", minute, hour, day), fn, params...)
}

// запуск каждый месяц в определенное время
func (s *Scheduler) MonthlyAt(day, hour, minute int, fn any, params ...any) error {
	return s.NewJob(fmt.Sprintf("%d %d %d * *", minute, hour, day), fn, params...)
}

// запуск каждый год в определенное время
func (s *Scheduler) YearlyAt(month, day, hour, minute int, fn any, params ...any) error {
	return s.NewJob(fmt.Sprintf("%d %d %d %d *", minute, hour, day, month), fn, params...)
}

// запуск через определенное время (повторяющийся)
func (s *Scheduler) Duration(duration time.Duration, fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.DurationJob(
			duration,
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через определенное время
func (s *Scheduler) After(duration time.Duration, fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(duration)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через секунду
func (s *Scheduler) AfterSecound(fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(1*time.Second)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через n секунд
func (s *Scheduler) AfterSecounds(n uint, fn any, params ...any) (gocron.Job, error) {
	if n < 1 {
		return nil, fmt.Errorf("invalid second interval: %d", n)
	}
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(time.Duration(n)*time.Second)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через минуту
func (s *Scheduler) AfterMinute(fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(1*time.Minute)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через n минут
func (s *Scheduler) AfterMinutes(n uint, fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(time.Duration(n)*time.Minute)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через час
func (s *Scheduler) AfterHour(fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(1*time.Hour)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через n часов
func (s *Scheduler) AfterHours(n uint, fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(time.Duration(n)*time.Hour)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через день
func (s *Scheduler) AfterDay(fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(24*time.Hour)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через n дней
func (s *Scheduler) AfterDays(n uint, fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(time.Duration(n)*24*time.Hour)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через неделю
func (s *Scheduler) AfterWeek(fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(7*24*time.Hour)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через месяц (30 дней)
func (s *Scheduler) AfterMonth(fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(30*24*time.Hour)),
		),
		gocron.NewTask(fn, params...),
	)
}

// запуск один раз через год (365 дней)
func (s *Scheduler) AfterYear(fn any, params ...any) (gocron.Job, error) {
	return s.scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(365*24*time.Hour)),
		),
		gocron.NewTask(fn, params...),
	)
}
