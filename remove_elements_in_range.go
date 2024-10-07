package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
    length := len(arr)

    // Обрабатываем отрицательные индексы
    if from < 0 {
        from += length
    }
    if to < 0 {
        to += length
    }

    // Если индексы выходят за пределы массива, корректируем их
    if from < 0 {
        from = 0
    }
    if from >= length {
        from = length - 1
    }
    if to < 0 {
        to = 0
    }
    if to >= length {
        to = length - 1
    }

    // Если индексы в неправильном порядке, меняем их местами
    if from > to {
        from, to = to, from
    }

    // Удаляем элементы от from до to-1 включительно, оставляем to
    result := append(arr[:from], arr[to:]...)

    return result
}