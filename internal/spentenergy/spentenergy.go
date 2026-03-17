package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	if steps <= 0 { // || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("steps must be more than 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be more than 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be more than 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be more than 0")
	}
	// Рассчитать среднюю скорость с помощью meanSpeed().
	speed := MeanSpeed(steps, height, duration)
	// Рассчитать количество калорий. Для этого:
	// Переведите продолжительность в минуты с помощью функции из пакета time.
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий.
	// Умножить полученное число калорий на корректирующий коэффициент walkingCaloriesCoefficient.
	// Соответствующая константа объявлена в пакете. Вернуть полученное значение.
	return ((duration.Minutes() * weight * speed) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	if steps <= 0 { // || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("steps must be more than 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be more than 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be more than 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be more than 0")
	}
	// Рассчитать среднюю скорость с помощью meanSpeed().
	speed := MeanSpeed(steps, height, duration)
	// Рассчитать и вернуть количество калорий. Для этого:
	// Переведите продолжительность в минуты с помощью функции из пакета time.
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий.
	// (weight * meanSpeed * durationInMinutes) / minInH
	return (duration.Minutes() * weight * speed) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	// Проверить, что продолжительность duration больше 0. Если это не так, вернуть 0.
	if duration <= 0 || steps <= 0 {
		return 0
	}
	// Вычислить дистанцию с помощью Distance().
	dist := Distance(steps, height)
	// Вычислить и вернуть среднюю скорость. Для этого разделите дистанцию на продолжительность в часах.
	// Чтобы перевести продолжительность в часы, воспользуйтесь функцией из пакета time.
	return dist / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	// рассчитайте длину шага. Для этого умножьте высоту пользователя на коэффициент длины шага stepLengthCoefficient.
	// Соответствующая константа уже определена в пакете.
	stepLength := height * stepLengthCoefficient
	// умножьте пройденное количество шагов на длину шага.
	// разделите полученное значение на число метров в километре (mInKm, константа определена в пакете)
	return (float64(steps) * stepLength) / mInKm
}
