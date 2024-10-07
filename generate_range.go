func GenerateRange(min, max int) []int {
    // Проверяем условие: если min >= max, возвращаем nil
    if min >= max {
        return nil
    }

    // Определяем размер среза как разницу между max и min
    size := max - min

    // Создаём срез необходимого размера
    result := make([]int, size)

    // Заполняем срез значениями от min до max-1
    for i := 0; i < size; i++ {
        result[i] = min + i
    }

    // Возвращаем срез после того, как все значения будут добавлены
    return result
}
