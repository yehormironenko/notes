# Reflection #

[источник](http://javadevblog.com/polnoe-rukovodstvo-po-java-reflection-api-refleksiya-na-primerah.html)

***Рефлексия в Java*** — *это механизм, с помощью которого можно вносить изменения и получать информацию о классах, интерфейсах, полях и методах во время выполнения, при этом не зная имен этих классов, методов и полей. Кроме того, Reflection API дает возможность создавать новые экземпляры классов, вызывать методы, а также получать или устанавливать значения полей.*


##### Минусы рефлексии:

+ **Низкая производительность** - поскольку рефлексия в Java определяет типы динамически, то она сканирует classpath, чтобы найти класс для загрузки, в результате чего снижается производительность программы.


+ **Ограничения системы безопасности** -  рефлексия требует разрешения времени выполнения, которые не могут быть доступны для систем, работающих под управлением менеджера безопасности (Java Security Manager).


+ **Нарушения безопасности приложения** - с помощью рефлексии мы можем получить доступ к части кода, к которой мы не должны получать доступ. Например, мы можем получить доступ к закрытым полям класса и менять их значения. Это может быть серьезной угрозой безопасности.


+ **Сложность в поддержке** - код, написанный с помощью рефлексии трудно читать и отлаживать, что делает его менее гибким и трудно поддерживаемым.


## Java Reflection

### Работа с классами

Важные методы при работе с `java.lang.Class`:

+ Все типы в Java, включая примитивные типы и массивы имеют связанный с ними `java.lang.Class` объект. Если мы знаем название класса во время компиляции, то сможем получить объект следующим образом:
` Class mClassObject = SomeObject.class`


+ Если мы не знаем имя во время компиляции, но знаем имя класса во время выполнения, то можно сделать так:
`Class mClassObject = Class.forName("здесь имя класса")`

#### Получаем название класса

+ Методом getName() , который вернет полное имя класса с пакетом:

`String fullClassName = mClassObject.getName();`

+ C помощью метода getSimpleName(), который вернет только название класса без имени пакета:

`String justClassName = mClassObject.getSimpleName();`

#### Работа с модификаторами доступа

Мы можем получить доступ к модификаторам доступа с помощью `Class` объекта. Модификаторы представляют собой ключевые слова `public, static, private` и т.д. Мы можем получить модификаторы с помощью метода `getModifiers()`:

`Class mClassObject = SomeObject.class
int classModifiers = mClassObject.getModifiers();`

Мы можем проверить модификаторы, используя следующие методы в классе `java.lang.reflect.Modifier`:

`Modifier.isAbstract(int modifiers)
Modifier.isFinal(int modifiers)
Modifier.isInterface(int modifiers)
Modifier.isNative(int modifiers)
Modifier.isPrivate(int modifiers)
Modifier.isProtected(int modifiers)
Modifier.isPublic(int modifiers)
Modifier.isStatic(int modifiers)
Modifier.isStrict(int modifiers)
Modifier.isSynchronized(int modifiers)
Modifier.isTransient(int modifiers)
Modifier.isVolatile(int modifiers)`

В параметры метода просто передадим modifiers и каждый из этих методов вернет `true` или `false`.


#### Получение информации о пакете

`Class mClassObject = SomeObject.class
Package package = mClassObject.getPackage();`


#### Получаем Superclass

`Class mClassObject = SomeObject.class
Class superclass = mClassObject.getSuperclass();`


#### Реализованные интерфейсы

`Class mClassObject = SomeObject.class
Class[] interfaces = mClassObject.getInterfaces();`

**Обратите внимание**: Метод вернул только интерфейсы, которые реализует указанный класс, а не его суперкласс. Для того, чтобы получить полный список интерфейсов, реализованных в этом классе, вы должны обратиться как к текущему классу, так и к его суперклассу.


### Рефлексия в Java: Конструкторы


С помощью Java Reflection API можно получать информацию про конструкторы классов и создавать объекты во время выполнения. Это делается с помощью класса Java java.lang.reflect.Constructor.

#### Получаем конструкторы класса

`Class mClassObject = SomeObject.class 
Constructor[] constructors = mClassObject.getConstructors();`

В коде выше мы получили массив конструкторов класса и теперь можем работать с ними. Если известны параметры конкретного конструктора, то массив можно не получать, а работать уже с известным. Например, следующий конструктор класса принимает String качестве параметра:

`Class mClassObject = SomeObject.class 
Constructor constructor = mClassObject.getConstructor(new Class[]{String.class});`

Если мы ошиблись с аргументом конструктора, то будет выброшено исключение `NoSuchMethodException`.

#### Получаем параметры конструктора

`Class[] parameterTypes = constructor.getParameterTypes();`

#### Создание объектов с помощью конструкторов

`Constructor constructor = SomeObject.class.getConstructor(String.class);

SomeObject myObject = (SomeObject) constructor.newInstance("здесь какой-то строковый аргумент");`


### Рефлексия в Java: Поля

Используется класс `java.lang.reflect.Field`


#### Получаем поля класса

`Class mClassObject = SomeObject.class 
Field[] fields = mClassObject.getFields();`

Если вы знаете имя поля:

`Class mClassObject = SomeObject.class 
Field field = mClassObject.getField("fieldName");`

#### Получаем название поля

`Field field = mClassObject.getField("fieldName"); 
String fieldName = field.getName();`

#### Получаем тип поля

`Field field = mClassObject.getField("fieldName");
Object fieldType = field.getType();`

#### Получаем и устанавливаем значения полей

`Class mClassObject = SomeObject.class 
Field field = mClassObject.getField("fieldName");
SomeObject instance = new SomeObject();
Object value = field.get(instance);
field.set(instance, value);`


### Java Reflection API: Методы


Используя ревфлексию в Java мы можем получать информацию о методах и вызывать их в рантайме. Для этого используется класс `java.lang.reflect.Method`

Например, у нас есть метод под названием "sayHello", который принимает `String` в качестве параметра. Получить объект `Method` для него можно так:

`Class mClassObject = SomeObject.class 
Method method = mClassObject.getMethod("sayHello", new Class[]{String.class});`

Если такого метода нет, то будет выброшен `NoSuchMethodException`.

Если метод sayHello() без параметров, то нужно передать null в методе getMethod():

`Method method = mClassObject.getMethod("sayHello", null);`

#### Параметры метода и типы возвращаемых значений

Получить параметры метода можно так:

`Class[] parameterTypes = method.getParameterTypes();`

получить тип возвращаемого значения:

`Class returnType = method.getReturnType();`

#### Вызов метода с помощью Java рефлексии

`Class mClassObject = SomeObject.class 
Method method = mClassObject.getDeclaredMethod("sayHello", parameterTypes)
method.invoke(objectToInvokeOn, params)`


### Java Reflection API: Методы установки и получения значения

Смотрим на примере определения метода как сеттера или геттера с помощью рефлексии в Java:

````
public static void printGettersOrSetters(Class aClass){
  Method[] methods = aClass.getMethods();
  for(Method method : methods){
    if(isGetter(method)) System.out.println("getter: " + method);
    if(isSetter(method)) System.out.println("setter: " + method);
  }
}

public static boolean isGetter(Method method){
  if (!method.getName().startsWith("get")) {
     return false;
  } 
  if (method.getParameterTypes().length != 0) {
     return false;  
  } 
  if (void.class.equals(method.getReturnType()) {
     return false;
  } 
  return true;
}

public static boolean isSetter(Method method){
  if (!method.getName().startsWith("set")) {
     return false;
  } 
  if (method.getParameterTypes().length != 1) {
     return false;
  } 
  return true;
}
````

### Java Reflection API: Приватные поля и методы

С помощью рефлексии в Java можно получить доступ к private полям других классов. Это очень удобно, например, для модульного тестирования.

