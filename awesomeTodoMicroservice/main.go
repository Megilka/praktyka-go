package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var tasks []task

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.GET("/", hello) // <-- endpoint
	e.GET("/create", create_task)
	e.GET("/tasks", get_tasks)
	e.GET("/tasks/:id", get_task)
	e.GET("/update/:id", update_task)
	e.GET("/delete/:id", delete_task)

	e.Logger.Fatal(e.Start(":1323"))
}

func create_task(c echo.Context) error {
	title := c.QueryParam("title")
	desc := c.QueryParam("description")

	res := task{title, desc, false}

	tasks = append(tasks, res)

	return c.JSON(http.StatusOK, res)
}

func get_tasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func get_task(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	return c.JSON(http.StatusOK, tasks[id])
}

func update_task(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	title := c.QueryParam("title")
	desc := c.QueryParam("description")
	done := c.QueryParam("done")

	if title != "" {
		tasks[id].Title = title
	}

	if desc != "" {
		tasks[id].Description = desc
	}

	if done == "true" {
		tasks[id].Done = true
	} else if done == "false" {
		tasks[id].Done = false
	}

	return c.JSON(http.StatusOK, tasks[id])
}

func delete_task(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	tasks = append(tasks[:id], tasks[id+1:]...)

	return c.JSON(http.StatusOK, tasks)
}
