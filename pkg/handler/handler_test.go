package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	testIP = "127.0.0.1"
)

func TestAddHandler(t *testing.T) {
	e := echo.New()

	handler := New()

	positiveTestCases := []struct {
		a      int
		b      int
		result float64
		code   int
	}{
		{2, 2, 4, http.StatusOK},
		{2, 0, 2, http.StatusOK},
		{4, -2, 2, http.StatusOK},
	}

	for _, tc := range positiveTestCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/add?a=%d&b=%d", tc.a, tc.b), nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = handler.Add(c)

		resp := &Response{}
		err := json.NewDecoder(rec.Body).Decode(resp)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.result, resp.Value)
		assert.Equal(t, tc.code, rec.Code)
	}

	negativeTestCases := []struct {
		a        int
		b        int
		code     int
		errValue string
	}{
		{a: 2, code: http.StatusBadRequest, errValue: "query parameter is not provided"},
		{a: 4, code: http.StatusBadRequest, errValue: "query parameter is not provided"},
	}
	for _, tc := range negativeTestCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/add?a=%d", tc.a), nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = handler.Add(c)

		resp := &Response{}
		err := json.NewDecoder(rec.Body).Decode(resp)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.errValue, resp.ErrCode)
		assert.Equal(t, tc.code, rec.Code)
	}

}

func TestSubHandler(t *testing.T) {
	e := echo.New()

	handler := New()

	positiveTestCases := []struct {
		a      int
		b      int
		result float64
		code   int
	}{
		{4, 2, 2, http.StatusOK},
		{2, 0, 2, http.StatusOK},
		{2, -4, 6, http.StatusOK},
	}

	for _, tc := range positiveTestCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/sub?a=%d&b=%d", tc.a, tc.b), nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = handler.Sub(c)

		resp := &Response{}
		err := json.NewDecoder(rec.Body).Decode(resp)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.result, resp.Value)
		assert.Equal(t, tc.code, rec.Code)
	}

	negativeTestCases := []struct {
		a        int
		b        int
		code     int
		errValue string
	}{
		{a: 2, code: http.StatusBadRequest, errValue: "query parameter is not provided"},
		{a: 4, code: http.StatusBadRequest, errValue: "query parameter is not provided"},
	}
	for _, tc := range negativeTestCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/sub?a=%d", tc.a), nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = handler.Sub(c)

		resp := &Response{}
		err := json.NewDecoder(rec.Body).Decode(resp)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.errValue, resp.ErrCode)
		assert.Equal(t, tc.code, rec.Code)
	}

}

func TestMulHandler(t *testing.T) {
	e := echo.New()

	handler := New()

	positiveTestCases := []struct {
		a      int
		b      int
		result float64
		code   int
	}{
		{4, 2, 8, http.StatusOK},
		{2, 0, 0, http.StatusOK},
		{2, -4, -8, http.StatusOK},
	}

	for _, tc := range positiveTestCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/mul?a=%d&b=%d", tc.a, tc.b), nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = handler.Mul(c)

		resp := &Response{}
		err := json.NewDecoder(rec.Body).Decode(resp)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.result, resp.Value)
		assert.Equal(t, tc.code, rec.Code)
	}

	negativeTestCases := []struct {
		a        int
		b        int
		code     int
		errValue string
	}{
		{a: 2, code: http.StatusBadRequest, errValue: "query parameter is not provided"},
		{a: 4, code: http.StatusBadRequest, errValue: "query parameter is not provided"},
	}
	for _, tc := range negativeTestCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/mul?a=%d", tc.a), nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = handler.Mul(c)

		resp := &Response{}
		err := json.NewDecoder(rec.Body).Decode(resp)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.errValue, resp.ErrCode)
		assert.Equal(t, tc.code, rec.Code)
	}

}

func TestDivHandler(t *testing.T) {
	e := echo.New()

	handler := New()

	positiveTestCases := []struct {
		a      int
		b      int
		result float64
		code   int
	}{
		{4, 2, 2, http.StatusOK},
		{9, 3, 3, http.StatusOK},
		{2, 2, 1, http.StatusOK},
	}

	for _, tc := range positiveTestCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/div?a=%d&b=%d", tc.a, tc.b), nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = handler.Div(c)

		resp := &Response{}
		err := json.NewDecoder(rec.Body).Decode(resp)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.result, resp.Value)
		assert.Equal(t, tc.code, rec.Code)
	}

	negativeTestCases := []struct {
		a        int
		b        int
		code     int
		errValue string
	}{
		{a: 5, b: 0, code: http.StatusBadRequest, errValue: "second query parameter is 0, cant divide by 0"},
	}
	for _, tc := range negativeTestCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/div?a=%d&b=%d", tc.a, tc.b), nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = handler.Div(c)

		resp := &Response{}
		err := json.NewDecoder(rec.Body).Decode(resp)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.errValue, resp.ErrCode)
		assert.Equal(t, tc.code, rec.Code)
	}

}
