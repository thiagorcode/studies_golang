package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thiagorcode/studies_golang/Part-2/model"
)

// 58:30
type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
	e.POST("/product", w.createProduct)

	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(c echo.Context) error {

	return c.JSON(http.StatusOK, w.Products)
}

func (w WebServer) createProduct(c echo.Context) error {
	product := model.NewProduct()
	// O Bind pega os dados que recebo do post
	// e gruda com a struct
	if err := c.Bind(product); err != nil {
		return err
	}
	w.Products.Add(*product)

	return c.JSON(http.StatusCreated, product)
}
