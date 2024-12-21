package api

import (
	"github.com/axidex/CryptBot/internal/problems"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Des3
// @Summary Des3
// @Description Des3
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param l query string true "Left block" default(5)
// @Param r query string true "Right block" default(1)
// @Param k1 query string true "Key 1" default(3)
// @Param k2 query string true "Key 2" default(7)
// @Param k3 query string true "Key 3" default(5)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/des3 [post]
func (app *App) Des3(c *gin.Context) {

	params := &Des3{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution := problems.DES3(params.L, params.R, params.K1, params.K2, params.K3)

	c.String(http.StatusOK, solution)

	return
}
