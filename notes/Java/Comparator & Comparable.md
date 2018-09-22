# Comparator & Comparable


Имеем класс: 

```
class Person{
     
    private String name;
    public Person(String value){
         
        name=value;
    }
    String getName(){return name;}
}
```

Для того, чтобы объекты Person можно было сравнить и сортировать, они должны применять интерфейс **Comparable<E>**. При применении интерфейса он типизируется текущим классом. Применим его к классу Person:

```
class Person implements Comparable<Person>{
     
    private String name;
    public Person(String value){
         
        name=value;
    }
    String getName(){return name;}
     
    public int compareTo(Person p){
     
        return name.compareTo(p.getName());
    }
}
```
Интерфейс Comparable содержит один единственный метод int compareTo(E item), который сравнивает текущий объект с объектом, переданным в качестве параметра. 

Интерфейс Comparator также содержит один метод:

```
public interface Comparator<E> {
 
    int compare(T a, T b);
}
```

Для применения интерфейса нам вначале надо создать класс компаратора, который реализует этот интерфейс:

```
class PersonComparator implements Comparator<Person>{
 
    public int compare(Person a, Person b){
     
        return a.getName().compareTo(b.getName());
    }
}
```
Начиная с JDK 8 в механизм работы компараторов были внесены некоторые дополнения. В частности, теперь мы можем применять сразу несколько компараторов по принципу приоритета, 
используя **thenComparing**

`Comparator<Person> pcomp = new PersonNameComparator().thenComparing(new PersonAgeComparator());
`

