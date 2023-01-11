package handlers

import (
	"net/http"

	"02_products_api/internal/domain"
	"02_products_api/internal/product"

	"github.com/gin-gonic/gin"
	"time"
	"errors"
	"strconv"

)





type productHandler struct {
	server product.Service
}


func NewProductHandler(s product.Service) *productHandler {
	return &productHandler{
		server: s,
	}
}


// recibe peticiones del front 
// valida params





func (ph *productHandler) Create() gin.HandlerFunc {

	return func(c *gin.Context){

		var req domain.Product 

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		

		if err := verificarFecha(req.Expiration); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error al verificar la fecha ",
			})
			return
		}

		if err := verificarVacios(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error verificando vacios",
			})
			return
		}


		ph.server.Create(req)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Created",
			"data":    req,
		})

	}
	
}




func verificarFecha(date string) error {
	format := "02/01/2006"
	_, err := time.Parse(format, date)
	return err
}

func verificarVacios(prod domain.Product) error {
	if prod.Price == 0 {
		return errors.New("Price vacio")}
	if prod.Name == "" {
		return errors.New("Name vacio")
	}
	if prod.Expiration == "" {
		return errors.New("Expiration vacio")
	}
	if prod.CodeValue == "" {
		return errors.New("CodeValue  vacio")
	}
	if prod.Quantity == 0 {
		return errors.New("Quantity vacio")
	}
	return nil
}


//Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
func (ph *productHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		data:=ph.server.GetAll()
		c.JSON(http.StatusOK, gin.H{"message": "Listado de productos", "data": data})
	}
}


func (ph *productHandler) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
	// request
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "ERROR ", "data": nil})
			return
		}




		response, err_func:=ph.server.GetProductById(id)
		if err_func == nil {
			c.JSON(http.StatusOK, gin.H{"message": "PRODUCTO ENCONTRADO", "data": response})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "NO SE ENCONTRO EL PRODUCTO", "data": nil})
		}
	}
}



// Crear una ruta /products/search que nos permita buscar por parámetro los productos 
// cuyo precio sean mayor a un valor priceGt.

func (ph *productHandler) GetProductsPriceGt() gin.HandlerFunc {
	return func(c *gin.Context) {
		priceQuery, err := strconv.Atoi(c.Query("price"))

		if err != nil {
			c.JSON(404, gin.H{
				"message": "NO SE ENCONTRARON PRODUCTOS",
				"data":    nil,
			})
			return
		}

		response:=ph.server.GetProductsPriceGt(priceQuery)

		c.JSON(200, gin.H{
			"message": "PRODUCTOS ENCONTRADOS:",
			"data":    response,
		})
	}
}





//nuevo


func (ph *productHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req domain.Product 

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		

		if err := verificarFecha(req.Expiration); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error al verificar la fecha ",
			})
			return
		}

		if err := verificarVacios(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error verificando vacios",
			})
			return
		}


		response,err:=ph.server.Update(req.Id, req.Name,req.Quantity, req.CodeValue, req.IsPublished, req.Expiration, req.Price )

		if err!=nil{
			c.JSON(http.StatusOK, gin.H{
				"message": "Updated",
				"data":    response,
			})

		}else{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Not found",
				"data":    nil,
			})
		}
	}
	
	
}










