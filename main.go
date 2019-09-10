package main

import (
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/linguofeng/douban-graphql-api/helpers"
	"github.com/linguofeng/douban-graphql-api/schema"
)

var isLambda = "false"

func main() {
	app := echo.New()

	app.GET("/graphql", func(c echo.Context) error {
		body := new(struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		})
		if err := c.Bind(body); err != nil {
			return err
		}
		acceptHeader := c.Request().Header.Get("Accept")
		if strings.Contains(acceptHeader, "text/html") {
			helpers.RenderGraphiQL(c, graphql.Params{
				Schema:         schema.New(),
				RequestString:  body.Query,
				VariableValues: body.Variables,
				OperationName:  body.OperationName,
			})
		} else {
			c.JSON(http.StatusOK, graphql.Do(graphql.Params{
				Schema:         schema.New(),
				RequestString:  body.Query,
				VariableValues: body.Variables,
				OperationName:  body.OperationName,
			}))
		}
		return nil
	})

	app.POST("/graphql", func(c echo.Context) error {
		body := new(struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		})
		if err := c.Bind(body); err != nil {
			return err
		}
		c.JSON(http.StatusOK, graphql.Do(graphql.Params{
			Schema:         schema.New(),
			RequestString:  body.Query,
			VariableValues: body.Variables,
			OperationName:  body.OperationName,
		}))
		return nil
	})

	if isLambda == "true" {
		lambda.Start((&helpers.LambdaAdapter{Echo: app}).Handler)
	} else {
		app.Start("localhost:1234")
	}
}
