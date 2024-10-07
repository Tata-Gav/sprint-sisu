package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	// Получаем длину массива
	length := len(arr)

	// Обрабатываем отрицательные индексы
	if from < 0 {
		from += length
	}
	if to < 0 {
		to += length
	}

	// Корректируем индексы, если они выходят за пределы
	if from >= length {
		from = length - 1
	}
	if to >= length {
		to = length - 1
	}
	if to < 0 {
		to = 0
	}

	// Если индексы в неправильном порядке, меняем их местами
	if from > to {
		from, to = to, from
	}

	// Создаём новый массив, который включает элементы до from и после to
	result := append(arr[:from], arr[to:]...)

	return result
}