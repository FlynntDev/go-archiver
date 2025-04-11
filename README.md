# go-archiver

`go-archiver` — это простой инструмент для сжатия и распаковки файлов, реализованный на языке Go. Он поддерживает методы сжатия, такие как VLC (Variable Length Coding), с использованием алгоритма Шеннона-Фано.

## Возможности

- **Сжатие файлов**: Преобразует текстовые файлы в сжатый формат.
- **Распаковка файлов**: Восстанавливает оригинальный текст из сжатого формата.
- **Поддержка алгоритма Шеннона-Фано**: Используется для генерации таблиц кодирования.

## Установка

1. Убедитесь, что у вас установлен Go версии 1.22 или выше.
2. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/FlynntDev/go-archiver.git
   ```
3. Перейдите в директорию проекта:
   ```bash
   cd go-archiver
   ```
4. Установите зависимости:
   ```bash
   go mod tidy
   ```

## Использование

### Сжатие файла

Для сжатия файла используйте команду `pack`:

```bash
go run main.go pack -m vlc <путь_к_файлу>
```

Пример:
```bash
go run main.go pack -m vlc example.txt
```

Результат будет сохранен в файл с расширением `.vlc`.

### Распаковка файла

Для распаковки файла используйте команду `unpack`:

```bash
go run main.go unpack -m vlc <путь_к_файлу>
```

Пример:
```bash
go run main.go unpack -m vlc example.vlc
```

Результат будет сохранен в файл с расширением `.txt`.

## Структура проекта

```
├── cmd/
│   ├── pack.go          # Реализация команды сжатия
│   ├── unpack.go        # Реализация команды распаковки
│   └── root.go          # Основная точка входа для команд
├── lib/
│   ├── compression/
│   │   ├── compression.go  # Интерфейсы Encoder и Decoder
│   │   └── vlc/
│   │       ├── vlc.go       # Реализация VLC кодирования
│   │       ├── chunks.go    # Работа с бинарными блоками
│   │       └── table/
│   │           ├── table.go # Таблицы кодирования и декодирования
│   │           └── shannon_fano/
│   │               ├── shannon_fano.go       # Генерация таблиц Шеннона-Фано
│   │               └── shannon_fano_test.go # Тесты для Шеннона-Фано
├── main.go              # Точка входа в приложение
├── go.mod               # Модуль Go
├── README.md            # Документация
```

## Тестирование

Для запуска тестов выполните:

```bash
go test ./...
```

## Требования

- Go 1.22 или выше