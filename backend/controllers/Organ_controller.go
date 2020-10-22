package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Pichais/app/ent"
	"github.com/Pichais/app/ent/organ"
	"github.com/gin-gonic/gin"
)

// OrganController defines the struct for the organ controller
type OrganController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateOrgan handles POST requests for adding organ entities
// @Summary Create organ
// @Description Create organ
// @ID create-organ
// @Accept   json
// @Produce  json
// @Param organ body ent.Organ true "Organ entity"
// @Success 200 {object} ent.Organ
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /organs [post]
func (ctl *OrganController) CreateOrgan(c *gin.Context) {
	obj := ent.Organ{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "organ binding failed",
		})
		return
	}

	u, err := ctl.client.Organ.
		Create().
		SetOrganName(obj.OrganName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetOrgan handles GET requests to retrieve a organ entity
// @Summary Get a organ entity by ID
// @Description get organ by ID
// @ID get-organ
// @Produce  json
// @Param id path int true "Organ ID"
// @Success 200 {object} ent.Organ
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /organs/{id} [get]
func (ctl *OrganController) GetOrgan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.Organ.
		Query().
		Where(organ.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListOrgan handles request to get a list of organ entities
// @Summary List organ entities
// @Description list organ entities
// @ID list-organ
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Organ
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /organs [get]
func (ctl *OrganController) ListOrgan(c *gin.Context) {
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

	organs, err := ctl.client.Organ.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, organs)
}

// DeleteOrgan handles DELETE requests to delete a organ entity
// @Summary Delete a organ entity by ID
// @Description get organ by ID
// @ID delete-organ
// @Produce  json
// @Param id path int true "Organ ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /organs/{id} [delete]
func (ctl *OrganController) DeleteOrgan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Organ.
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

// UpdateOrgan handles PUT requests to update a organ entity
// @Summary Update a organ entity by ID
// @Description update organ by ID
// @ID update-organ
// @Accept   json
// @Produce  json
// @Param id path int true "Organ ID"
// @Param organ body ent.Organ true "Organ entity"
// @Success 200 {object} ent.Organ
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /organs/{id} [put]
func (ctl *OrganController) UpdateOrgan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Organ{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "organ binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Organ.
		UpdateOneID(int(id)).
		SetOrganName(obj.OrganName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewOrganController creates and registers handles for the organ controller
func NewOrganController(router gin.IRouter, client *ent.Client) *OrganController {
	p := &OrganController{
		client: client,
		router: router,
	}
	p.register()
	return p
}

// InitOrganController registers routes to the main engine
func (ctl *OrganController) register() {
	organs := ctl.router.Group("/organs")

	organs.GET("", ctl.ListOrgan)

	// CRUD
	organs.POST("", ctl.CreateOrgan)
	organs.GET(":id", ctl.GetOrgan)
	organs.PUT(":id", ctl.UpdateOrgan)
	organs.DELETE(":id", ctl.DeleteOrgan)
}
