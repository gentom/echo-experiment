package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.Use(middleware.Logger())

	e.GET("/", hello)

	e.POST("/post", post)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func hello(c echo.Context) error {
	output := fmt.Sprintf("%#v", c.Request().Header)
	os.Stdout.Write([]byte(output + "\n"))
	return c.String(http.StatusOK, "ʕ◔ϖ◔ʔ")
}

type Post struct {
	name string `json:"name"`
}

func post(c echo.Context) (err error) {
	p := new(Post)
	if err = c.Bind(p); err != nil {
		return err
	}
	result := fmt.Sprintf("name:%s", p.name)
	os.Stdout.Write([]byte(result + "\n"))
	return c.JSON(http.StatusOK, p)
}
