Part 2.
ООП
=======================

***Назовите принципы ООП и расскажите о каждом.***

**Инкапсуляция
Абстракция
Полиморфизм
Наследование**
Абстрация - обобщение как результат (написание кода с использованием абстрактных классов и интерфейсов, когда во внимание не берется реализация деталей, а выделяются лишь какие-то общие признаки на основании которых выстраивается иерархия наследования)

Инкапсуляция — это механизм языка, позволяющий ограничить прямой доступ к полям и методам класса извне (другим классам), с целью запретить бесконтрольную модификацию состояния объекта или вызов методов, которые также могут изменить его состояние.
(public, private, protected)

Наследование - принцип позволяющий писать классы на основе существующего (используя готовую функциональность)

Полиморфизм (многообразие форм) — способность вызывать нужные методы у объектов (имеющих разные типы) но находящихся в одной иерархии


***Дайте определение понятию “класс”.***
Шаблон или описание объекта.
Класс это набор каких-либо параметров и функций. 

***Что такое поле/атрибут класса?***
Допустим мы создаём класс человек. У человека есть характеристики - вес, рост. Это и будет полями класса. Грубо говоря - это характеристики класса.

***Как правильно организовать доступ к полям класса?***
Обычно поля класса инкапсулируют. Делают их модификаторы видимости private, что бы можно было их использовать только в данном классе и реализуют методы, которые могут позволить обратится к полям этого класса - геттеры и сеттеры.

***Дайте определение понятию “конструктор”.***
Когда мы создаём новый объект класса, первым делом вызывается его конструктор. По умолчанию при создании нового класса конструктор является пустым, но его можно переопределить и с помощью конструктора выполнять какие-либо действия в момент создания объекта.

***Чем отличаются конструкторы по-умолчанию, копирования и конструктор с параметрами?***
Конструктор по умолчанию вызовется когда мы создадим новый объект класса. Test t = new Test(); В этот момент сработает конструктор по умолчанию, но что делать, если нам надо создать новый объект с какими либо предопределёнными параметрами, тогда мы можем воспользоваться конструктором с параметрами. Test t = new Test(1, 4); Конструктор с параметрами необходимо реализовать в классе Test. Так же рядом может быть реализован конструктор без параметров, если понадобится создать объект класса не указав его параметры. Конструктор копирования, отвечает за копирование объекта. Если нам необходимо будет использовать копирование объекта, нам необходимо реализовать этот конструктор в классе, предварительно Имплементируясь от интерфейса Cloneable.
`` 
public DummyBean(DummyBean another) {
 this.dummy = another.dummy; // you can access 
  }``

