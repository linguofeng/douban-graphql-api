package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/linguofeng/douban-graphql-api/helpers"
	_schema "github.com/linguofeng/douban-graphql-api/schema"
)

var isLambda = "false"
var isProxy = "false"
var schema = _schema.New()

func handlerGraphQL(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var result = []byte("")
	if strings.ToUpper(req.HTTPMethod) == http.MethodGet {
		query := req.QueryStringParameters["query"]
		operationName := req.QueryStringParameters["operationName"]
		variablesString := req.QueryStringParameters["variables"]
		variables := make(map[string]interface{})
		_ = json.Unmarshal([]byte(variablesString), &variables)
		// Browser render GraphiQL
		if strings.Contains(req.Headers["accept"], "text/html") {
			result = []byte(helpers.RenderGraphiQLHtml(graphql.Params{
				Schema:         schema,
				RequestString:  query,
				VariableValues: variables,
				OperationName:  operationName,
			}))
		} else {
			result, _ = json.Marshal(graphql.Do(graphql.Params{
				Schema:         schema,
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
			Schema:         schema,
			RequestString:  body.Query,
			VariableValues: body.Variables,
			OperationName:  body.OperationName,
		}))
	}
	return events.APIGatewayProxyResponse{
		Body:       string(result),
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json; charset=UTF-8",
		},
	}, nil
}

func handlerProxy(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://frodo.douban.com%s", req.Path))
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string(err.Error()),
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string(err.Error()),
			StatusCode: 500,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json; charset=UTF-8",
		},
	}, nil
}

func main() {
	// netlify functions
	if isLambda == "true" {
		if isProxy == "true" {
			lambda.Start(handlerGraphQL)
		} else {
			lambda.Start(handlerProxy)
		}
	} else {
		app := echo.New()
		app.Use(middleware.Recover())
		app.Use(middleware.Logger())
		app.Use(middleware.CORS())

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
					Schema:         schema,
					RequestString:  body.Query,
					VariableValues: body.Variables,
					OperationName:  body.OperationName,
				})
			} else {
				c.JSON(http.StatusOK, graphql.Do(graphql.Params{
					Schema:         schema,
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
				Schema:         schema,
				RequestString:  body.Query,
				VariableValues: body.Variables,
				OperationName:  body.OperationName,
			}))
			return nil
		})

		app.GET("/api/*", func(ctx echo.Context) error {
			resp, err := http.Get(fmt.Sprintf("https://frodo.douban.com%s", ctx.Request().URL.String()))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			ctx.JSONBlob(http.StatusOK, body)
			return nil
		})

		app.Start(":1234")
	}
}
