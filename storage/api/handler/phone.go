package handler

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetPhone(c *gin.Context) {
	id := c.Param("id")
	log.Println(id)
	phoneID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("bad request")
		return
	}
	resp, err := h.storage.Phone().GetPhone(phoneID)
	if err != nil {
		c.JSON(http.StatusNotFound, resp)
		return
	}
	n := rand.Intn(2)
	keys := [2]int{http.StatusOK, http.StatusInternalServerError}
	c.JSON(keys[n], resp)
}
