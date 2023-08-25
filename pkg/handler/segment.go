package handler

import (
	avito_segment "avito-segment"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Create segment
// @Tags segments
// @Description create segment
// @ID create-segment
// @Accept  json
// @Produce  json
// @Param input body avito_segment.Segment true "segment slug"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/segments [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input avito_segment.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	slug, err := h.services.AvitoSegment.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"slug": slug,
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
