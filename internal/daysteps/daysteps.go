package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps                 int           // количество шагов.
	Duration              time.Duration // длительность прогулки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print()./*
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("Неверный формат данных")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("Количество шагов должно быть больше 0")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("Продолжительность должна быть больше 0")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	// Вычислите дистанцию.
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	// Вычислите количество сожжённых калорий. При возникновении ошибки верните пустую строку и ошибку.
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	// Сформируйте и верните строку с информацией.
	/*Количество шагов: 792.
	Дистанция составила 0.51 км.
	Вы сожгли 221.33 ккал.*/
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
