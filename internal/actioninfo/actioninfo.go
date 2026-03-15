package actioninfo

import (
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	// Перебрать все значения слайса dataset в цикле.
	// Распарсить каждое значение с помощью метода Parse().
	if len(dataset) == 0 {
		log.Println("Нет данных для обработки")
		return
	}
	for _, data := range dataset {
		// Обработать ошибку парсинга. Если она возникает, нужно залогировать ошибку и перейти к следующей итерации цикла.
		err := dp.Parse(data)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	// Сформировать и вывести строку с информацией об активности с помощью метода ActionInfo(). При возникновении ошибки ее нужно залогировать.
	info, err := dp.ActionInfo()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(info)
}
