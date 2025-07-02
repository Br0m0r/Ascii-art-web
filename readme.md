# ASCII-Art-Web

Web application to generate ASCII art with three banner styles: Standard, Shadow, and Thinkertoy.

## Authors
- **mfoteino**, **gpapadopoulos**, **vtsoucha**

## Quick Start

### Local Development
```bash
go run main.go
# Open http://localhost:8080
```

### Docker Deployment
```bash
# Build and run
docker build -t ascii-art-web .
docker run -p 8080:8080 ascii-art-web

# Or use audit script for testing
./audit.sh
```

## Features
- **Web Interface**: Text input with banner selection (radio buttons)
- **Three Banners**: Standard, Shadow, Thinkertoy
- **Export**: Download ASCII art as .txt files
- **Error Handling**: Proper HTTP status codes (200, 400, 404, 500)
- **Docker**: Containerized with security best practices

## Endpoints
- `GET /` - Main page
- `POST /ascii-art` - Generate ASCII art (with optional export)

## Project Structure
```
├── main.go              # Server logic and routing
├── loadAsciiArt.go      # Template loading
├── formatAsciiArt.go    # ASCII art generation
├── templates/           # HTML templates
├── banners/             # ASCII art templates
├── Dockerfile           # Multi-stage Docker build
├── .dockerignore        # Build optimization
└── audit.sh             # Docker testing script
```

## Requirements Implemented

### ✅ ASCII-Art-Web (Mandatory)
- Web GUI with three banner styles
- GET / and POST /ascii-art endpoints
- Proper HTTP status codes
- HTML templates and responsive design

### ✅ Dockerize (Optional)
- Multi-stage Dockerfile with security best practices
- Non-root user execution
- Health checks and proper metadata
- Optimized build with .dockerignore

### ✅ Export-File (Optional)
- File download with proper HTTP headers
- Content-Type, Content-Length, Content-Disposition
- Timestamped filenames for uniqueness

## Testing
Run the audit script to validate all functionality:
```bash
chmod +x audit.sh
./audit.sh
```

## Dependencies
- Go standard library only
- Docker (for containerization)
