#  Todo Backend API (Go + Gin + PostgreSQL)

A simple, clean **Todo Backend API** built using **Golang**, **Gin framework**, and **PostgreSQL**.
This project demonstrates how to build a production-style REST API in Go without authentication.

---

##  Tech Stack

* **Language**: Go (Golang)
* **Framework**: Gin
* **Database**: PostgreSQL
* **ORM**: GORM
* **Environment Config**: godotenv

---

##  Project Structure

```
go_todo/
│── main.go
│── go.mod
│── .env
│
├── database/
│   └── db.go
│
├── models/
│   └── todo.go
│
├── controllers/
│   └── todo_handler.go
│
├── routes/
│   └── routes.go
│
├── services/
│   └── todo_service.go
```

---

##  Environment Variables

Create a `.env` file in the root directory:

```env
PORT=8080
DB_URL=postgres://todo_user:todo_pass@localhost:5432/todo_db?sslmode=disable
```

---

## ▶ Run the Project

```bash
go mod tidy
go run main.go
```

Server will start on:

```
http://localhost:8080
```

---

##  API Endpoints

###  Create Todo

**POST** `/todos`

```json
{
  "Title": "Learn Golang",
  "Done": false
}
```

---

###  Get All Todos

**GET** `/todos`

---

###  Get Todo by ID

**GET** `/todos/:id`

---

###  Update Todo

**PUT** `/todos/:id`

```json
{
  "Title": "Learn Gin Framework",
  "Done": true
}
```

---

###  Delete Todo

**DELETE** `/todos/:id`

---




