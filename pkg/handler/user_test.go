package handler

import (
	"avito-segment/pkg/service"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"avito-segment/models"
	service_mocks "avito-segment/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_updateUserSegments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mocks.NewMockAvitoUser(ctrl)
	s := &service.Service{AvitoUser: mockService}
	h := NewHandler(s)

	r := gin.Default()
	r.POST("/api/users/:user_id/segments", h.updateUserSegments)

	t.Run("Success", func(t *testing.T) {
		userID := 1
		input := models.UserSegmentsRequest{
			AddSegments:    []string{"segment1", "segment2"},
			RemoveSegments: []string{"segment3"},
		}

		mockService.EXPECT().UpdateUserSegments(userID, input.AddSegments, input.RemoveSegments).Return(nil)

		body := `{"add_segments":["segment1","segment2"],"remove_segments":["segment3"]}`
		req := httptest.NewRequest("POST", "/api/users/1/segments", strings.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse := `{"status":"ok"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("Error", func(t *testing.T) {
		userID := 1
		input := models.UserSegmentsRequest{
			AddSegments:    []string{"segment1", "segment2"},
			RemoveSegments: []string{"segment3"},
		}

		mockService.EXPECT().UpdateUserSegments(userID, input.AddSegments, input.RemoveSegments).Return(errors.New("update error"))

		body := `{"add_segments":["segment1","segment2"],"remove_segments":["segment3"]}`
		req := httptest.NewRequest("POST", "/api/users/1/segments", strings.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expectedResponse := `{"message":"update error"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestHandler_getUserSegments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mocks.NewMockAvitoSegment(ctrl)
	s := &service.Service{AvitoSegment: mockService}
	h := NewHandler(s)

	r := gin.Default()
	r.GET("/api/users/:user_id/segments", h.getUserSegments)

	t.Run("Success", func(t *testing.T) {
		userID := 1

		mockService.EXPECT().GetUserSegments(userID).Return([]models.Segment{
			{Slug: "segment1"},
			{Slug: "segment2"},
		}, nil)

		req := httptest.NewRequest("GET", "/api/users/1/segments", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse := `["segment1","segment2"]`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("Error", func(t *testing.T) {
		userID := 1

		mockService.EXPECT().GetUserSegments(userID).Return(nil, errors.New("get segments error"))

		req := httptest.NewRequest("GET", "/api/users/1/segments", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expectedResponse := `{"message":"get segments error"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestHandler_addUserToSegmentWithTTL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mocks.NewMockAvitoUser(ctrl)
	s := &service.Service{AvitoUser: mockService}
	h := NewHandler(s)

	r := gin.Default()
	r.POST("/api/users/:user_id/segments/add_with_ttl", h.addUserToSegmentWithTTL)

	t.Run("Success", func(t *testing.T) {
		userID := 1
		input := models.UserToSegmentWithTTLRequest{
			SegmentSlug: "segment1",
			TTLMinutes:  60,
		}

		mockService.EXPECT().AddUserToSegmentWithTTL(userID, input.SegmentSlug, gomock.Any()).Return(nil)

		body := `{"segment_slug":"segment1","ttl_minutes":60}`
		req := httptest.NewRequest("POST", "/api/users/1/segments/add_with_ttl", strings.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse := `{"status":"ok"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("InvalidUserID", func(t *testing.T) {
		body := `{"segment_slug":"segment1","ttl_minutes":60}`
		req := httptest.NewRequest("POST", "/api/users/invalid/segments/add_with_ttl", strings.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedResponse := `{"message":"invalid user id param"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("InvalidInput", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/users/1/segments/add_with_ttl", strings.NewReader("invalid json"))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedResponse := `{"message":"invalid input"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("AddUserToSegmentError", func(t *testing.T) {
		userID := 1
		input := models.UserToSegmentWithTTLRequest{
			SegmentSlug: "segment1",
			TTLMinutes:  60,
		}

		mockService.EXPECT().AddUserToSegmentWithTTL(userID, input.SegmentSlug, gomock.Any()).Return(errors.New("add user error"))

		body := `{"segment_slug":"segment1","ttl_minutes":60}`
		req := httptest.NewRequest("POST", "/api/users/1/segments/add_with_ttl", strings.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedResponse := `{"message":"Failed to add user to segment"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestHandler_generateUserSegmentHistoryReport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mocks.NewMockAvitoUser(ctrl)
	h := NewHandler(&service.Service{AvitoUser: mockService})

	r := gin.Default()
	r.GET("/api/users/:user_id/segments/history", h.generateUserSegmentHistoryReport)

	t.Run("Success", func(t *testing.T) {
		userID := 1
		year := 2023
		month := 8

		mockService.EXPECT().GenerateUserSegmentHistoryReport(userID, year, time.Month(month)).Return([]byte("csv data"), nil)

		req := httptest.NewRequest("GET", fmt.Sprintf("/api/users/%d/segments/history?year=%d&month=%d", userID, year, month), nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "attachment; filename=segment_history_report_1_2023_08.csv", w.Header().Get("Content-Disposition"))
		assert.Equal(t, "text/csv", w.Header().Get("Content-Type"))
		assert.Equal(t, "csv data", w.Body.String())
	})

	t.Run("InvalidUserID", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/invalid/segments/history?year=2023&month=8", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message":"invalid user id param"}`, w.Body.String())
	})

	t.Run("InvalidYear", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/1/segments/history?year=invalid&month=8", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message":"invalid year query param"}`, w.Body.String())
	})

	t.Run("InvalidMonth", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/1/segments/history?year=2023&month=invalid", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message":"invalid month query param"}`, w.Body.String())
	})

	t.Run("InternalServerError", func(t *testing.T) {
		userID := 1
		year := 2023
		month := 8

		mockService.EXPECT().GenerateUserSegmentHistoryReport(userID, year, time.Month(month)).Return(nil, errors.New("error"))

		req := httptest.NewRequest("GET", fmt.Sprintf("/api/users/%d/segments/history?year=%d&month=%d", userID, year, month), nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message":"error"}`, w.Body.String())
	})
}
