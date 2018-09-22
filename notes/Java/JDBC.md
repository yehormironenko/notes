# JDBC

**JDBC Driver** – (Java DataBase Connectivity — соединение с базами данных на Java) — платформенно-независимый промышленный стандарт взаимодействия Java-приложений с различными СУБД, реализованный в виде пакета java.sql, входящего в состав Java SE.


### Подключнеие к БД

**Mysql**

```
Class.forName("com.mysql.jdbc.Driver");
Connection conn = DriverManager.getConnection("jdbc:mysql://hostname:port/dbname","username", "password");
conn.close();
```

**PostgreSQL**

```
Class.forName("org.postgresql.Driver");
Connection connection = DriverManager.getConnection("jdbc:postgresql://hostname:port/dbname","username", "password");
connection.close();
```


**Oracle**

```
Class.forName("oracle.jdbc.driver.OracleDriver");
Connection connection = DriverManager.getConnection("jdbc:oracle:thin:@localhost:1521:mkyong","username","password");
connection.close();
```


В 1-й строке мы указываем наш JDBC драйвер. Не забудьте добавить его в ClassPath иначе его компилятор его не увидит.
Во 2-й строке JDBC Manager который открывает соединение с базой данных и обеспечит нам дальнейшее обращение к ней.
И последняя строка закрывает соединение с БД.

Желательно строку для определения JDBC поместить в блок try для того чтобы контролировать его наличия в вашем приложении.

```
try {
	Class.forName("com.mysql.jdbc.Driver");
} catch (ClassNotFoundException e) {
	System.out.println("Where is your MySQL JDBC Driver?");
	e.printStackTrace();
	return;
}
```

### Создание таблиц БД

Вынесем в отдельный метод соединение с БД.
```
private static Connection getDBConnection() {
	Connection dbConnection = null;
	try {
		Class.forName(DB_DRIVER);
	} catch (ClassNotFoundException e) {
		System.out.println(e.getMessage());
	}
	try {
		dbConnection = DriverManager.getConnection(DB_CONNECTION, DB_USER,DB_PASSWORD);
		return dbConnection;
	} catch (SQLException e) {
		System.out.println(e.getMessage());
	}
	return dbConnection;
}
```

Этот метод будет создавать в БД таблицу:

```
private static void createDbUserTable() throws SQLException {
	Connection dbConnection = null;
	Statement statement = null;

	String createTableSQL = "CREATE TABLE DBUSER("
			+ "USER_ID NUMBER(5) NOT NULL, "
			+ "USERNAME VARCHAR(20) NOT NULL, "
			+ "CREATED_BY VARCHAR(20) NOT NULL, "
			+ "CREATED_DATE DATE NOT NULL, " + "PRIMARY KEY (USER_ID) "
			+ ")";

	try {
		dbConnection = getDBConnection();
		statement = dbConnection.createStatement();

                // выполнить SQL запрос
		statement.execute(createTableSQL);
		System.out.println("Table \"dbuser\" is created!");
	} catch (SQLException e) {
		System.out.println(e.getMessage());
	} finally {
		if (statement != null) {
			statement.close();
		}
		if (dbConnection != null) {
			dbConnection.close();
		}
	}
}
```

и в `main` методе вызываем метод `createDbTable()` который создаст таблицу в БД.

```
public static void main(String[] argv) {
	try {
		createDbUserTable();
	} catch (SQLException e) {
		System.out.println(e.getMessage());
	}
}
```

### Получение данных с БД

`String selectTableSQL = "SELECT USER_ID, USERNAME from DBUSER";
`

Выполняем запрос:

```
try {
	dbConnection = getDBConnection();
	statement = dbConnection.createStatement();

	// выбираем данные с БД
	ResultSet rs = statement.executeQuery(selectTableSQL);

	// И если что то было получено то цикл while сработает   
	while (rs.next()) {
		String userid = rs.getString("USER_ID");
		String username = rs.getString("USERNAME");

		System.out.println("userid : " + userid);
		System.out.println("username : " + username);
	}
} catch (SQLException e) {
	System.out.println(e.getMessage());
}
```

