package handler

import (
	avito_segment "avito-segment"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createSegment(c *gin.Context) {
	var input avito_segment.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.AvitoSegment.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteSegment(c *gin.Context) {
	err := h.services.AvitoSegment.Delete(c.Param("slug"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
