package location

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const PathPrefix = "/location/"

type Handler struct {
	Service
}

func NewHandler() Handler {
	repo := NewRepository()
	service := NewService(repo)
	return Handler{Service: service}
}

type createLocationRequest struct {
	City         string `json:"city"`
	Country      string `json:"country"`
	MeetingPoint string `json:"meeting_point"`
}

type createLocationResponse struct {
	LocationID int64 `json:"location_id"`
}

type getLocationRequest struct {
	ID int64 `json:"id"`
}

type getLocationResponse struct {
	ID           int64  `json:"id"`
	City         string `json:"city"`
	Country      string `json:"country"`
	MeetingPoint string `json:"meeting_point"`
}

type getLocationsRequest struct {
	City         string `json:"city"`
	Country      string `json:"country"`
	MeetingPoint string `json:"meeting_point"`
}

type getLocationsResponse struct {
	Locations []Location
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	logrus.Println(req.URL)
	_, method := parseRequestPath(req.URL.Path)
	logrus.Println(method)

	switch method {
	case "ping":
		res.Write([]byte(`pong`))
		return
	case "getLocation":
		h.serveGetLocation(res, req)
		return
	case "getLocations":
		h.serveGetLocations(res, req)
		return
	case "createLocation":
		h.serveCreateLocation(res, req)
		return
	default:
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
}

func (h Handler) serveGetLocation(res http.ResponseWriter, req *http.Request) {
	var getLocationReq getLocationRequest
	var getLocationRes getLocationResponse

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&getLocationReq)
	if err != nil {
		logrus.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	err = h.Service.GetLocation(&getLocationReq, &getLocationRes)
	if err != nil {
		logrus.Error("serveGetLocation: ", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(getLocationRes)
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

func (h Handler) serveGetLocations(res http.ResponseWriter, req *http.Request) {
	var getLocationsReq getLocationsRequest
	var getLocationsRes getLocationsResponse

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&getLocationsReq)
	if err != nil {
		logrus.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	err = h.Service.GetLocations(&getLocationsReq, &getLocationsRes)
	if err != nil {
		logrus.Error("serveGetLocations: ", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(getLocationsRes)
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

func (h Handler) serveCreateLocation(res http.ResponseWriter, req *http.Request) {
	var createLocationReq createLocationRequest
	var createLocationRes createLocationResponse

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&createLocationReq)
	if err != nil {
		logrus.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	err = h.Service.CreateLocation(&createLocationReq, &createLocationRes)
	if err != nil {
		logrus.Error("serveCreateLocation: ", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(createLocationRes)
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
