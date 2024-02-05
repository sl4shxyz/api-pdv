package rest_error

import "net/http"

type RestError struct {
  Message string   `json:"message"`
  Err     string   `json:"error"`
  Code    int      `json:"code"`
  Causes  []Causes `json:"causes"`
}

type Causes struct {
  Field   string `json:"field"`
  Message string `json:"message"`
}

func (re *RestError) Error() string {
  return re.Message
}

func NewRestError(message, err string, code int, causes []Causes) *RestError {
  return &RestError {
    Message: message,
    Err: err,
    Code: code,
    Causes: causes,
  }
}

func NewBadRequestError(message string) *RestError {
  return &RestError {
    Message: message,
    Err: "bad request",
    Code: http.StatusBadRequest,
  }
}

func NewBadRequestValidationError(message string, causes []Causes) *RestError {
  return &RestError {
    Message: message,
    Err: "bad request",
    Code: http.StatusBadRequest,
    Causes: causes,
  }
}

func NewInternalServerError(message string) *RestError {
  return &RestError {
    Message: message,
    Err: "internal server error",
    Code: http.StatusInternalServerError,
  }
}

func NewNotFoundError(message string) *RestError {
  return &RestError {
    Message: message,
    Err: "not found",
    Code: http.StatusNotFound,
  }
}
