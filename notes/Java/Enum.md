# Enum

Тип enum — это тип, поля которого состоят из набора некоторых констант. Простой пример — стороны света, такой объект будет содержать 4 константы: NORTH, SOUTH, EAST, и WEST. Переменные объявлены большими буквами, потому что это константы.

## Синтаксис
Для объявления используется ключевое слово enum. Например, нам необходимо перечислить дни недели:

```
public enum Day {
    SUNDAY, MONDAY, TUESDAY, WEDNESDAY,
    THURSDAY, FRIDAY, SATURDAY
}
```
Enum тип необходимо использовать, если нужно определить некоторое количество констант, значения которых известны заранее, например, варианты пунктов меню или планеты нашей солнечной системы.