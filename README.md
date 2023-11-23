# simple-todo-list

To Do List with:
- client in pure JS
- server in Go

### Start server
Run this command: `go run server/main.go`

- Create task with curl:
```
curl -X POST -H "Content-Type: application/json" -d '{"title": "Just do it", "status": 1}' http://localhost:8080/tasks/create
```
- Get all task:
```
curl -X GET http://localhost:8080/tasks/get
```

### Use client in browser
Open file `index.html` in browser
