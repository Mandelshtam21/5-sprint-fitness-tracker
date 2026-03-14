package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int                   // количество шагов, проделанных за тренировку.
	TrainingType string                // тип тренировки(бег или ходьба).
	Duration     time.Duration         // длительность тренировки.
	Personal     personaldata.Personal // встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	// Разделить строку datastring на слайс строк.
	parts := strings.Split(datastring, ",")
	// Проверить, чтобы длина слайса была равна 3, так как в строке данных у нас количество шагов,
	// вид активности и продолжительность.
	if len(parts) != 3 {
		return fmt.Errorf("Неверный формат данных")
	}
	// Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	// При возникновении ошибки вернуть её из метода.
	// Сохранить полученное значение в соответствующем поле структуры Training.
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("Количество шагов должно быть больше 0")
	}
	t.Steps = steps
	// Сохранить значение типа тренировки в соответствующем поле структуры Training.
	t.TrainingType = parts[1]
	// Преобразовать третий элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration.
	// Обработать возможные ошибки. При их возникновении вернуть ошибку,
	// в противном случае сохранить полученное значение в соответствующем поле структуры Training.
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("Продолжительность должна быть больше 0")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	// - возвращать ошибку и строку
	// - описать сообщения об ошибках
	// Вычислить дистанцию, используя функцию из пакета spentenergy.
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	// Вычислить среднюю скорость, используя функцию из пакета spentenergy.
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	// Проверить, какой вид тренировки содержится в структуре Training.
	// Для каждого из видов тренировок рассчитать калории, используя функцию из пакета spentenergy.
	var calories float64
	var err error
	if t.TrainingType == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	} else if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("Неизвестный тип тренировки")
	}
	// Сформируйте и верните строку, образец которой был выше.
	// Если был передан неизвестный тип тренировки, верните ошибку с текстом неизвестный тип тренировки.
	/* Тип тренировки: Бег
	Длительность: 0.75 ч.
	Дистанция: 10.00 км.
	Скорость: 13.34 км/ч
	Сожгли калорий: 18621.75 */
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f\nСкорость: %.2f\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration, distance, speed, calories), nil
}
