package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tranduckhuy/go-ecommerce-backend-api/internal/services"
	"github.com/tranduckhuy/go-ecommerce-backend-api/pkg/responses"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (p *ProductController) GetByID(c *gin.Context) {
	id := c.Param("id")

	product := map[string]any{"id": id, "name": "Product 1", "price": 10.0}

	responses.Respond(c, responses.Success, product)
}

// Paginate products
func (p *ProductController) List(c *gin.Context) {
	// Get query parameters
	limit := c.Query("limit")
	page := c.Query("page")

	// Mock data
	products := []map[string]any{
		{"id": 1, "name": "Product 1", "price": 10.0},
		{"id": 2, "name": "Product 2", "price": 20.0},
	}

	// Create response object with data and pagination
	responseData := map[string]any{
		"items":       products,
		"total_items": len(products),
		"page":        page,
		"limit":       limit,
	}

	// Replace direct product response with structured response object
	responses.Respond(c, responses.Success, responseData)
}

// CreateProduct creates a new product
func (p *ProductController) Create(c *gin.Context) {
	responses.Respond(c, responses.Success, "Product created successfully")
}

// UpdateProduct updates an existing product
func (p *ProductController) Update(c *gin.Context) {
	id := c.Param("id")
	responses.Respond(c, responses.Success, "Product "+id+" updated successfully")
}

// DeleteProduct deletes a product
func (p *ProductController) Delete(c *gin.Context) {
	id := c.Param("id")
	responses.Respond(c, responses.Success, "Product "+id+" deleted successfully")
}

// LikeProduct likes a product
func (p *ProductController) Like(c *gin.Context) {
	id := c.Param("id")
	responses.Respond(c, responses.Success, "Product "+id+" liked successfully")
}

// AddReview adds a review to a product
func (p *ProductController) AddReview(c *gin.Context) {
	id := c.Param("id")
	review := c.PostForm("review")

	// Mock data
	product := map[string]any{"id": id, "review": review}

	responses.Respond(c, responses.Success, product)
}
