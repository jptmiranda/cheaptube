package utils

import (
    "fmt"
    "net/http"
)

func MultipartFile(req *http.Request, key string) ([]byte, string, error) {
    file, handler, err := req.FormFile(key)
    if err != nil {
        return nil, "", fmt.Errorf("error getting file from multipart request: %v", err)
    }
    defer file.Close()
    var data = make([]byte, handler.Size)
    _, err = file.Read(data)
    if err != nil {
        return nil, "", fmt.Errorf("error getting file from multipart request: %v", err)
    }

    return data, handler.Filename, nil
}
