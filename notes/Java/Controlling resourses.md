# Controlling resourses

Начиная с седьмой версии Java предлагает улучшенное управление ресурсами, которые должны быть закрыты после окончания работы с ними. 

Языковая конструкция:
`try-with-resources`

 Для того чтобы это автоматическое закрытие работало создан специальный интерфейс `AutoCloseable`. 
 
 Сигнатура:
 ```
 public interface AutoCloseable {
    void close() throws Exception;
}
```

`close()` автоматически вызывается для объектов обслуживаемых try-with-resourses.

Для классов которые не реализуют этот интерфейс нужно:
1. Наследуйте класс ресурса, который должен участвовать в конструкции try-with-resources. 
2. Использование в кострукции

#### Пример

После окончания работы с ITextRenderer, должен быть вызван метод finishPDF(). Обычно это делается и это в блоке finally. Но создавая новый класс, расширяющий ITextRenderer и реализующий интерфейс AutoCloseable, вы можете включить его в конструкцию try-with-resources. Новый класс AutoCloseableITextRenderer будет при этом выглядеть так:

```
public class AutoCloseableITextRenderer extends ITextRenderer implements AutoCloseable {
    @Override
    public void close() {
        super.finishPDF();
    }
}
```
Расширение оригинального класса в потомке, является наиболее разумным решением, поскольку новый класс по прежнему будет ITextRenderer. В случае, если исходный класс объявлен как final, необходимо использовать композицию. И вот как при этом будет выглядеть использование:
```
try (final AutoCloseableITextRenderer iTextRenderer = new AutoCloseableITextRenderer()) {
            ByteArrayOutputStream out; // contains the data to be converted to PDF, not shown here.
            iTextRenderer.setDocumentFromString(new String(out.toByteArray()));
            iTextRenderer.layout();
            iTextRenderer.createPDF(pdfOutputStream);
            pdfOutputStream.flush();
        }
```