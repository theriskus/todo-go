# ToDo List

Requiriements:
1. Go 1.12 => above
2. MySQL 5.8 => above (Testing on MySQL 8 Community Edition)

For running this app you needed this packages:
```sh
go get github.com/gorilla/websocket
go get github.com/go-sql-driver/mysql
```
Also, you need rename ```config.example.json``` to ```config.json``` and edit this

Sample sql:
```sql
CREATE DATABASE todo;
USE todo;
CREATE TABLE quests
(
    id     INT(11) AUTO_INCREMENT PRIMARY KEY,
    quest  VARCHAR(3000) NOT NULL,
    status int(1) NOT NULL
);
```

#### Frontend
Frontend-part locate to ```index.html```

#### Backend
For running server, in folder with project run: 
```go build && ./go```

#### P.S.
Default WebSocket and server port is 9900
