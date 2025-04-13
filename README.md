
# ✅ ToDo CLI App

A simple and efficient command-line ToDo app built with Golang and [Cobra](https://github.com/spf13/cobra).

---

## ✨ Features

- ➕ Add new tasks
- 📋 List all tasks
- ✅ Mark tasks as done
- ❌ Delete tasks
- 💾 Persists tasks in a local `tasks.json` file

---

## 🚀 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/yourusername/todo-cli.git
cd todo-cli
```

### 2. Build the App

```bash
go build -o todo
```

### 3. Run Commands

```bash
./todo add "Buy groceries"
./todo list
./todo done 1
./todo delete 1
```

---

## 📦 Commands

| Command         | Description                     | Example                          |
|-----------------|---------------------------------|----------------------------------|
| `add [task]`     | Add a new task                  | `./todo add "Read Go docs"`      |
| `list`           | List all tasks                  | `./todo list`                    |
| `done [id]`      | Mark a task as done             | `./todo done 1`                  |
| `delete [id]`    | Delete a task                   | `./todo delete 1`                |

---

## 📁 Task Storage

- Tasks are saved in a `tasks.json` file in the app's root directory.
- Each task has an `ID`, `Title`, and `Done` status.

---

## 🛠 Built With

- [Go](https://golang.org)
- [Cobra CLI](https://github.com/spf13/cobra)

---

