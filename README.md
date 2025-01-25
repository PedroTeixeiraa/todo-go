# Todo App in Go

A task management application (Todo List) developed in Go, using the Gin Framework and PostgreSQL.

## 🚀 Technologies

- [Go](https://golang.org/)
- [Gin Framework](https://github.com/gin-gonic/gin)

## 📋 Prerequisites

- Go 1.16 or higher
- PostgreSQL

## 🔧 Installation

1. Clone the repository
```bash
git clone https://github.com/PedroTeixeiraa/todo-go.git
```

2. Navigate to the project directory
```bash
cd todo-go
```

3. Install dependencies
```bash
go mod download
```

4. Configure the environment variables

5. Run the command to create the database
```bash
docker compose up -d
```

6. Start the application
```bash
go run main.go
```

The server will be running at `http://localhost:8080`

## 🛠️ API Endpoints

### Tasks

- `GET /tasks` - List all tasks
- `POST /tasks` - Create a new task
- `GET /tasks/:id` - Get a specific task
- `PUT /tasks/:id` - Update a task
- `DELETE /tasks/:id` - Delete a task

### Task Format
```json
{
  "title": "Example task",
  "description": "Task description",
  "completed": false
}
```

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## ✒️ Author

* **Pedro Teixeira** - [PedroTeixeiraa](https://github.com/PedroTeixeiraa)

---
⌨️ with ❤️ by [Pedro Teixeira](https://github.com/PedroTeixeiraa) 😊