package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Pichais/app/ent"
	"github.com/Pichais/app/ent/organ"
	"github.com/Pichais/app/ent/physician"
	"github.com/Pichais/app/ent/typedisease"
	"github.com/gin-gonic/gin"
)

// SpacialityController defines the struct for the spaciality controller
type SpacialityController struct {
	client *ent.Client
	router gin.IRouter
}

// Spaciality defines the struct for the spaciality
type Spaciality struct {
	Organdata       int
	Physiciandata   int
	Typediseasedata int
}

// CreateSpaciality handles POST requests for adding spaciality entities
// @Summary Create spaciality
// @Description Create spaciality
// @ID create-spaciality
// @Accept   json
// @Produce  json
// @Param spaciality body Spaciality true "Spaciality entity"
// @Success 200 {object} Spaciality
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /spacialitys [post]
func (ctl *SpacialityController) CreateSpaciality(c *gin.Context) {
	obj := Spaciality{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "spaciality binding failed",
		})
		return
	}

	o, err := ctl.client.Organ.
		Query().
		Where(organ.IDEQ(int(obj.Organdata))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "organ not found",
		})
		return
	}

	p, err := ctl.client.Physician.
		Query().
		Where(physician.IDEQ(int(obj.Physiciandata))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "physiciandata diagnostic  not found",
		})
		return
	}

	t, err := ctl.client.TypeDisease.
		Query().
		Where(typedisease.IDEQ(int(obj.Typediseasedata))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	sa, err := ctl.client.Spaciality.
		Create().
		SetOgan(o).
		SetUser(p).
		SetType(t).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, sa)
}

// ListSpaciality handles request to get a list of spaciality entities
// @Summary List spaciality entities
// @Description list spaciality entities
// @ID list-spaciality
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Spaciality
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /spacialitys [get]
func (ctl *SpacialityController) ListSpaciality(c *gin.Context) {
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

	spacialitys, err := ctl.client.Spaciality.
		Query().
		WithUser().
		WithOgan().
		WithType().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, spacialitys)
}

// DeleteSpaciality handles DELETE requests to delete a spaciality entity
// @Summary Delete a spaciality entity by ID
// @Description get spaciality by ID
// @ID delete-spaciality
// @Produce  json
// @Param id path int true "Spaciality ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /spacialitys/{id} [delete]
func (ctl *SpacialityController) DeleteSpaciality(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Spaciality.
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

// NewSpacialityController creates and registers handles for the spaciality controller
func NewSpacialityController(router gin.IRouter, client *ent.Client) *SpacialityController {
	drc := &SpacialityController{
		client: client,
		router: router,
	}
	drc.register()
	return drc
}

// InitUserController registers routes to the main engine
func (ctl *SpacialityController) register() {
	spacialitys := ctl.router.Group("/spacialitys")

	spacialitys.GET("", ctl.ListSpaciality)

	// CRUD
	spacialitys.POST("", ctl.CreateSpaciality)
	spacialitys.DELETE(":id", ctl.DeleteSpaciality)
}
