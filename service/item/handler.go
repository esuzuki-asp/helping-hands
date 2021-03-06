package item

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	Items []Item `json:"items"`
}

type addToCartRequest struct {
	ItemID string `json:"item_id"`
	UserID string `json:"item_id"`
}

type addToCartResponse struct {
}

type createItemRequest struct {
	UserID          int64  `json:"user_id"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Category        string `json:"category"`
	Type            string `json:"type"`
	Subtype         string `json:"subtype"`
	CategoryFilters string `json:"category_filter"`
	AvailableStart  string `json:"available_start"`
	AvailableEnd    string `json:"available_end"`
	LocationID      int64  `json:"location_id"`
}

type createItemResponse struct {
	ID int64 `json:"id"`
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	logrus.Println(req.URL)
	_, method := parseRequestPath(req.URL.Path)
	logrus.Println(method)

	switch method {
	case "ping":
		res.Write([]byte(`pong`))
		return
	case "getItems":
		h.serveGetItems(res, req)
		return
	case "addToCart":
		h.serveAddToCart(res, req)
		return
	case "createItem":
		h.serveCreateItem(res, req)
		return
	default:
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
}

func (h Handler) serveGetItems(res http.ResponseWriter, req *http.Request) {
	var getItemsReq getItemsRequest
	var getItemsRes getItemsResponse

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&getItemsReq)
	if err != nil {
		logrus.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	err = h.Service.GetItems(&getItemsReq, &getItemsRes)
	if err != nil {
		logrus.Error("serveGetLocations: ", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(getItemsRes)
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

func (h Handler) serveAddToCart(res http.ResponseWriter, req *http.Request) {
	var addToCartReq addToCartRequest
	var addToCartRes addToCartResponse

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&addToCartReq)
	if err != nil {
		logrus.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	err = h.Service.AddToCart(&addToCartReq, &addToCartRes)
	if err != nil {
		logrus.Error("serveGetLocations: ", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(addToCartRes)
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

func (h Handler) serveCreateItem(res http.ResponseWriter, req *http.Request) {
	var createItemReq createItemRequest
	var createItemRes createItemResponse

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&createItemReq)
	if err != nil {
		logrus.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	err = h.Service.CreateItem(&createItemReq, &createItemRes)
	if err != nil {
		logrus.Error("serveGetLocations: ", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(createItemRes)
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
