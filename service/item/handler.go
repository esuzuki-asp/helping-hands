package item

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

const PathPrefix = "/item/"

type Handler struct {
	Service
}

func NewHandler() Handler {
	repo := NewRepository()
	service := NewService(repo)
	return Handler{Service: service}
}

type getItemsRequest struct {
	Category        string `json:"category"`
	Type            string `json:"type"`
	Subtype         string `json:"subtype"`
	CategoryFilters string `json:"category_filter"`
}

type getItemsResponse struct {
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	logrus.Println(req.URL)
	_, method := parseRequestPath(req.URL.Path)
	logrus.Println(method)

	switch method {
	case "ping":
		res.Write([]byte(`pong`))
		return
	default:
		print("unsupported request")
		// Error?? through res
		return
	}
}

//func (h Handler) serveGetItems(res http.ResponseWriter, req *http.Request) {
//	var getItemsRequest getItemsRequest
//	jsonpb.Unmarshal(req.Body, &getItemsRequest)
//}

func parseRequestPath(path string) (string, string) {
	parts := strings.Split(path, "/")
	logrus.Println(parts)
	if len(parts) < 3 {
		return "", ""
	}
	return parts[1], parts[2]
}
