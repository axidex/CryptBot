package api

import (
	"github.com/axidex/CryptBot/internal/problems"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Des3
// @Summary Des3
// @Description Des3
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param l query int true "Left block" default(5)
// @Param r query int true "Right block" default(1)
// @Param k1 query int true "Key 1" default(3)
// @Param k2 query int true "Key 2" default(7)
// @Param k3 query int true "Key 3" default(5)
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

	solution := problems.DES3(*params.L, *params.R, *params.K1, *params.K2, *params.K3)

	c.String(http.StatusOK, solution)

	return
}

// A5
// @Summary A5
// @Description A5
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param a query string true "a" default(010011)
// @Param b query string true "b" default(1001011)
// @Param c query string true "c" default(10101011110)
// @Param text query string true "Text" default(f34ff541f04fed5ee04fff59fa50)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/a5 [post]
func (app *App) A5(c *gin.Context) {

	params := &A5{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution, err := problems.A5(params.A, params.B, params.C, params.Text)
	if err != nil {
		app.logger.Errorf("Error solve problem: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.String(http.StatusOK, solution)

	return
}

// Aes
// @Summary Aes
// @Description Aes
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param text query string true "text" default(Перу)
// @Param key query string true "key" default(Ключ)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/aes [post]
func (app *App) Aes(c *gin.Context) {

	params := &Aes{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution := problems.Aes(params.Text, params.Key)

	c.String(http.StatusOK, solution)

	return
}

// Blowfish
// @Summary Blowfish
// @Description Blowfish
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param l query int true "Left block" default(4)
// @Param r query int true "Right block" default(9)
// @Param k1 query int true "Key 1" default(1)
// @Param k2 query int true "Key 2" default(5)
// @Param k3 query int true "Key 3" default(3)
// @Param k4 query int true "Key 4" default(4)
// @Param k5 query int true "Key 5" default(5)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/blowfish [post]
func (app *App) Blowfish(c *gin.Context) {

	params := &Blowfish{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution := problems.BlowFish(*params.L, *params.R, *params.K1, *params.K2, *params.K3, *params.K4, *params.K5)

	c.String(http.StatusOK, solution)

	return
}

// Des
// @Summary Des
// @Description Des
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param text query string true "text" default(гана)
// @Param key query string true "key" default(куб)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/des [post]
func (app *App) Des(c *gin.Context) {

	params := &Des{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	textBinaryResult, err := problems.TransformText(params.Text)
	if err != nil {
		app.logger.Errorf("TransformText: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	keyBinaryResult, err := problems.TransformText(params.Key)
	if err != nil {
		app.logger.Errorf("TransformText: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hexText := problems.BinToHexStr(textBinaryResult.Value)
	hexKey := problems.BinToHexStr(keyBinaryResult.Value)

	resp, err := app.desClient.R().
		SetQueryParam("key", hexKey).
		SetQueryParam("data", hexText).
		Post("/encrypt")
	if err != nil {
		app.logger.Errorf("Des: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if resp.StatusCode() != 200 {
		app.logger.Errorf("Des: %v | %d", resp.Body(), resp.StatusCode())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	trimmedResp := strings.Trim(resp.String(), `"`)
	normalResp := strings.ReplaceAll(trimmedResp, `\n`, "\n")

	c.String(http.StatusOK, normalResp)

	return
}

// Diffie
// @Summary Diffie
// @Description Diffie
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param n query int true "n" default(17)
// @Param q query int true "q" default(11)
// @Param x query int true "x" default(2)
// @Param y query int true "y" default(3)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/diffie [post]
func (app *App) Diffie(c *gin.Context) {

	params := &Diffie{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution := problems.DiffieHellman(*params.N, *params.Q, *params.X, *params.Y)

	c.String(http.StatusOK, solution)

	return
}

// Elgamal
// @Summary Elgamal
// @Description Elgamal
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param p query int true "p" default(19)
// @Param g query int true "g" default(3)
// @Param k query int true "k" default(7)
// @Param x query int true "x" default(8)
// @Param m query int true "m" default(5)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/elgamal [post]
func (app *App) Elgamal(c *gin.Context) {

	params := &Elgamal{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution, err := problems.ElGamal(*params.P, *params.G, *params.K, *params.X, *params.M)
	if err != nil {
		app.logger.Errorf("Error solve problem: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.String(http.StatusOK, solution)

	return
}

// Enigma
// @Summary Enigma
// @Description Enigma
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param rot query int true "rotor" default(315)
// @Param ref query string true "reflector" default(C)
// @Param pp query string true "patch panel" default(A-U,D-Y)
// @Param str query string true "text" default(MJL)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/enigma [post]
func (app *App) Enigma(c *gin.Context) {

	params := &Enigma{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution, err := problems.Enigma(*params.Rot, params.Ref, params.Str, params.Pp)
	if err != nil {
		app.logger.Errorf("Error solve problem: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.String(http.StatusOK, solution)

	return
}

// Feistel
// @Summary Feistel
// @Description Feistel
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param data query string true "text" default(047b16c2)
// @Param keys query string true "key" default(67e99b3c)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/feistel [post]
func (app *App) Feistel(c *gin.Context) {

	params := &Feistel{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution, err := problems.Feistel(params.Data, params.Keys)
	if err != nil {
		app.logger.Errorf("Error solve problem: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.String(http.StatusOK, solution)

	return
}

// Hash
// @Summary Hash
// @Description Hash
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param word query string true "word" default(сотовыйтелефон)
// @Param h0 query string true "h0" default(7)
// @Param p query int true "p" default(11)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/hash [post]
func (app *App) Hash(c *gin.Context) {

	params := &Hash{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution, err := problems.Hash(params.Word, *params.H0, *params.P)
	if err != nil {
		app.logger.Errorf("Error solve problem: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.String(http.StatusOK, solution)

	return
}

// InvMix
// @Summary InvMix
// @Description InvMix
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param char query string true "char" default(h)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/invMix [post]
func (app *App) InvMix(c *gin.Context) {

	params := &InvMix{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invMixCol := problems.InvMixCol{}
	err = invMixCol.Solve(params.Char)
	if err != nil {
		app.logger.Errorf("Error solve problem: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.String(http.StatusOK, invMixCol.GetSolution())

	return
}

// Md5
// @Summary Md5
// @Description Md5
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param a query int true "a" default(5)
// @Param b query int true "b" default(1)
// @Param c query int true "c" default(9)
// @Param d query int true "d" default(1)
// @Param m query int true "m" default(6)
// @Param k query int true "k" default(10)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/md5 [post]
func (app *App) Md5(c *gin.Context) {

	params := &Md5{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution := problems.Md5(*params.A, *params.B, *params.C, *params.D, *params.M, *params.K)

	c.String(http.StatusOK, solution)

	return
}

// Rsa
// @Summary Rsa
// @Description Rsa
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param p query int true "p" default(3)
// @Param q query int true "q" default(7)
// @Param e query int true "e" default(5)
// @Param x query int true "x" default(101)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/rsa [post]
func (app *App) Rsa(c *gin.Context) {

	params := &Rsa{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution := problems.RSA(*params.P, *params.Q, *params.E, *params.X)

	c.String(http.StatusOK, solution)

	return
}

// SBlock
// @Summary SBlock
// @Description SBlock
// @Tags Crypt
// @Accept json
// @Produce plain/text
// @Param a query int true "a" default(11)
// @Param b query int true "b" default(17)
// @Param c query int true "c" default(15)
// @Param z0 query int true "z0" default(12)
// @Param n query int true "n" default(7)
// @Success 200 {string} string "Processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/sBlock [post]
func (app *App) SBlock(c *gin.Context) {

	params := &SBlock{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		app.logger.Errorf("Error parsing params: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solution := problems.SBlock(*params.A, *params.B, *params.C, *params.Z0, *params.N)

	c.String(http.StatusOK, solution)

	return
}
