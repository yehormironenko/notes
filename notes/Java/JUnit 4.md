# JUnit 4

**JUnit** — библиотека для модульного тестирования программного обеспечения на языке Java. 

Чтобы runner тестов мог определить, кто есть кто, тестовые методы необходимо помечать аннтоацией @Test.

У аннотации могут быть проставлены такие параметры:

+ expected - указываем какое исключение будет сгенерировано методом (см. пример ниже);
+ timeout - через какое время в милисекундах прекратить выполнение теста и засчитать его как неуспешный.

Если вы хотите указать, что определенный тест необхдимо пропустить, то пометьте его аннотацией @Ignore. Хотя можно просто удалить аннотацию @Test.

@Before - выполняется до запуска теста
@After  - после запуска
@BeforeClass & @AfterClass - выполняется один раз до или после

### Методы

+ `assertTrue` - проверяет, является ли результат выражения верным
+ `assertEquals` - ожидаемый результат и полученный результат совпадают;
+ `assertNull` - результатом выражения есть null;
+ `assertNotNull` - результат выражения отличен от null;
+ `assertSame` - ожидаемый и полученный объекты это один и тот же объект.
+ `fail` - метод генерирует исключение 
+ `AssertionError` - добавляем туда, куда не должен дойти ход выполнения программы.
