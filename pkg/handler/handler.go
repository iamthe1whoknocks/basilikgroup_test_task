package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	errParamNotProvided       = errors.New("query parameter is not provided")
	errQueryParamIsNotNumeric = errors.New("query parameter is not a number")
	errSecondQueryParamIsZero = errors.New("second query parameter is 0, cant divide by 0")
)

type Handler struct {
}

type Response struct {
	Success bool    `json:"success"`
	ErrCode string  `json:"error_code"`
	Value   float64 `json:"value"`
}

//NewHandler returns new handler instance
func New() *Handler {
	return &Handler{}
}

// parse numbers and checks, if a and b input is provided and if are they numbers
func parseNumbers(c echo.Context) (firstNum, secondNum int, err error) {
	params := c.Request().URL.Query()

	a := params["a"]
	b := params["b"]

	aStr := strings.Join(a, "")
	bStr := strings.Join(b, "")

	if aStr == "" || bStr == "" {
		return 0, 0, errParamNotProvided

	}

	firstNum, err = strconv.Atoi(aStr)
	if err != nil {
		return 0, 0, errQueryParamIsNotNumeric
	}

	secondNum, err = strconv.Atoi(bStr)
	if err != nil {
		return 0, 0, errQueryParamIsNotNumeric
	}

	return
}

//Add function
func (h *Handler) Add(c echo.Context) error {
	a, b, err := parseNumbers(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Value:   0,
			ErrCode: err.Error(),
		})
	}
	result := a + b

	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Value:   float64(result),
		ErrCode: "",
	})

}

//Sub function
func (h *Handler) Sub(c echo.Context) error {
	a, b, err := parseNumbers(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Value:   0,
			ErrCode: err.Error(),
		})
	}

	result := a - b

	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Value:   float64(result),
		ErrCode: "",
	})

}

//Mul function
func (h *Handler) Mul(c echo.Context) error {
	a, b, err := parseNumbers(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Value:   0,
			ErrCode: err.Error(),
		})
	}

	result := a * b

	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Value:   float64(result),
		ErrCode: "",
	})

}

//Divide function
func (h *Handler) Div(c echo.Context) error {
	a, b, err := parseNumbers(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Value:   0,
			ErrCode: err.Error(),
		})
	}

	if b == 0 {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Value:   0,
			ErrCode: errSecondQueryParamIsZero.Error(),
		})
	}

	result := float64(a) / float64(b)

	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Value:   result,
		ErrCode: "",
	})

}
