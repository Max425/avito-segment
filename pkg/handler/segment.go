package handler

import (
	avito_segment "avito-segment"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Create segment
// @Tags segments
// @Description Create a new segment
// @Accept json
// @Produce json
// @Param input body avito_segment.Segment true "Segment data"
// @Success 200 {object} avito_segment.Segment
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

// @Summary Delete segment
// @Tags segments
// @Description Delete a segment by slug
// @Param slug path string true "Segment slug"
// @Success 200 {object} statusResponse
// @Router /api/segments/{slug} [delete]
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
