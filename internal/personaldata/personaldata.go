package personaldata

import (
	"fmt"
)

type Personal struct {
	// TODO: добавить поля
	// Name — имя пользователя;
	// Weight— вес пользователя;
	// Height — рост пользователя.
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	// Метод ничего не принимает и ничего не возвращает. Он просто выводит данные структуры на экран
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
