package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserSegmentsRequest struct {
	AddSegments    []string `json:"add_segments"`
	RemoveSegments []string `json:"remove_segments"`
}

func (h *Handler) updateUserSegments(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid user id param")
		return
	}

	var userSegments UserSegmentsRequest
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

func (h *Handler) getUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	items, err := h.services.AvitoSegment.GetUserSegments(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
