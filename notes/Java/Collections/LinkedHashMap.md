# LinkedHashMap

![Image](https://habrastorage.org/storage2/6c8/9a8/15d/6c89a815d82b6b9f1c362a50f182e7ee.png)

*Данная структура является симбиозом связанных списков и хэш-мапов.*

### Создание объекта

`Map<Integer, String> linkedHashMap = new LinkedHashMap<Integer, String>();`

Объект **linkedHashMap**, помимо свойств унаследованных от **HashMap** (такие как table, loadFactor, threshold, size, entrySet и т.п.), так же содержит два доп. свойства:
+ **header** — «голова» двусвязного списка. При инициализации указывает сам на себя;
+ **accessOrder** — указывает каким образом будет осуществляться доступ к элементам при использовании итератора. При значении *true* — по порядку последнего доступа (об этом в конце). При значении *false() доступ осуществляется в том порядке, в каком элементы были вставлены.

### Добавление элементов

linkedHashMap.put(1, "obj1");
При добавлении элемента, первым вызывается метод `createEntry(hash, key, value, bucketIndex)` (по цепочке `put()` -> `addEntry()` -> `createEntry()`)

#### accessOrder == true

Когда свойство `accessOrder` имеет значение true. В такой ситуации поведение **LinkedHashMap** меняется и при вызовах методов `get()` и `put()` порядок элементов будет изменен — элемент к которому обращаемся будет помещен в конец.


### Итераторы

```
// 1.
Iterator<Entry<Integer, String>> itr1 = linkedHashMap.entrySet().iterator();
while (itr1.hasNext()) {
    Entry<Integer, String> entry = itr1.next();
    System.out.println(entry.getKey() + " = " + entry.getValue());
}

// 2.
Iterator<Integer> itr2 = linkedHashMap.keySet().iterator();
while(itr2.hasNext())
    System.out.println(itr2.next());

// 3.
Iterator<String> itr3 = linkedHashMap.values().iterator();
while (itr3.hasNext())
    System.out.println(itr3.next());
```

### Итоги


Данная структура может слегка уступать по производительности родительскому **HashMap**, при этом время выполнения операций `add()`, `contains()`, `remove()` остается константой — O(1). Понадобится чуть больше места в памяти для хранения элементов и их связей, но это совсем небольшая плата за дополнительные фишечки.

Вообще, из-за того что всю основную работу на себя берет родительский класс, серьезных отличий в реализации HashMap и LinkedHashMap не много. Можно упомянуть о мелких:

+ Методы `transfer()` и `containsValue()` устроены чуть проще из-за наличия двунаправленной связи между элементами;
+ В классе `LinkedHashMap.Entry` реализованы методы `recordRemoval()` и `recordAccess()` (тот самый, который помещает элемент в конец при accessOrder = true). В **HashMap** оба этих метода пустые.

## Ссылки 
[original page](https://habrahabr.ru/post/129037/)