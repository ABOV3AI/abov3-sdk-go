# ABOV3 Go SDK v0.14.0

Official Go SDK for ABOV3 AI - Genesis CodeForger Edition

---

## What's New in v0.14.0

### 🔄 Repository Migration
- ✅ Updated repository URL to `github.com/ABOV3AI/abov3-sdk-go`
- ✅ Migrated from personal repository to ABOV3AI organization
- ✅ Updated all import paths and documentation

### 🔄 Version Alignment
- Updated to align with ABOV3 Genesis CodeForger v0.1.6
- Synchronized with latest ABOV3 API endpoints

### 🐛 Bug Fixes
- Improved error handling for connection timeouts
- Enhanced retry logic for failed requests
- Fixed edge cases in JSON parsing

### 📚 Documentation
- Updated all repository references
- Enhanced README with usage examples
- Added comprehensive GoDoc comments

### 🔧 Technical Improvements
- Improved type safety with enhanced interfaces
- Better context handling for request cancellation
- Enhanced test coverage

---

## Installation

```bash
go get -u github.com/ABOV3AI/abov3-sdk-go@v0.14.0
```

---

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ABOV3AI/abov3-sdk-go"
)

func main() {
    // Initialize client
    client := abov3.NewClient()

    // List sessions
    sessions, err := client.Session.List(context.TODO(), abov3.SessionListParams{})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Sessions: %+v\n", sessions)

    // Create a new session
    session, err := client.Session.New(context.TODO(), abov3.SessionNewParams{})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Created session: %s\n", session.ID)
}
```

---

## Features

- ✨ **Type-Safe Interfaces** - Compile-time type checking
- 🔄 **Built-in Retry Logic** - Automatic retry for transient failures
- 📝 **Comprehensive Test Coverage** - Well-tested codebase
- 🔒 **Context Support** - Full context.Context support for cancellation
- 📊 **Structured Errors** - Detailed error information
- 🚀 **Zero Dependencies** - Only standard library dependencies

---

## Requirements

- Go 1.21 or later

---

## Key Changes from v0.13.0

### Breaking Changes
- Repository URL changed from `github.com/fajardofahad/abov3-sdk-go` to `github.com/ABOV3AI/abov3-sdk-go`
- Update your import statements:

```go
// Old
import "github.com/fajardofahad/abov3-sdk-go"

// New
import "github.com/ABOV3AI/abov3-sdk-go"
```

### Migration Guide

To upgrade from v0.13.x to v0.14.0:

1. Update your go.mod:
```bash
go get -u github.com/ABOV3AI/abov3-sdk-go@v0.14.0
```

2. Update import statements in your code:
```bash
# Unix/macOS
find . -name "*.go" -type f -exec sed -i '' 's|github.com/fajardofahad/abov3-sdk-go|github.com/ABOV3AI/abov3-sdk-go|g' {} \;

# Linux
find . -name "*.go" -type f -exec sed -i 's|github.com/fajardofahad/abov3-sdk-go|github.com/ABOV3AI/abov3-sdk-go|g' {} \;
```

3. Run your tests:
```bash
go test ./...
```

---

## Documentation

- 📚 **ABOV3 Documentation**: https://www.abov3.ai/docs
- 📦 **Go SDK Repository**: https://github.com/ABOV3AI/abov3-sdk-go
- 📖 **Go Package Documentation**: https://pkg.go.dev/github.com/ABOV3AI/abov3-sdk-go

---

## Support

- 💬 **Issues**: https://github.com/ABOV3AI/abov3-sdk-go/issues
- 🏠 **Website**: https://www.abov3.ai
- 📧 **Email**: support@abov3.ai

---

**Author**: Fahad Ibn Omar Fajardo
**License**: MIT