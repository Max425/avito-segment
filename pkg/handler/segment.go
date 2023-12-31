package handler

import (
	"avito-segment/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Create segment
// @Tags segments
// @Description Create a new segment
// @Accept json
// @Produce json
// @Param input body models.Segment true "Segment data"
// @Success 200 {object} models.Segment
// @Router /api/segments [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input models.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	slug, err := h.Services.AvitoSegment.Create(input)
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
	err := h.Services.AvitoSegment.Delete(c.Param("slug"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
