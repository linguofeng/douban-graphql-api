package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/linguofeng/douban-graphql-api/helpers"
	"github.com/linguofeng/douban-graphql-api/schema"
)

var isLambda = "false"

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	acceptHeader := req.Headers["Accept"]
	query := req.QueryStringParameters["query"]
	var result = []byte("")
	if query != "" {
		operationName := req.QueryStringParameters["operationName"]
		variablesString := req.QueryStringParameters["variables"]
		variables := make(map[string]interface{})
		_ = json.Unmarshal([]byte(variablesString), &variables)
		if strings.Contains(acceptHeader, "text/html") {
			result = []byte("graphiql")
		} else {
			result, _ = json.Marshal(graphql.Do(graphql.Params{
				Schema:         schema.New(),
				RequestString:  query,
				VariableValues: variables,
				OperationName:  operationName,
			}))
		}
	} else {
		body := new(struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		})
		_ = json.Unmarshal([]byte(req.Body), &body)
		result, _ = json.Marshal(graphql.Do(graphql.Params{
			Schema:         schema.New(),
			RequestString:  body.Query,
			VariableValues: body.Variables,
			OperationName:  body.OperationName,
		}))
	}
	return events.APIGatewayProxyResponse{
		Body:       string(result),
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
			"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		},
	}, nil
}

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
		lambda.Start(handler)
	} else {
		app.Start("localhost:1234")
	}
}
