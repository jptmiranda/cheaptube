package utils

import (
    "net/http"
    "strings"
)

func IsContentType(req *http.Request, name string) bool {
    contentType := strings.Split(req.Header.Get("Content-Type"), ";")[0]
    if contentType == name {
        return true
    }
    return false
}
