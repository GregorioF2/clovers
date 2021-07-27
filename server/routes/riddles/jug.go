package riddles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	riddlesController "github.com/gregorioF2/clovers/controllers/riddles"
	. "github.com/gregorioF2/clovers/lib/consts"
	. "github.com/gregorioF2/clovers/lib/errors"
)

func readAndValdidateJugRiddleQueryParams(queryParams map[string][]string) (int, int, int, error) {
	validateParam := func(key string) (int, error) {
		param, ok := queryParams[key]
		if !ok {
			return 0, &InvalidParametersError{Err: fmt.Sprintf("query param '%s' is required.", key)}
		}
		value, err := strconv.ParseInt(param[0], 10, 64)
		if err != nil {
			return 0, &InvalidParametersError{Err: fmt.Sprintf("query param '%s' must be an integer.", key)}
		}
		return int(value), nil
	}

	x, err := validateParam("x")
	if err != nil {
		return 0, 0, 0, err
	}
	y, err := validateParam("y")
	if err != nil {
		return 0, 0, 0, err
	}
	z, err := validateParam("z")
	if err != nil {
		return 0, 0, 0, err
	}
	return x, y, z, nil
}

func JugRiddleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	x, y, z, err := readAndValdidateJugRiddleQueryParams(r.URL.Query())
	if err != nil {
		var responseError *ResponseError
		switch e := err.(type) {
		case *InvalidParametersError:
			responseError = &ResponseError{Err: e.Error(), StatusCode: HttpStatusCode["ClientError"]["BadRequest"]}
		default:
			responseError = &ResponseError{Err: e.Error(), StatusCode: HttpStatusCode["ServerError"]["InternalServerError"]}
		}
		http.Error(w, responseError.Error(), responseError.StatusCode)
		return
	}

	data, err := riddlesController.JugRiddle(x, y, z)
	if err != nil {
		var responseError *ResponseError
		switch e := err.(type) {
		case *InvalidParametersError:
			responseError = &ResponseError{Err: e.Error(), StatusCode: HttpStatusCode["ClientError"]["BadRequest"]}
		default:
			responseError = &ResponseError{Err: e.Error(), StatusCode: HttpStatusCode["ServerError"]["InternalServerError"]}
		}
		http.Error(w, responseError.Error(), responseError.StatusCode)
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		responseError := &ResponseError{Err: "failed parse response to []byte", StatusCode: HttpStatusCode["ServerError"]["InternalServerError"]}
		http.Error(w, responseError.Error(), responseError.StatusCode)
		return
	}

	w.Write(res)
}
