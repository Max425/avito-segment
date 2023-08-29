package handler

import (
	"avito-segment/pkg/service"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"avito-segment/models"
	service_mocks "avito-segment/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createSegment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mocks.NewMockAvitoSegment(ctrl)
	s := &service.Service{AvitoSegment: mockService}
	h := NewHandler(s)

	r := gin.Default()
	r.POST("/api/segments", h.createSegment)

	t.Run("Success", func(t *testing.T) {
		expectedSlug := "test-slug"
		input := models.Segment{Slug: "test-slug"}

		mockService.EXPECT().Create(input).Return(expectedSlug, nil)

		body := `{"slug":"test-slug"}`
		req := httptest.NewRequest("POST", "/api/segments", strings.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse := `{"slug":"test-slug"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("Error", func(t *testing.T) {
		input := models.Segment{Slug: "test-slug"}

		mockService.EXPECT().Create(input).Return("", errors.New("something went wrong"))

		body := `{"slug":"test-slug"}`
		req := httptest.NewRequest("POST", "/api/segments", strings.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expectedResponse := `{"message":"something went wrong"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("Invalid JSON Format", func(t *testing.T) {
		body := `invalid_json`
		req := httptest.NewRequest("POST", "/api/segments", strings.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Segment Creation Error", func(t *testing.T) {
		input := models.Segment{Slug: "test-slug"}
		mockService.EXPECT().Create(input).Return("", errors.New("segment creation error"))

		body := `{"slug":"test-slug"}`
		req := httptest.NewRequest("POST", "/api/segments", strings.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expectedResponse := `{"message":"segment creation error"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestHandler_deleteSegment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mocks.NewMockAvitoSegment(ctrl)
	s := &service.Service{AvitoSegment: mockService}
	h := NewHandler(s)

	r := gin.Default()
	r.DELETE("/api/segments/:slug", h.deleteSegment)

	t.Run("Success", func(t *testing.T) {
		mockService.EXPECT().Delete("test-slug").Return(nil)

		req := httptest.NewRequest("DELETE", "/api/segments/test-slug", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse := `{"status":"ok"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("Error", func(t *testing.T) {
		mockService.EXPECT().Delete("test-slug").Return(errors.New("something went wrong"))

		req := httptest.NewRequest("DELETE", "/api/segments/test-slug", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expectedResponse := `{"message":"something went wrong"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}
