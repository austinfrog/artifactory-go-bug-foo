package foo

import "github.com/google/uuid"

func Hello() string {
    return "Hello from foo v2: " + uuid.New().String()
}