***Какие модификации уровня доступа вы знаете, расскажите про каждый из них.***
![](http://javaway.info/wp-content/uploads/2016/01/image002-300x151.jpg)
***Расскажите об особенностях класса с единственным закрытым (private) конструктором.***
Если не рассматривать `reflection` то можно создать объект класса только через статический метод, который должен вернуть объект класса. 
Похоже на singleton, Разница состоит в том, что во втором случае, объект будет существовать в единственном экземпляре.

***О чем говорят ключевые слова “this”, “super”, где и как их можно использовать?***
this - представляет текущий экземпляр класса
super - текущий экземпляр родительского класса

И this, и super могут использоваться внутри конструкторов для вызова других конструкторов по цепочке, нпр., this() и super() вызывают конструктор без аргументов наследующего и родительского классов соответственно.



***Дайте определение понятию “метод”.***
Метод -  это законченная последовательность действий (инструкций), направленных на решение отдельной задачи.
***Что такое сигнатура метода?***
Сигнатура метода — это имя метода плюс параметры (причем порядок параметров имеет значение).
***Какие методы называются перегруженными?***
Одинковое название метода и разные сигнатуры
***Могут ли нестатические методы перегрузить статические?***
нет, происходит переопределение

***Расскажите про переопределение методов.***
переопределение метода: ovveride. Часто может быть в классах потомках. 

***Может ли метод принимать разное количество параметров (аргументы переменной длины)?***
да, может принимать массивы. 
через ..

***Можно ли сузить уровень доступа/тип возвращаемого значения при переопределении метода?***

Можно расширить, но не сузить (***перепроверь***)

***Как получить доступ к переопределенным методам родительского класса?***
используя `super()`. Или прямым вызовом из класса

***Какие преобразования называются нисходящими и восходящими?***

![](https://metanit.com/java/tutorial/pics/3.12.png)

Восходящее преобразование `upcasting` происходит автоматически
```
Object tom = new Person("Tom");
Object sam = new Employee("Sam", "Oracle");
Object kate = new Client("Kate", "DeutscheBank", 2000);
Person bob = new Client("Bob", "DeutscheBank", 3000);
Person alice = new Employee("Alice", "Google");
```

Обратное не всегда верно. Например, объект Person не всегда является объектом Employee или Client. Поэтому нисходящее преобразование или `downcasting` от суперкласса к подклассу автоматически не выполняется. В этом случае нам надо использовать операцию преобразования типов.
```
Object sam = new Employee("Sam", "Oracle");      
// нисходящее преобразование от Object к типу Employee
Employee emp = (Employee)sam;
emp.display();
System.out.println(emp.getCompany());
```

***Чем отличается переопределение от перегрузки?***
переопределение используется при наследовании, перегрузка при полиморфизме.

***Где можно инициализировать статические/нестатические поля?***
Статические поля или переменные инициализируются после загрузки класса в память. Порядок инициализации сверху вниз, в том же порядке, в каком они описаны в исходном файле Java класса.
***Зачем нужен оператор instanceof?***
для определения к какому типу принадлежит объект

***Зачем нужны и какие бывают блоки инициализации?***
нестатический (instance initializer)

`{
   //...
}`
статический (class initializer)


`static {
   //...
}`
блоки инициализации делают код читабельнее, и позволяют вызывать любые методы.


***Каков порядок вызова конструкторов и блоков инициализации двух классов: потомка и его предка?***
Сначала вызываются все статические блоки от первого предка до последнего наследника. Потом попарно вызываются динамический блок инициализации и конструктор в той же последовательности (от предка до последнего потомка). 

***Где и для чего используется модификатор abstract?***
Абстрактным называется класс, на основе которого не могут создаваться объекты. При этом наследники класса могут быть не абстрактными, на их основе объекты создавать, соответственно, можно. Для того, чтобы превратить класс в абстрактный перед его именем надо указать модификатор abstract. 

Абстрактный метод — метод, который не имеет реализации. Если в классе есть хотя бы один абстрактный метод, то весь класс должен быть объявлен абстрактным.

Использование абстрактных классов и методов позволяет описать некую абстракцию, которая должна быть реализована в других классах. Например, мы можем создать абстрактный класс Fighter и объявить в нём абстрактный метод fight(). Т.к. стилей борьбы может быть много, то, например, для JudoFighter extends Fighter метод fight() будет описывать приемы в стиле дзюдо и т.д. 
***Можно ли объявить метод абстрактным и статическим одновременно?***
Нет. Получите: Illegal combination of modifiers: ‘abstract’ and ‘static’. Модификатор abstract говорит, что метод будет реализован в другом классе, а static наоборот указывает, что этот метод будет доступен по имени класса. 

***Что означает ключевое слово static?***
Модификатор static говорит о том, что метод или поле класса принадлежат не объекту, а классу. Т.е. доступ можно будет получить и не создавая объекта класса. Поля помеченные static инициализируются при инициализации класса. К примеру, Class.forName(«MyClass», true, currentClassLoader), где второй параметр указывает на необходимость проведения инициализации. 

На методы, объявленные как static, накладывается ряд ограничений. 

Они могут вызывать только другие статические методы. 
Они должны осуществлять доступ только к статическим переменным. 
Они не могут ссылаться на члены типа this или super. 


***К каким конструкциям Java применим модификатор static?***
+ К методу.
+ К внутреннему классу.
+ К полю.
+ К импортируемым классам (с 5-ой java). Например, import static org.junit.Assert.assertThat;

***Что будет, если в static блоке кода возникнет исключительная ситуация?***
Если в явном виде написать любое исключение в static-блоке, то компилятор не скомпилирует исходники. Это все от того, что компилятор умный. В остальном, взаимодействие с исключениями такое же как и в любом другом месте. Если unchecked исключение вывалится в static-блоке, то класс не будет инициализирован. 
Какое исключение выбрасывается при ошибке в блоке инициализации? 

Для static:
java.lang.ExceptionInInitializerError — если исключение наследуется от RuntimeException.
Для init:
exception, который и вызвал исключение, если он наследуется от RuntimeException.
Верно для static и init:
java.lang.Error — если исключение вызвано Error.
java.lang.ThreadDeath — смерть потока. Ничего не вываливается.

***Можно ли перегрузить static метод?***
Перегрузить можно, но переопределить нельзя.

***Что такое статический класс, какие особенности его использования?***
Это вложенный класс, который может обращаться только к статическим полям обертывающего его класса, в том числе и приватным. Доступ к нестатическим полям обрамляющего класса может быть осуществлен только через ссылку на экземпляр обрамляющего объекта. К классу высшего уровня модификатор static неприменим. 
***Какие особенности инициализации final static переменных?***
Переменные должны быть инициализированы во время объявления или в static блоке. 

***Как влияет модификатор static на класс/метод/поле?***
Модификатор static говорит о том, что метод или поле класса принадлежат не объекту, а классу. 
Внутри static метода нельзя вызвать не статический метод по имени класса. 
Про static класс смотрите ответ выше. 
***О чем говорит ключевое слово final?***
Может быть применено к полям, методам или классам. В зависимости к какой сущности приложено данное ключевое слово — будет и различный смысл в его применении. 

+ Для класса. Класс помеченный при помощи final не может иметь наследников.
+ Для метода. Метод помеченный при помощи final не может быть переопределен в классах наследниках.
+ Для поля. Поле помеченное при помощи слова final не может изменить свое значение после инициализации (инициализируется либо при описании, либо в конструкторе/статическом блоке).
+ Значение локальных переменных, а так же параметров метода помеченных при помощи слова final не могут быть изменены после присвоения.

***Дайте определение понятию “интерфейс”.***
Ключевое слово `interface` используется для создания полностью абстрактных классов. Создатель интерфейса определяет имена методов, списки аргументов и типы возвращаемых значений, но не тела методов. 

Наличие слова interface означает, что именно так должны выглядеть все классы, которые реализуют данный интерфейс. Таким образом, любой код, использующий конкретный интерфейс, знает только то, какие методы вызываются для этого интерфейса, но не более того.
```
public interface SomeName{ 
    void method(); 
    int getSum(); 
}
```
***Какие модификаторы по умолчанию имеют поля и методы интерфейсов?***
Интерфейс может содержать поля, но они автоматически являются статическими (static) и неизменными (final). Все методы и переменные неявно объявляются как public. 


***Почему нельзя объявить метод интерфейса с модификатором final или static?***
Вообще с 8й версии можно static, но нужно чтобы было тело метода. Например
```
public interface Shape { 
    static void draw() { 
        System.out.println("Wow! It is impossible!"); 
    }; 
}
```
final модификатор просто бессмысленный. Все методы по умолчанию абстрактные, т.е. их невозможно создать не реализовав где-то еще, но это нельзя будет сделать, если у метода идентификатор final. 
***Какие типы классов бывают в java (вложенные… и.т.д.)***
***Какие особенности создания вложенных классов: простых и статических.***
+ Обычные классы (`Top level classes`)
+ Интерфейсы (`Interfaces`)
+ Перечисления (`Enum`)
+ Статические вложенные классы (`Static nested classes`)
  + Есть возможность обращения к внутренним статическим полям и методам класса обертки.
  + Внутренние статические интерфейсы могут содержать только статические методы.

+ Внутренние классы-члены (`Member inner classes`)
   + Есть возможность обращения к внутренним полям и методам класса обертки.
   + Не может иметь статических объявлений.
   + Нельзя объявить таким образом интерфейс. А если его объявить без идентификатора `static`, то он автоматически будет добавлен.
   + Внутри такого класса нельзя объявить перечисления.
   + Если нужно явно получить `this` внешнего класса — `OuterClass.this`

+ Локальный класс (Local inner classes)
  + Видны только в пределах блока, в котором объявлены.
  + Не могут быть объявлены как `private/public/protected` или `static` (по этой причине интерфейсы нельзя объявить локально).
  + Не могут иметь внутри себя статических объявлений (полей, методов, классов).
  + Имеют доступ к полям и методам обрамляющего класса.
  + Можно обращаться к локальным переменным и параметрам метода, если они объявлены с модификатором `final`.
+ Анонимные классы (Anonymous inner classes)
  + Локальный класс без имени.

***Что вы знаете о вложенных классах, зачем они используются? Классификация, варианты использования, о нарушении инкапсуляции.***
***В чем разница вложенных и внутренних классов?***
***Какие классы называются анонимными?***
Вложенный класс — это класс, который объявлен внутри объявления другого класса. Вложенные классы делятся на статические и нестатические (non-static). Собственно нестатические вложенные классы имеют и другое название — внутренние классы (inner classes).

Внутренние классы в Java делятся на такие три вида:
+ внутренние классы-члены (member inner classes);
+ локальные классы (local classes);
+ анонимные классы (anonymous classes).
Внутренние классы-члены ассоциируются не с самим внешним классом, а с его экземпляром. При этом они имеют доступ ко всем его полям и методам.

Локальные классы (local classes) определяются в блоке Java кода. На практике чаще всего объявление происходит в методе некоторого другого класса. Хотя объявлять локальный класс можно внутри статических и нестатических блоков инициализации.

Анонимный класс (anonymous class) — это локальный класс без имени.

Использование вложенных классов всегда приводит к некоторому нарушению инкапсуляции — вложенный класс может обращаться к закрытым членам внешнего класса (но не наоборот!). Если это обстоятельство учитывается в архитектуре вашего приложения, не стоит уделять ему особого внимания, поскольку внутренний класс всего лишь является специализированным членом внешнего класса.

***Каким образом из вложенного класса получить доступ к полю внешнего класса?***
Если вложенный класс не статический и поле не статическое, то можно просто обратиться к этому полю из внутреннего класса, если только у внутреннего класса не существует поля с таким же литералом, в этом случае нужно обращаться через ссылку на внешний класс так — OuterClass.this.имяПоля


***Каким образом можно обратиться к локальной переменной метода из анонимного класса, объявленного в теле этого метода? Есть ли какие-нибудь ограничения для такой переменной?***
Также как и локальные классы, анонимные могут захватывать переменные, доступ к локальным переменным происходит по тем же правилам:

+ Анонимный класс имеет доступ к полям внешнего класса.
+ Анонимный класс не имеет доступ к локальным переменным области, в которой он определен, если они не финальные (final) или неизменяемые (effectively final).
+ Как и у других внутренних классов, объявление переменной с именем, которое уже занято, затеняет предыдущее объявление.
+ Вы не можете определять статические члены анонимного класса.
 
 
***Как связан любой пользовательский класс с классом Object?***
Все классы являются наследниками суперкласса Object. Это не нужно указывать явно. В результате объект Object o может ссылаться на объект любого другого класса. 

***Расскажите про каждый из методов класса Object.***
+ public final native Class getClass() — возвращает в рантайме класс данного объекта.
+ public native int hashCode() — возвращает хеш-код
+ public boolean equals(Object obj) — сравнивает объекты.
+ protected native Object clone() throws CloneNotSupportedException — клонирование объекта
+ public String toString() — возвращает строковое представление объекта.
+ public final native void notify() — просыпается один поток, который ждет на “мониторе” данного объекта.
+ public final native void notifyAll() — просыпаются все потоки, которые ждут на “мониторе” данного объекта.
+ public final native void wait(long timeout) throws InterruptedException — поток переходит в режим ожидания в течение указанного времени.
+ public final void wait() throws InterruptedException — приводит данный поток в ожидание, пока другой поток не вызовет notify() или notifyAll() методы для этого объекта.
+ public final void wait(long timeout, int nanos) throws InterruptedException — приводит данный поток в ожидание, пока другой поток не вызовет notify() или notifyAll() для этого метода, или пока не истечет указанный промежуток времени.
+ protected void finalize() throws Throwable — вызывается сборщиком мусора, когда garbage collector определил, что ссылок на объект больше нет.
***Что такое метод equals(). Чем он отличается от операции ==.***
***Если вы хотите переопределить equals(), какие условия должны удовлетворяться для переопределенного метода?***
***Если equals() переопределен, есть ли какие-либо другие методы, которые следует переопределить?***
***В чем особенность работы методов hashCode и equals? Каким образом реализованы методы hashCode и equals в классе Object? Какие правила и соглашения существуют для реализации этих методов? Когда они применяются?***
`equals` метод определенный в классе `Object`
и служит для сравнения объектов.
`==`  - ссылочное сравнение , оба объекта указывают на ожно и то же место в памяти.
`equals()` - оценивает сравнение значений в объектах.

При переопределении необходимо переопределить метод `hashCode`, хэшкоды должны совпадать.

+ Если хеш-коды разные, то и входные объекты гарантированно разные.
+ Если хеш-коды равны, то входные объекты не всегда равны.
Ситуация, когда у разных объектов одинаковые хеш-коды называется — коллизией. Вероятность возникновения коллизии зависит от используемого алгоритма генерации хеш-кода.


***Какой метод возвращает строковое представление объекта?***
`toString()`
***Что будет, если переопределить equals не переопределяя hashCode? Какие могут возникнуть проблемы?***
Нарушается связь. Так для объекта HashMap это может привести к тому, что пара, которая была помещена в Map возможно не будет найдена в ней при обращении к Map, если используется новый экземпляр ключа. 
***Есть ли какие-либо рекомендации о том, какие поля следует использовать при подсчете hashCode?***
Те, которые используют при определении метода equals(). Хэш код должен быть равномерно распределен на области возможных принимаемых значений.

***Как вы думаете, будут ли какие-то проблемы, если у объекта, который используется в качестве ключа в hashMap изменится поле, которое участвует в определении hashCode?***
Будут. При обращении по ключу мы можем не найти значение.

***Чем отличается абстрактный класс от интерфейса, в каких случаях что вы будете использовать?***
Абстрактные классы используются только тогда, когда есть «is a» тип отношений; интерфейсы могут быть реализованы классами которые не связаны друг с другом.

Абстрактный класс может реализовывать методы; интерфейс может реализовывать статические методы начиная с 8й версии.

Интерфейс может описывать константы и методы. Все методы интерфейса по умолчанию являются публичными (public) и абстрактными (abstract), а поля — public static final. С java 8 в интерфейсах можно реализовывать default и статические методы.

В Java класс может наследоваться (реализовывать) от многих интерфейсов, но только от одного абстрактного класса.

С абстрактными классами вы теряете индивидуальность класса, наследующего его; с интерфейсами вы просто расширяете функциональность каждого класса.

***Можно ли получить доступ к private переменным класса и если да, то каким образом?***
Reflection
***Что такое volatile и transient? Для чего и в каких случаях можно было бы использовать default?***
`volatile`  — не используется кэш (имеется ввиду область памяти в которой JVM может сохранять локальную копию переменной, чтобы уменьшить время обращения к переменной) при обращении к полю. Для volatile переменной JVM гарантирует синхронизацию для операций чтения/записи, но не гарантирует для операций изменения значения переменной.

`transient` — указание того, что при сериализации/десериализации данное поле не нужно сериализовать/десериализовывать.


***Расширение модификаторов при наследовании, переопределение и сокрытие методов. Если у класса-родителя есть метод, объявленный как private, может ли наследник расширить его видимость? А если protected? А сузить видимость?***
Действует общий принцип: расширять видимость можно, сужать нельзя. private методы видны только внутри класса, для потомков они не видны. Поэтому их и расширить нельзя.

***Имеет ли смысл объявлять метод private final?***
Нет, такой метод и так не виден для наследников, а значит не может быть ими переопределен.

***Какие особенности инициализации final переменных?***

+ Для поля. Поле помеченное при помощи слова final не может изменить свое значение после инициализации. Не статическое final поле можно инициализировать: при описании, в конструкторе (во всех), в статическом блоке, в динамическом блоке. Статическое final поле (static final) инициализируется либо в статическом блоке, либо при описании.
+ Значение локальных переменных, а так же параметров метода помеченных при помощи слова final не могут быть изменены после присвоения.


***Что будет, если единственный конструктор класса объявлен как final?***
К конструктору не применимо ключевое слово `final`.


***Что такое finalize? Зачем он нужен? Что Вы можете рассказать о сборщике мусора и алгоритмах его работы.***
Метод finalize() вызывается перед тем, как объект будет удален garbage collector (сборщик мусора, далее gc). Существует много различных реализаций gc. Основа работы следующая: gc помечает объекты, на которые больше не ссылаются другие объекты для их удаления. Затем на одном из проходов помеченные объекты удаляются.
Вызов finalize() не гарантируется, т.к. приложение может быть завершено до того, как будет запущена ещё одна сборка мусора. Да, можно отменить сборку объекта с помощью метода finalize(), присвоив его ссылку какому-то статическому методу.



***Почему метод clone объявлен как protected? Что необходимо для реаМожет быть применено к полям, методам или классам. В зависимости к какой сущности приложено данное ключевое слово — будет и различный смысл в его применении.***
Это указывает на то, что хоть метод и есть в классе Object и разработчик желает им воспользоваться, то его нужно переопределить. Для этого нужно реализовать интерфейс Clonable, чтобы соблюсти контракт.

***Знакомы ли Вам какие-либо паттерны проектирования?***
+ шаблонный
+ стратегия
+ синглтон
+ итератор
***Напишите Singleton… А с ленивой загрузкой. А если он должен быть потоко-безопасным? А в каких случаях ленивая загрузка хуже?***

***Singleton*** - одиночка, порождающий паттерн проектирования, который гарантирует существование только одного объекта определенного класса, а также позволяет достучаться до этого объекта из любого места программы. 

Признаки применения паттерна: Одиночку можно определить по статическому создающему методу, который возвращает один и тот же объект.

Топорно реализовать Одиночку очень просто — достаточно скрыть конструктор и предоставить статический создающий метод.

```java
public final class Singleton {
    private static volatile Singleton instance;
    public String value;

    private Singleton(String value) {
        this.value = value;
    }

    public static Singleton getInstance(String value) {
        if (instance == null) {
            synchronized (Singleton.class) {
                if (instance == null) {
                    instance = new Singleton(value);
                }
            }
        }
        return instance;
    }
}
```
```java
public class Singleton {
	public static final Singleton INSTANCE = new Singleton();
}
```

Ленивая инициализация:
```java
public class Singleton {
	private static Singleton instance;
	
	public static synchronized Singleton getInstance() {
		if (instance == null) {
			instance = new Singleton();
		}
		return instance;
	}
}
```

***Что можете сказать про MVC? Нарисуйте диаграмму и объясните, как MVC работает.***

![](https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/MVC-Process.png/240px-MVC-Process.png)

**Модель** - Модель предоставляет данные и методы работы с ними: запросы в базу данных, проверка на корректность. Модель не зависит от представления (не знает как данные визуализировать) и контроллера (не имеет точек взаимодействия с пользователем) просто предоставляя доступ к данным и управлению ими.

**Представление** - Представление отвечает за получение необходимых данных из модели и отправляет их пользователю. Представление не обрабатывает введённые данные пользователя 

**Контроллер** - Контроллер обеспечивает «связи» между пользователем и системой. Контролирует и направляет данные от пользователя к системе и наоборот. Использует модель и представление для реализации необходимого действия.


***Напишите функцию вычисления факториала.***
***Дана функция вычисления чисел Фибоначчи, известно, что она работает. Найдите логическую ошибку. Оцените сложность получившегося алгоритма.***