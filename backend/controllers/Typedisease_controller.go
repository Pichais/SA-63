package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Pichais/app/ent"
	"github.com/Pichais/app/ent/typedisease"
	"github.com/gin-gonic/gin"
)

// TypeDiseaseController defines the struct for the typedisease controller
type TypeDiseaseController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTypeDisease handles POST requests for adding typedisease entities
// @Summary Create typedisease
// @Description Create typedisease
// @ID create-typedisease
// @Accept   json
// @Produce  json
// @Param typedisease body ent.TypeDisease true "TypeDisease entity"
// @Success 200 {object} ent.TypeDisease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typediseases [post]
func (ctl *TypeDiseaseController) CreateTypeDisease(c *gin.Context) {
	obj := ent.TypeDisease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "typedisease binding failed",
		})
		return
	}

	u, err := ctl.client.TypeDisease.
		Create().
		SetDiseaseName(obj.DiseaseName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetTypeDisease handles GET requests to retrieve a typedisease entity
// @Summary Get a typedisease entity by ID
// @Description get typedisease by ID
// @ID get-typedisease
// @Produce  json
// @Param id path int true "TypeDisease ID"
// @Success 200 {object} ent.TypeDisease
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typediseases/{id} [get]
func (ctl *TypeDiseaseController) GetTypeDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.TypeDisease.
		Query().
		Where(typedisease.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListTypeDisease handles request to get a list of typedisease entities
// @Summary List typedisease entities
// @Description list typedisease entities
// @ID list-typedisease
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.TypeDisease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typediseases [get]
func (ctl *TypeDiseaseController) ListTypeDisease(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	typediseases, err := ctl.client.TypeDisease.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, typediseases)
}

// DeleteTypeDisease handles DELETE requests to delete a typedisease entity
// @Summary Delete a typedisease entity by ID
// @Description get typedisease by ID
// @ID delete-typedisease
// @Produce  json
// @Param id path int true "TypeDisease ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typediseases/{id} [delete]
func (ctl *TypeDiseaseController) DeleteTypeDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.TypeDisease.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateTypeDisease handles PUT requests to update a typedisease entity
// @Summary Update a typedisease entity by ID
// @Description update typedisease by ID
// @ID update-typedisease
// @Accept   json
// @Produce  json
// @Param id path int true "TypeDisease ID"
// @Param typedisease body ent.TypeDisease true "TypeDisease entity"
// @Success 200 {object} ent.TypeDisease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typediseases/{id} [put]
func (ctl *TypeDiseaseController) UpdateTypeDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.TypeDisease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "typedisease binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.TypeDisease.
		UpdateOneID(int(id)).
		SetDiseaseName(obj.DiseaseName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewTypeDiseaseController creates and registers handles for the typedisease controller
func NewTypeDiseaseController(router gin.IRouter, client *ent.Client) *TypeDiseaseController {
	p := &TypeDiseaseController{
		client: client,
		router: router,
	}
	p.register()
	return p
}

// InitTypeDiseaseController registers routes to the main engine
func (ctl *TypeDiseaseController) register() {
	typediseases := ctl.router.Group("/typediseases")

	typediseases.GET("", ctl.ListTypeDisease)

	// CRUD
	typediseases.POST("", ctl.CreateTypeDisease)
	typediseases.GET(":id", ctl.GetTypeDisease)
	typediseases.PUT(":id", ctl.UpdateTypeDisease)
	typediseases.DELETE(":id", ctl.DeleteTypeDisease)
}
