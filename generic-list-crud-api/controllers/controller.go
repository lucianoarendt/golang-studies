package controllers

import "generic-list-crud-api/controllers/lists"

type Controller struct {
	List lists.ListService
}

func Start() *Controller {
	controllers := &Controller{
		List: lists.NewListService(),
	}
	return controllers
}
