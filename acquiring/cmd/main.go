package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/google/uuid"
)

func main() {
    a := AA{}
    v := `{"id": "d69375ea-bb50-4540-b9e2-055a7a91b06e", "name": "hello"}`

    fmt.Println(v)

    err := json.Unmarshal([]byte(v), &a)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", a)
}

// AA TODO
type AA struct {
    ID   uuid.UUID `json:"id"`
    Name string    `json:"name"`
}
