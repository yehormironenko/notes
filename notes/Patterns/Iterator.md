# Iterator

[original link ](https://refactoring.guru/ru/design-patterns/iterator/java/example)

***Итератор*** — *это поведенческий паттерн, позволяющий последовательно обходить сложную коллекцию, без раскрытия деталей её реализации.*

**Применимость**: Паттерн можно часто встретить в Java-коде, особенно в программах, работающих с разными типами коллекций, и где требуется обход разных сущностей.

**Примеры Итератора в стандартных библиотеках Java**:

Все реализации java.util.Iterator (среди прочего также java.util.Scanner).

Все реализации java.util.Enumeration

**Признаки применения паттерна:** Итератор легко определить по методам навигации (например, получения следующего/предыдущего элемента и т.д.). Код использующий итератор зачастую вообще не имеет ссылок на коллекцию, с которой работает итератор. Итератор либо принимает коллекцию в параметрах конструкторе при создании, либо возвращается самой коллекцией.