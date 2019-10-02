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

#### Frontend
Frontend-part locate to ```index.html```

#### Backend
For running server, in folder with project run: 
```go build && ./go```

#### P.S.
Default WebSocket and server port is 9900
