# st0pcha/go-fiber-example

## 📒 About

**go-fiber-example** is a template for the [**Fiber (v3)**](https://gofiber.io) framework for Go. It includes a clean project structure and integrates live reloads using [**Air**](https://github.com/air-verse/air)

## 💻 Installation

1. Clone the repository:

```bash
git clone https://github.com/st0pcha/go-fiber-example.git
cd go-fiber-example
```

2. Create .env file

```
SERVER_HOST="0.0.0.0"
SERVER_PORT="8080"
SERVER_ALLOWED_ORIGINS="*" # "http://localhost:8080,http://localhost:3000"
```

3. Install dependencies

- with **go**:

```bash
go mod download
```

- with [**godo**](https://github.com/st0pcha/godo)

```bash
godo deps
```

4. Run

- with **go**:

```bash
go run cmd/api/main.go
```

- with [**godo**](https://github.com/st0pcha/godo)

```bash
godo run
```

## 🛫 How to run app using Air?

- with **go**:

```bash
air
```

- with [**godo**](https://github.com/st0pcha/godo)

```bash
godo dev
```
