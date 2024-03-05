// lib/http/server.go
package http

import (
	"dsched/db"
	"fmt"
	"io/ioutil"
	stdhttp "net/http"
	"strconv"
	"strings"
)

type Server struct {
	database db.Database
}

func (s *Server) ServeHTTP(writer stdhttp.ResponseWriter, request *stdhttp.Request) {
	if request.Method == "POST" {
		bodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(400)
			_, _ = writer.Write([]byte("error while reading the body"))
			return
		}
		id, err := s.database.StoreText(string(bodyBytes))
		if err != nil {
			writer.WriteHeader(500)
			_, _ = writer.Write([]byte("error while storing the text"))
			return
		}
		writer.WriteHeader(200)
		writer.Write([]byte(strconv.Itoa(int(id))))
	} else {
		pathSplit := strings.Split(request.URL.Path, "/")
		id, err := strconv.Atoi(pathSplit[1])
		if err != nil {
			writer.WriteHeader(400)
			fmt.Println(err)
			_, _ = writer.Write([]byte("error while reading ID from URL"))
			return
		}
		text, err := s.database.GetTextByID(id)
		if err != nil {
			writer.WriteHeader(400)
			fmt.Println(err)
			_, _ = writer.Write([]byte("error while reading text from database"))
			return
		}
		_, _ = writer.Write([]byte(text))
	}
}

func NewServer(db db.Database) *Server {
	return &Server{database: db}
}
