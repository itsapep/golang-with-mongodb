package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-with-mongodb/delivery/api"
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/model/dto"
	"github.com/itsapep/golang-with-mongodb/usecase"
	"github.com/itsapep/golang-with-mongodb/utils"
)

type ProductController struct {
	router *gin.Engine
	api.BaseApi
	prodRegUc     usecase.ProductRegistrationUsecase
	prodFindAllUc usecase.FindAllProductUsecase
	prodUpdUc     usecase.UpdateProductUsecase
	prodDelUc     usecase.DeleteProductUsecase
	prodGetIdUc   usecase.GetProductByIdUsecase
	prodGetCatUc  usecase.GetProductByCategoryUsecase
}

func (pc *ProductController) registerNewProduct(ctx *gin.Context) {
	var newProduct model.Product
	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = pc.prodRegUc.Register(&newProduct)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
	})
}

func (pc *ProductController) findAllProduct(ctx *gin.Context) {
	var pageData dto.Paging
	err := ctx.ShouldBindJSON(&pageData)
	if err != nil {
		log.Println(err.Error())
		return
	}
	convertPage, _ := strconv.Atoi(pageData.Page)
	convertLimit, _ := strconv.Atoi(pageData.Limit)
	products, err := pc.prodFindAllUc.FindAllProduct(int64(convertPage), int64(convertLimit))
	if err != nil {
		pc.Failed(ctx, err)
		return
	}
	pc.Success(ctx, &products)
}

func (pc *ProductController) updateProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedProduct model.Product
	err := pc.ParseRequestBody(ctx, &updatedProduct)
	if err != nil {
		pc.Failed(ctx, utils.RequiredError())
		return
	}
	err = pc.prodUpdUc.UpdateProductById(id, &updatedProduct)
	if err != nil {
		pc.Failed(ctx, err)
		return
	}
	pc.Success(ctx, updatedProduct)
}

func (pc *ProductController) deleteProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	err := pc.prodDelUc.DeleteProductById(id)
	if err != nil {
		pc.Failed(ctx, err)
		return
	}
	pc.Success(ctx, id)
}

func (pc *ProductController) getProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := pc.prodGetIdUc.GetProductById(id)

	if err != nil {
		pc.Failed(ctx, err)
		return
	}
	pc.Success(ctx, product)
}

func (pc *ProductController) getProductByCategory(ctx *gin.Context) {
	var products []model.Product
	category := ctx.Param("category")
	products, err := pc.prodGetCatUc.GetProductByCategory(category)
	if err != nil {
		pc.Failed(ctx, err)
		return
	}
	pc.Success(ctx, products)
}

func NewProductController(router *gin.Engine,
	prodRegUc usecase.ProductRegistrationUsecase,
	prodFindAllUc usecase.FindAllProductUsecase,
	prodUpdUc usecase.UpdateProductUsecase,
	prodDelUc usecase.DeleteProductUsecase,
	prodGetIdUc usecase.GetProductByIdUsecase,
	prodGetCatUc usecase.GetProductByCategoryUsecase) *ProductController {
	controller := ProductController{
		router:    router,
		prodRegUc: prodRegUc,
	}
	router.POST("/product", controller.registerNewProduct)
	router.GET("/product/all", controller.findAllProduct)
	router.PUT("/product/id/:id", controller.updateProductById)
	router.DELETE("/product/id/:id", controller.deleteProductById)
	router.GET("/product/id/:id", controller.getProductById)
	router.GET("/product/category/:category", controller.getProductByCategory)
	return &controller
}
