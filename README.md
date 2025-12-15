# Portfolio Backend

A centralized backend API built with Go and Gin framework that connects and manages all deployed projects from web to mobile applications and vice versa, using GitHub Gist as the primary data store.

## 👨‍💻 Creator

**Ryann Kim Sesgundo [MPOP Reverse II]**

## 📋 Overview

This portfolio backend serves as a central hub for managing and connecting all your deployed projects across different platforms (web, mobile, etc.). The application leverages GitHub Gist as a primary data storage solution, providing a lightweight yet powerful way to store and retrieve project information, configurations, and metadata.

## 🎯 Main Purpose

The primary goal of this repository is to:
- **Centralize Project Management**: Unify all deployed projects under one backend API
- **Cross-Platform Integration**: Connect web and mobile applications seamlessly
- **GitHub Gist Integration**: Use GitHub Gist as a reliable, version-controlled data store
- **Unified API**: Provide consistent endpoints for all connected applications
- **Project Orchestration**: Enable communication and data sharing between different projects

## 🚀 Features

- **Lightweight & Fast**: Built with Go and Gin framework for optimal performance
- **GitHub Gist Integration**: Seamless integration with GitHub Gist API for data storage
- **Modular Architecture**: Clean separation of concerns with dedicated packages
- **Hot Reload**: Configured with Air for development hot reloading
- **RESTful API**: Standard HTTP methods support (GET, POST, PUT, DELETE)
- **Custom Middleware**: Flexible middleware system for request handling
- **Cross-Platform Ready**: Designed to serve both web and mobile applications
- **Environment Configuration**: Secure API key management through environment variables

## 🛠️ Tech Stack

- **Language**: Go 1.25.3
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Data Store**: GitHub Gist API
- **Development Tool**: [Air](https://github.com/cosmtrek/air) (Hot reload)
- **Package Management**: Go Modules
- **Authentication**: GitHub Personal Access Token

## 📁 Project Structure

```
portfolio-backend/
├── endpoints/              # API endpoint handlers
│   ├── index.go           # Main welcome endpoint
│   ├── fetch.go           # Data fetching endpoints
│   └── routers.go         # Route definitions and registration
├── middleware/            # Custom middleware
│   └── server_handler.go  # Server setup and route registration
├── utils/                 # Utility functions and structures
│   ├── structures.go      # Data structures and types
│   └── gist.go           # GitHub Gist API integration
├── tmp/                  # Temporary files (ignored by git)
├── .air.toml             # Air configuration for hot reload
├── .env                  # Environment variables (not tracked)
├── .gitignore            # Git ignore rules
├── go.mod                # Go module dependencies
├── go.sum                # Go module checksums
└── index.go              # Application entry point
```

## 🔧 Installation & Setup

### Prerequisites

- Go 1.25.3 or higher
- Git
- GitHub Personal Access Token (for Gist API access)

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

3. **Set up environment variables**

   Create a `.env` file in the root directory:
   ```bash
   API_KEY=your_github_personal_access_token
   ```

4. **Install Air for development (optional)**
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

| Method | Endpoint | Description                    |
| ------ | -------- | ------------------------------ |
| GET    | `/`      | Welcome endpoint               |
| GET    | `/fetch` | Fetch data from GitHub Gist    |

### Example Responses

#### Welcome Endpoint (`GET /`)
```json
{
  "error": "Hello World"
}
```

#### Fetch Endpoint (`GET /fetch`)
```json
{
  "message": "A successful connection"
}
```

### 404 Handler

All undefined routes return a custom HTML 404 page with a friendly message.

## 🔧 Configuration

### Server Configuration

The server can be configured to run on different ports by modifying the `StartServer()` function call in `index.go`:

```go
mw.StartServer(3000) // Run on port 3000
```

### GitHub Gist Configuration

Set up your GitHub Personal Access Token in the `.env` file:

```bash
API_KEY=ghp_your_personal_access_token_here
```

The token should have `gist` scope permissions.

### Air Configuration

Development hot reload is configured via `.air.toml`. Key settings:

- **Build Command**: `go build -o ./tmp/main .`
- **Watch Extensions**: `.go`, `.tpl`, `.tmpl`, `.html`
- **Excluded Directories**: `assets`, `tmp`, `vendor`, `testdata`

## 🏗️ Architecture

### GitHub Gist Integration

The application includes a comprehensive GitHub Gist API integration:

- **GET Operations**: Retrieve data from specified Gists
- **POST Operations**: Create or update Gist content
- **Authentication**: Secure API access using Personal Access Tokens
- **Error Handling**: Robust error handling for API operations

### Middleware System

The application uses a custom middleware system that:

- Registers routes dynamically
- Supports all HTTP methods (GET, POST, PUT, DELETE)
- Provides centralized server management
- Handles 404 errors with custom HTML responses

### Route Registration

Routes are defined as structs and automatically registered:

```go
var Routes = []utils.Route{
    Index,
    Fetch,
}
```

Each route includes:
- **Path**: URL endpoint
- **Method**: HTTP method (defaults to GET)
- **Handler**: Gin handler function

## 🔗 Integration Examples

### Connecting Web Applications

```javascript
// Example fetch from web application
fetch('http://localhost:8000/fetch')
  .then(response => response.json())
  .then(data => console.log(data));
```

### Connecting Mobile Applications

```dart
// Example HTTP request from Flutter/Dart
import 'package:http/http.dart' as http;

Future<Map<String, dynamic>> fetchData() async {
  final response = await http.get(
    Uri.parse('http://localhost:8000/fetch'),
  );
  return json.decode(response.body);
}
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 Development Guidelines

- Follow Go naming conventions and best practices
- Add appropriate error handling for all operations
- Update documentation for new endpoints
- Test GitHub Gist integration thoroughly
- Ensure environment variables are properly configured
- Test cross-platform compatibility

## 🔒 Security Considerations

- Never commit your `.env` file or API keys to version control
- Use environment variables for all sensitive configuration
- Regularly rotate your GitHub Personal Access Tokens
- Implement proper input validation for all endpoints
- Consider rate limiting for production deployments

## 📄 License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## 📞 Contact

**Ryann Kim Sesgundo**

- [GitHub](https://github.com/RyannKim327)
- Email: weryses19@gmail.com
- [LinkedIn](https://www.linkedin.com/in/ryannkim327/)

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin) for the excellent HTTP framework
- [Air](https://github.com/cosmtrek/air) for hot reload functionality
- GitHub for providing the Gist API as a data storage solution
- ChatGPT for helping me out

---

⭐ If you found this project helpful, please give it a star!

## 🚀 Future Enhancements

- [ ] Add authentication middleware
- [ ] Implement caching for Gist data
- [ ] Add WebSocket support for real-time updates
- [ ] Create admin dashboard for project management
- [ ] Add comprehensive API documentation with Swagger
- [ ] Implement project deployment automation
- [ ] Add monitoring and logging capabilities
