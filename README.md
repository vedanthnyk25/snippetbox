# SnippetBox - A Secure Web Application

SnippetBox is a full-stack web application built with Go that allows users to create, store, and share text snippets. The application follows modern web development practices and implements robust security measures.

## 🚀 Features

### Core Functionality
- Create, read, and manage text snippets
- User authentication and authorization
- Session management
- Responsive web interface
- RESTful API endpoints

### Security Features
- TLS encryption with modern cipher suites (X25519, CurveP256)
- Secure session management with MySQL-backed storage
- Protection against common web vulnerabilities
- Input validation and sanitization
- Secure password handling

### Technical Features
- Structured logging with slog
- Template caching for improved performance
- Database connection pooling
- Graceful server shutdown
- Error handling and recovery

## 📁 Project Structure

```
snippetbox/
├── cmd/
│   └── web/           # Application entry point and server setup
├── internal/
│   ├── models/        # Database models and business logic
│   ├── validator/     # Input validation
│   └── assert/        # Custom assertions and helpers
├── ui/
│   ├── html/         # HTML templates
│   ├── static/       # Static assets (CSS, JS, images)
│   └── efs.go        # Embedded file system
├── tls/              # TLS certificates
└── tmp/              # Temporary files and build artifacts
```

## 🛠️ Technology Stack

- **Backend**: Go (Golang)
- **Database**: MySQL
- **Session Management**: SCS (Session Cookie Store)
- **Form Processing**: go-playground/form
- **Templates**: Go's html/template
- **Security**: TLS, Secure Cookies

## 🔧 Setup and Installation

1. **Prerequisites**
   - Go 1.21 or later
   - MySQL 8.0 or later
   - TLS certificates (for development, you can use self-signed certificates)

2. **Database Setup**
   ```sql
   CREATE DATABASE snippetbox;
   USE snippetbox;
   ```

3. **Configuration**
   - Set up your database connection string in the environment variables
   - Place your TLS certificates in the `tls` directory

4. **Running the Application**
   ```bash
   go run ./cmd/web
   ```

## 🔒 Security Measures

- **TLS Configuration**
  - Modern cipher suites
  - Secure curve preferences
  - Proper certificate handling

- **Session Security**
  - Secure cookie settings
  - Session timeout management
  - MySQL-backed session storage

- **Database Security**
  - Parameterized queries
  - Connection pooling
  - Proper error handling

## 📝 API Endpoints

- `GET /` - Home page
- `GET /snippet/:id` - View a specific snippet
- `POST /snippet/create` - Create a new snippet
- `GET /user/signup` - User registration
- `POST /user/signup` - Process registration
- `GET /user/login` - User login
- `POST /user/login` - Process login

## 🧪 Testing

The application includes comprehensive testing:
- Unit tests for models
- Integration tests for handlers
- Security testing for authentication

## 📈 Performance Optimizations

- Template caching
- Database connection pooling
- Static file serving
- Proper timeout configurations

## 🔍 Logging and Monitoring

- Structured logging with slog
- Error tracking
- Performance monitoring
- Request logging

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👥 Authors

- Vedanth Nayak

## 🙏 Acknowledgments

- Go standard library
- All open-source contributors
- The Go community 
