# hw3

## Message Formatter

Реализуйте интерфейс Formatter с методом Format, который возвращает
отформатированную строку.

Определите структуры, удовлетворяющие интерфейсу Formatter: обычный
текст(как есть), жирным шрифтом(** **), код(` `), курсив(_ _)

Опционально: иметь возможность задавать цепочку модификаторов

chainFormatter.AddFormatter(plainText)
chainFormatter.AddFormatter(bold)
chainFormatter.AddFormatter(code)