func GenerateRange(min, max int) []int {
    // Если min >= max, возвращаем nil
    if min >= max {
        return nil
    }

    // Создаём срез нужного размера
    size := max - min
    result := make([]int, size)

    // Заполняем срез числами от min до max-1
    for i := 0; i < size; i++ {
        result[i] = min + i
    }

    // Возвращаем результат после завершения цикла
    return result
}