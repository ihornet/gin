package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, dto{
		Id:   id,
		Name: "ihornet",
	})
}

func AddUser(ctx *gin.Context) {

	var body body

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}
	body.Id += 1

	log.Println("2323")

	ctx.JSON(http.StatusOK, body)

}
