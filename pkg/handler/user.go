package handler

import (
	"avito-segment/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	if err := h.services.AvitoUser.UpdateUserSegments(userID, userSegments.AddSegments, userSegments.RemoveSegments); err != nil {
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

	segments, err := h.services.AvitoSegment.GetUserSegments(userId)
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
