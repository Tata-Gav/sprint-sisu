#!/bin/bash

# Создаем файлы a, !, \ и "
touch a
touch '!'
touch '\\'
touch '"'

# Создаем директорию, если она не существует
if [ ! -d '`' ]; then
    mkdir '`'
fi

# Копируем файл ! в директорию `
cp '!' '`'

# Если переменная окружения MOVE_A == "yes", перемещаем файл a в директорию `
if [ "$MOVE_A" = "yes" ]; then
    mv a '`'
elif [ "$MOVE_A" = "no" ]; then
    rm a   # Если переменная MOVE_A == "no", удаляем файл a
fi
