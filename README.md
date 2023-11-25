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
go run server/cmd/main.go
```

### Test API locally with curl

1) Register user

- Request:

```
curl -X POST -d '{"email": "email@google.com", "password": "123"}' http://localhost:8080/register
```

---

2) Login user

- Request:

```
curl -X POST -d '{"email": "email@google.com", "password": "123"}' http://localhost:8080/login
```

- Response:

```
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDA5NDg1NTgsImlhdCI6IjIwMjMtMTEtMjVUMjA6NDI6MzguMjU3NTExWiIsIm5iZiI6MTcwMDk0NDk1OCwic3ViIjoiM2MyYjNiZTItMDkzOC00Y2FmLWJlOGEtMTJiOGJmMjRhYTE5In0.fqRR3WodYEVu8MTVCk0BI4L8lgocWfAgZhrkk2ISosY"}
```

---

3) Create task

- Request:

```
curl -X POST -H "Authorization: Bearer YOUR_ACCESS_TOKEN" -d '{"title": "Just do it", "status": 1}' http://localhost:8080/tasks
```

- Response:

```
{"id":"c3db9081-9523-42e3-b6b0-efe65a4a8b89","userId":"3c2b3be2-0938-4caf-be8a-12b8bf24aa19","title":"Just do it","status":1,"createdAt":"2023-11-25T20:46:04.656102+06:00","updatedAt":"2023-11-26T02:46:04.656102+06:00"}
```

---

4) Get tasks for user

- Request:

```
curl -X GET -H "Authorization: Bearer YOUR_ACCESS_TOKEN" http://localhost:8080/tasks
```

- Response:

```
[{"id":"c3db9081-9523-42e3-b6b0-efe65a4a8b89","title":"Just do it","status":1,"createdAt":"2023-11-25T20:46:04.656102+06:00","updatedAt":"2023-11-26T02:46:04.656102+06:00"}]
```

---

5) Update task by id

- Request:

```
curl -X PUT -H "Authorization: Bearer YOUR_ACCESS_TOKEN" -d '{"id":"c3db9081-9523-42e3-b6b0-efe65a4a8b89","title":"Just not do it","status":3}' http://localhost:8080/tasks
```

- Response:

```
{"id":"c3db9081-9523-42e3-b6b0-efe65a4a8b89","title":"Just not do it","status":3,"createdAt":"2023-11-25T20:46:04.656102+06:00","updatedAt":"2023-11-26T02:57:41.768572+06:00"}
```

### Use client in browser

Open file `index.html` in browser
