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

func readAndValdidateJugRiddleQueryParams(queryParams map[string][]string) (int, int, int, *Error) {
	validateParam := func(key string) (int, *Error) {
		param, ok := queryParams[key]
		if !ok {
			return 0, NewError(fmt.Sprintf("Error: query param '%s' is required.", key), HttpStatusCode["ClientError"]["BadRequest"])
		}
		value, err := strconv.ParseInt(param[0], 10, 64)
		if err != nil {
			return 0, NewError(fmt.Sprintf("Error: query param '%s' must be an integer.", key), HttpStatusCode["ClientError"]["BadRequest"])
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

// public

func JugRiddleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	x, y, z, responseError := readAndValdidateJugRiddleQueryParams(r.URL.Query())
	if responseError != nil {
		http.Error(w, responseError.Err, responseError.StatusCode)
		return
	}

	data, responseError := riddlesController.JugRiddle(x, y, z)
	if responseError != nil {
		http.Error(w, responseError.Err, responseError.StatusCode)
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		responseError = NewError("Error parsin response :: "+err.Error(), HttpStatusCode["ServerError"]["InternalServerError"])
		http.Error(w, responseError.Err, responseError.StatusCode)
		return
	}

	w.Write(res)
}
