package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
    length := len(arr)

    // Обработка отрицательных индексов
    if from < 0 {
        from += length
    }
    if to < 0 {
        to += length
    }

    // Убедимся, что индексы не выходят за пределы массива
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

    // Если индексы в неправильном порядке, поменяем их
    if from > to {
        from, to = to, from
    }

    // Удаляем элементы от from до to-1 включительно
    result := append(arr[:from], arr[to:]...)

    return result
}