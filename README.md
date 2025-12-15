# Portfolio Backend

A lightweight and modular backend API built with Go and Gin framework for portfolio applications.

## 👨‍💻 Creator

**Ryann Kim Sesgundo [MPOP Reverse II]**

## 📋 Overview

This is a RESTful API backend designed to serve portfolio data and handle various endpoints for a portfolio website. The project follows a clean, modular architecture with separate concerns for routing, middleware, and endpoint handling.

## 🚀 Features

- **Lightweight & Fast**: Built with Go and Gin framework for optimal performance
- **Modular Architecture**: Clean separation of concerns with dedicated packages
- **Hot Reload**: Configured with Air for development hot reloading
- **RESTful API**: Standard HTTP methods support (GET, POST, PUT, DELETE)
- **Custom Middleware**: Flexible middleware system for request handling
- **Error Handling**: Comprehensive error responses with proper HTTP status codes

## 🛠️ Tech Stack

- **Language**: Go 1.25.3
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Development Tool**: [Air](https://github.com/cosmtrek/air) (Hot reload)
- **Package Management**: Go Modules

## 📁 Project Structure

```
portfolio-backend/
├── endpoints/           # API endpoint handlers
│   ├── index.go        # Main endpoint handlers
│   └── routers.go      # Route definitions
├── middleware/         # Custom middleware
│   └── server_handler.go # Server setup and route registration
├── tmp/               # Temporary files (ignored by git)
├── .air.toml          # Air configuration for hot reload
├── .gitignore         # Git ignore rules
├── go.mod             # Go module dependencies
├── go.sum             # Go module checksums
└── index.go           # Application entry point
```

## 🔧 Installation & Setup

### Prerequisites

- Go 1.25.3 or higher
- Git

### Installation Steps

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd portfolio-backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Install Air for development (optional)**
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

## 🚀 Running the Application

### Development Mode (with hot reload)

```bash
air
```

### Production Mode

```bash
go run index.go
```

The server will start on port `8000` by default.

## 📡 API Endpoints

### Base URL
```
http://localhost:8000
```

### Available Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/`      | Welcome endpoint |

### Example Response

```json
{
  "error": "Hello World"
}
```

### 404 Handler

All undefined routes return:
```json
{
  "error": "Miss ko na sya"
}
```

## 🔧 Configuration

### Server Configuration

The server can be configured to run on different ports by modifying the `StartServer()` function call in `index.go`:

```go
mw.StartServer(3000) // Run on port 3000
```

### Air Configuration

Development hot reload is configured via `.air.toml`. Key settings:

- **Build Command**: `go build -o ./tmp/main .`
- **Watch Extensions**: `.go`, `.tpl`, `.tmpl`, `.html`
- **Excluded Directories**: `assets`, `tmp`, `vendor`, `testdata`

## 🏗️ Architecture

### Middleware System

The application uses a custom middleware system that:
- Registers routes dynamically
- Supports all HTTP methods (GET, POST, PUT, DELETE)
- Provides centralized server management

### Route Registration

Routes are defined in `endpoints/routers.go` and automatically registered through the middleware system:

```go
var Routes = []Route {
    {
        Method: "GET",
        Path: "/",
        Handler: Index,
    },
}
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 Development Guidelines

- Follow Go naming conventions
- Add appropriate error handling
- Update documentation for new endpoints
- Test your changes before submitting

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 📞 Contact

**Ryann Kim Sesgundo**
- GitHub: [Your GitHub Profile]
- Email: [Your Email]

---

⭐ If you found this project helpful, please give it a star!
