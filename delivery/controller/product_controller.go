package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/usecase"
)

type ProductController struct {
	router         *gin.Engine
	productUsecase usecase.ProductRegistrationUsecase
}

func (pc *ProductController) registerNewProduct(ctx *gin.Context) {
	var newProduct model.Product
	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = pc.productUsecase.Register(&newProduct)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
	})
}

func NewProductController(router *gin.Engine, productUsecase usecase.ProductRegistrationUsecase) *ProductController {
	controller := ProductController{
		router:         router,
		productUsecase: productUsecase,
	}
	router.POST("/product", controller.registerNewProduct)
	return &controller
}
