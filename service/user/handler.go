package user

import (
	"encoding/json"
	"helping-hands/service/item"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const PathPrefix = "/user/"

type Handler struct {
	Service
}

func NewHandler() Handler {
	repo := NewRepository()
	service := NewService(repo)
	return Handler{Service: service}
}

type getCartRequest struct {
	UserID int64 `json:"user_id"`
}

type getCartResponse struct {
	Items []item.ItemWithLocation `json:"items"`
}

type getOrdersRequest struct {
	UserID int64 `json:"user_id"`
}

type getOrdersResponse struct {
	Items []OrderWithFullDetails `json:"items"`
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	logrus.Println(req.URL)
	_, method := parseRequestPath(req.URL.Path)
	logrus.Println(method)

	switch method {
	case "ping":
		res.Write([]byte(`pong`))
		return
	case "getCart":
		h.serveGetCart(res, req)
		return
	case "getOrders":
		h.serveGetOrders(res, req)
		return
	default:
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
}

func (h Handler) serveGetCart(res http.ResponseWriter, req *http.Request) {
	var getCartReq getCartRequest
	var getCartRes getCartResponse

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&getCartReq)
	if err != nil {
		logrus.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	err = h.Service.GetCart(&getCartReq, &getCartRes)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(getCartRes)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	res.WriteHeader(http.StatusOK)

	if _, err := res.Write(buf); err != nil {
		logrus.Error("Failed to write response")
	}
}

func (h Handler) serveGetOrders(res http.ResponseWriter, req *http.Request) {
	var getOrdersReq getOrdersRequest
	var getOrdersRes getOrdersResponse

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&getOrdersReq)
	if err != nil {
		logrus.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	err = h.Service.GetOrders(&getOrdersReq, &getOrdersRes)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(getOrdersRes)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	res.WriteHeader(http.StatusOK)

	if _, err := res.Write(buf); err != nil {
		logrus.Error("Failed to write response")
	}
}

func (h Handler) serveRequest(
	res http.ResponseWriter,
	req *http.Request,
	f func(*interface{}, *interface{}) error,
	internalRes *interface{}, internalReq *interface{},
) {

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(internalReq)
	if err != nil {
		logrus.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	err = f(internalReq, internalRes)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(internalRes)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	res.WriteHeader(http.StatusOK)

	if _, err := res.Write(buf); err != nil {
		logrus.Error("Failed to write response")
	}
}

func parseRequestPath(path string) (string, string) {
	parts := strings.Split(path, "/")
	logrus.Println(parts)
	if len(parts) < 3 {
		return "", ""
	}
	return parts[1], parts[2]
}
