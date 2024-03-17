// Package handlers provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get All Customers
	// (GET /customers)
	GetCustomers(c *gin.Context, params GetCustomersParams)
	// Create New Customer
	// (POST /customers)
	PostCustomers(c *gin.Context)
	// User Login
	// (POST /login)
	PostLogin(c *gin.Context)
	// Get Products
	// (GET /products)
	GetProducts(c *gin.Context, params GetProductsParams)
	// Create New Sales Invoice
	// (POST /sales-invoices)
	PostSalesInvoices(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetCustomers operation middleware
func (siw *ServerInterfaceWrapper) GetCustomers(c *gin.Context) {

	var err error

	c.Set(BearerTokenScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCustomersParams

	// ------------- Optional query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, false, "search", c.Request.URL.Query(), &params.Search)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter search: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCustomers(c, params)
}

// PostCustomers operation middleware
func (siw *ServerInterfaceWrapper) PostCustomers(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostCustomers(c)
}

// PostLogin operation middleware
func (siw *ServerInterfaceWrapper) PostLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostLogin(c)
}

// GetProducts operation middleware
func (siw *ServerInterfaceWrapper) GetProducts(c *gin.Context) {

	var err error

	c.Set(BearerTokenScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProductsParams

	// ------------- Optional query parameter "group" -------------

	err = runtime.BindQueryParameter("form", true, false, "group", c.Request.URL.Query(), &params.Group)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter group: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProducts(c, params)
}

// PostSalesInvoices operation middleware
func (siw *ServerInterfaceWrapper) PostSalesInvoices(c *gin.Context) {

	c.Set(BearerTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostSalesInvoices(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/customers", wrapper.GetCustomers)
	router.POST(options.BaseURL+"/customers", wrapper.PostCustomers)
	router.POST(options.BaseURL+"/login", wrapper.PostLogin)
	router.GET(options.BaseURL+"/products", wrapper.GetProducts)
	router.POST(options.BaseURL+"/sales-invoices", wrapper.PostSalesInvoices)
}