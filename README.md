# simple-todo-list

To Do List with:
- client in pure JS
- server in Go

### Init local db
Install postgresql:
```
brew install postgresql
```
Set up database:
```
db-reset
```

### Start server
Run command for starting Go server:
```
go run server/main.go
```

### Test API locally
- Create task with curl:
```
curl -X POST -H "Content-Type: application/json" -d '{"title": "Just do it", "status": 1}' http://localhost:8080/tasks/
```
- Get all task:
```
curl -X GET http://localhost:8080/tasks/
```
- Update task by id:
```
curl -X PUT -H "Content-Type: application/json" -d '{"id":"99234191-3414-428f-9e6b-7bee7f00e815","title":"Just not do it","status":3}' http://localhost:8080/tasks/
```

### Use client in browser
Open file `index.html` in browser
