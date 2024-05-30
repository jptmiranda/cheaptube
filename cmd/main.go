package main

import (
	"bytes"
	"fmt"
	"github.com/jptmiranda/cheaptube/db"
	"github.com/jptmiranda/cheaptube/utils"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	dbConn, err := db.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	query := db.New(dbConn)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{id}", func(resp http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		video, err := query.GetVideo(req.Context(), int64(id))
		if err != nil {
			resp.WriteHeader(http.StatusNotFound)
			return
		}

		resp.Header().Set("Content-Type", "application/octet-stream")
		resp.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", video.Name))
		fileReader := bytes.NewReader(video.Data)
		_, err = io.Copy(resp, fileReader)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte("error writing file"))
		}
	})

	mux.HandleFunc("POST /", func(resp http.ResponseWriter, req *http.Request) {
		if !utils.IsContentType(req, "multipart/form-data") {
			resp.WriteHeader(http.StatusBadRequest)
			resp.Write([]byte("bad request"))
			return
		}

		data, name, err := utils.MultipartFile(req, "video")
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte("error reading request"))
			return
		}

		err = query.CreateVideo(req.Context(), db.CreateVideoParams{
			Name: name,
			Data: data,
		})
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(fmt.Errorf("error saving video: %w", err).Error()))
		}
	})

	server := &http.Server{
		Addr:    net.JoinHostPort("localhost", "8080"),
		Handler: mux,
	}

	fmt.Println("server running")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
