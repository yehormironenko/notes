Static
========================

Модификатор *static* в Java напрямую связан с классом, если поле статично, значит оно принадлежит классу, если метод статичный, аналогично - он принадлежит классу

## Основные моменты 

1)   Вы НЕ можете получить доступ к НЕ статическим членам класса, внутри статического контекста, как вариант, метода или блока. Результатом компиляции приведенного ниже кода будет ошибка:
 ```
public class Counter{
private int count;
public static void main(String args[]){
   System.out.println(count); //compile time error
}}
```

2)  В отличие от локальных переменных, статические поля и методы НЕ потокобезопасны (Thread-safe) в Java. Каждый экземпляр класса имеет одну и ту же переменную. Поэтому при использовании статических переменных, убедитесь, что они должным образом синхронизированы (synchronized), во избежание проблем, например таких как «состояние гонки» (race condition). 

3)  Статические методы имеют преимущество в применении, т.к. отсутствует необходимость каждый раз создавать новый объект для доступа к таким методам. Статический метод можно вызвать, используя тип класса, в котором эти методы описаны. Именно поэтому, подобные методы как нельзя лучше подходят в качестве методов-фабрик (factory), и методов-утилит (utility). 

4)  Другим важным моментом является то, что вы НЕ можете переопределять (Override) статические методы. Если вы объявите такой же метод в классе-наследнике (subclass), т.е. метод с таким же именем  и сигнатурой, вы лишь «спрячете» метод суперкласса (superclass) вместо переопределения.  Это явление известно как сокрытие методов (hiding methods). Это означает, что при обращении к статическому методу, который объявлен как в родительском, так и в дочернем классе, во время компиляции всегда будет вызван метод  исходя из типа переменной.
5)  Объявить статическим также можно и класс, за исключением классов верхнего уровня. Такие классы известны как «вложенные статические классы» (nested static class). Они бывают полезными для представления улучшенных связей. Яркий пример вложенного статического класса - HashMap.Entry, который предоставляет структуру данных внутри HashMap.  
6)  Модификатор static также может быть объявлен в статичном блоке, более известным как «Статический блок инициализации» (Static initializer block), который будет выполнен во время загрузки класса. Если вы не объявите такой блок, то Java соберёт все статические поля в один список и выполнит его во время загрузки класса. Однако, статичный блок НЕ может пробросить перехваченные исключения, но может выбросить не перехваченные. В таком случае возникнет «Exception Initializer Error». 
7)  Полезно знать, что статические методы связываются **во время компиляции**,  в отличие от связывания виртуальных или не статических методов, которые связываются во время исполнения на реальном объекте. Следовательно, статические методы не могут быть переопределены в Java, т.к. полиморфизм во время выполнения не распространяется на них. Это важное ограничение, которое необходимо учитывать, объявляя метод статическим.
8)   Важным свойством статического блока является инициализация. Статические поля или переменные инициализируются после загрузки класса в память. Порядок инициализации сверху вниз, в том же порядке, в каком они описаны в исходном файле Java класса. 
9)   Во время сериализации, также как и transient переменные, статические поля не сериализуются. 
10)   И напоследок, поговорим о static import. Данный модификатор имеет много общего со стандартным оператором import, но в отличие от него позволяет импортировать один или все статические члены класса. При импортировании статических методов, к ним можно обращаться как будто они определены в этом же классе, аналогично при импортировании полей, мы можем получить доступ без указания имени класса.