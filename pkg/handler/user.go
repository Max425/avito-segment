package handler

import (
	"avito-segment/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// @Summary Update user segments
// @Tags users
// @Description Update segments for a user
// @Param user_id path int true "User ID"
// @Param input body models.UserSegmentsRequest true "Segments data"
// @Success 200 {object} statusResponse
// @Router /api/users/{user_id}/segments [post]
func (h *Handler) updateUserSegments(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid user id param")
		return
	}

	var userSegments models.UserSegmentsRequest
	if err := c.ShouldBindJSON(&userSegments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Services.AvitoUser.UpdateUserSegments(userID, userSegments.AddSegments, userSegments.RemoveSegments); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Get user segments
// @Tags users
// @Description Get segments of a user
// @Param user_id path int true "User ID"
// @Success 200 {array} string
// @Router /api/users/{user_id}/segments [get]
func (h *Handler) getUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	segments, err := h.Services.AvitoSegment.GetUserSegments(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Преобразование объектов сегментов в массив строк (slug'ов)
	var segmentSlugs []string
	for _, segment := range segments {
		segmentSlugs = append(segmentSlugs, segment.Slug)
	}

	c.JSON(http.StatusOK, segmentSlugs)
}

// addUserToSegmentWithTTL adds a user to a segment with a specified TTL (time to live).
// @Summary Add user to segment with TTL
// @Tags users
// @Description Add user to a segment with a specified TTL
// @Param user_id path int true "User ID"
// @Param request body models.UserToSegmentWithTTLRequest true "User to Segment with TTL request"
// @Success 200 {object} statusResponse
// @Router /api/users/{user_id}/segments/add_with_ttl [post]
func (h *Handler) addUserToSegmentWithTTL(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	var input models.UserToSegmentWithTTLRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	ttl := time.Duration(input.TTLMinutes) * time.Minute

	if err := h.Services.AvitoUser.AddUserToSegmentWithTTL(userId, input.SegmentSlug, ttl); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to add user to segment")
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// generateUserSegmentHistoryReport generates a user segment history report for a specified user within a given year and month.
// @Summary Generate user segment history report
// @Tags users
// @Description Generate a CSV report of user segment history for a specific user within a given year and month
// @Param user_id path int true "User ID"
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {file} csv "CSV report"
// @Router /api/users/{user_id}/segments/history [get]
func (h *Handler) generateUserSegmentHistoryReport(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	yearStr := c.Query("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid year query param")
		return
	}

	monthStr := c.Query("month")
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid month query param")
		return
	}

	reportData, err := h.Services.AvitoUser.GenerateUserSegmentHistoryReport(userID, year, time.Month(month))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=segment_history_report_%s_%d_%02d.csv", userIDStr, year, month))
	c.Data(http.StatusOK, "text/csv", reportData)
}
