package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"transaction-service/config"
	"transaction-service/errs"
	"transaction-service/pkg/log"
	"transaction-service/repository"
)

const (
	CT     = "Content-Type"
	CTJson = "application/json"
)

// Server struct
type Server struct {
	svc repository.DBOps
	cfg *config.Config
}

// NewServer constructor
func NewServer(cfg *config.Config, s repository.DBOps) *Server {
	return &Server{cfg: cfg, svc: s}
}

// Run server
func (s *Server) Run() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/accounts", s.createAccount).Methods(http.MethodPost)
	r.HandleFunc("/accounts/{account_id}", s.getAccount).Methods(http.MethodGet)
	r.HandleFunc("/transactions", s.createTransaction).Methods(http.MethodPost)

	return r
}

//GetParams unmarshalls request body to required struct
func (s *Server) GetParams(o interface{}, Response http.ResponseWriter, Request *http.Request) (err error) {
	ct := getContentType(Request)
	if ct != CTJson {
		Response.Header().Set("Accept", CTJson)
		err = errors.New("Unsupported media type")
		return
	}

	body, _ := ioutil.ReadAll(Request.Body)
	// Restore the io.ReadCloser to its original state
	Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if len(body) < 1 {
		err = errs.ErrorEmptyBodyContent
		return
	}

	err = json.Unmarshal(body, o)
	if err != nil {
		err = errs.ErrorRequestBodyInvalid
		return
	}

	return
}

// GetContentType ...
func getContentType(req *http.Request) (ct string) {
	ct = req.Header.Get(CT)
	return
}

func (s *Server) FormatException(r http.ResponseWriter, err error) {
	s.JSON(r, http.StatusBadRequest, errs.FormatErrorResponse(err))
}

// JSON sends a JSON response body
func (s *Server) JSON(r http.ResponseWriter, code int, content interface{}) {
	if fmt.Sprint(content) == "[]" {
		emptyResponse, _ := json.Marshal(make([]int64, 0))
		Output(r, code, CTJson, emptyResponse)
		return
	}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	enc.Encode(content)
	Output(r, code, CTJson, b.Bytes())
}

// Output sets a full HTTP output detail
func Output(r http.ResponseWriter, code int, ctype string, content []byte) {
	log.Info("Response ", zap.Any("Message", string(content)))
	r.Header().Set("Content-Type", ctype)
	r.WriteHeader(code)
	r.Write(content)
}
