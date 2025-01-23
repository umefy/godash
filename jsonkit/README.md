# jsonkit

A comprehensive JSON and Protocol Buffer utility package for Go HTTP applications. Provides streamlined JSON handling with strict validation, Protocol Buffer support, and convenient HTTP request/response helpers.

## Features

- **HTTP Integration**: Direct JSON binding for HTTP requests and responses
- **Protocol Buffer Support**: Seamless JSON serialization for protobuf messages
- **Strict Validation**: Prevents unknown fields and malformed JSON
- **Error Handling**: Comprehensive error reporting for debugging
- **Type Safe**: Full Go type safety with generics support

## Installation

```bash
go get github.com/umefy/godash/jsonkit
```

## Quick Start

```go
package main

import (
    "net/http"
    "github.com/umefy/godash/jsonkit"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func handleUser(w http.ResponseWriter, r *http.Request) {
    var user User

    // Bind JSON request body to struct
    if err := jsonkit.BindRequestBody(r, &user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Send JSON response
    if err := jsonkit.JSONResponse(w, http.StatusOK, user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

## API Reference

### Core JSON Functions

#### Marshal

```go
func Marshal(v interface{}) ([]byte, error)
```

Marshals a Go struct to JSON bytes.

**Example:**

```go
user := User{Name: "Alice", Age: 30}
data, err := jsonkit.Marshal(user)
// data = []byte(`{"name":"Alice","age":30}`)
```

#### UnMarshal

```go
func UnMarshal(data []byte, v interface{}) error
```

Unmarshals JSON bytes to a Go struct with strict validation.

**Example:**

```go
data := []byte(`{"name":"Bob","age":25}`)
var user User
err := jsonkit.UnMarshal(data, &user)
// user = User{Name: "Bob", Age: 25}
```

### HTTP Request/Response Functions

#### BindRequestBody

```go
func BindRequestBody(r *http.Request, v interface{}) error
```

Binds JSON from HTTP request body to a Go struct.

**Features:**

- Disallows unknown fields
- Validates JSON structure
- Handles request body reading

**Example:**

```go
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := jsonkit.BindRequestBody(r, &user); err != nil {
        // Handle validation errors
        return
    }
    // Process user...
}
```

#### JSONResponse

```go
func JSONResponse(w http.ResponseWriter, statusCode int, v interface{}) error
```

Writes a Go struct as JSON to HTTP response.

**Features:**

- Sets proper Content-Type header
- Handles HTTP status codes
- Error handling for marshaling failures

**Example:**

```go
func handleGetUser(w http.ResponseWriter, r *http.Request) {
    user := User{Name: "Alice", Age: 30}
    if err := jsonkit.JSONResponse(w, http.StatusOK, user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

### Protocol Buffer Functions

#### MarshalProto

```go
func MarshalProto(v proto.Message) ([]byte, error)
```

Marshals a Protocol Buffer message to JSON.

**Features:**

- Emits unpopulated fields
- Optimized for protobuf serialization
- Compatible with protobuf JSON spec

**Example:**

```go
message := &pb.User{Name: "Alice", Age: 30}
data, err := jsonkit.MarshalProto(message)
```

#### UnMarshalProto

```go
func UnMarshalProto(data []byte, v proto.Message) error
```

Unmarshals JSON to a Protocol Buffer message.

**Example:**

```go
data := []byte(`{"name":"Bob","age":25}`)
message := &pb.User{}
err := jsonkit.UnMarshalProto(data, message)
```

#### BindProtoRequestBody

```go
func BindProtoRequestBody(r *http.Request, v proto.Message) error
```

Binds JSON from HTTP request body to a Protocol Buffer message.

**Example:**

```go
func handleCreateProtoUser(w http.ResponseWriter, r *http.Request) {
    message := &pb.User{}
    if err := jsonkit.BindProtoRequestBody(r, message); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Process protobuf message...
}
```

#### ProtoJSONResponse

```go
func ProtoJSONResponse(w http.ResponseWriter, statusCode int, v proto.Message) error
```

Writes a Protocol Buffer message as JSON to HTTP response.

**Example:**

```go
func handleGetProtoUser(w http.ResponseWriter, r *http.Request) {
    message := &pb.User{Name: "Alice", Age: 30}
    if err := jsonkit.ProtoJSONResponse(w, http.StatusOK, message); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

## Error Handling

The package provides comprehensive error handling:

```go
// Validation errors
if err := jsonkit.BindRequestBody(r, &user); err != nil {
    switch {
    case strings.Contains(err.Error(), "unexpected extra JSON data"):
        // Handle extra JSON data
    case strings.Contains(err.Error(), "unknown field"):
        // Handle unknown fields
    default:
        // Handle other JSON errors
    }
}
```

## Validation Features

- **Unknown Field Detection**: Prevents extra fields in JSON
- **Type Validation**: Ensures correct data types
- **Required Field Checking**: Validates struct tags
- **Malformed JSON Detection**: Catches syntax errors

## Performance Considerations

- **Efficient Memory Usage**: Minimal allocations
- **Streaming Support**: Direct HTTP body reading
- **Optimized Marshaling**: Fast JSON serialization
- **Zero-Copy Where Possible**: Efficient data handling

## Testing

```bash
go test ./jsonkit/...
```

## Dependencies

- `encoding/json` - Core JSON functionality
- `google.golang.org/protobuf` - Protocol Buffer support
- `net/http` - HTTP integration

## License

MIT License - see [LICENSE](../LICENSE) for details.
