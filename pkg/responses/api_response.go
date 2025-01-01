package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewAPIResponse(statusCode int, message string, data interface{}) *APIResponse {
	return &APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

// Ok sends a JSON response with HTTP status 200, a custom message, and additional data to the client.
func Ok(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, okResponse(message, data))
}

// okResponse constructs an APIResponse with a 200 OK status, a custom message, and optional data payload.
func okResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	}
}

// Created sends an HTTP 201 Created response with a message and optional data payload in JSON format.
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, createdResponse(message, data))
}

// createdResponse creates an APIResponse with HTTP 201 Created status, a message, and optional data payload.
func createdResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		StatusCode: http.StatusCreated,
		Message:    message,
		Data:       data,
	}
}

// NoContent sends a 204 No Content JSON response with the provided message to the client.
func NoContent(c *gin.Context, message string) {
	c.JSON(http.StatusNoContent, noContentResponse(message))
}

// noContentResponse creates an APIResponse with a 204 No Content status and the provided message.
func noContentResponse(message string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusNoContent,
		Message:    message,
	}
}

// BadRequest Bad Request: For validation failures or incorrect input
func BadRequest(c *gin.Context, message string, errors interface{}) {
	c.JSON(http.StatusBadRequest, badRequestResponse(message, errors))
}

// badRequestResponse creates an APIResponse with HTTP 400 Bad Request status, a message, and errors.
func badRequestResponse(message string, errors interface{}) APIResponse {
	return APIResponse{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Data:       errors, // Include validation errors or additional details
	}
}

// Unauthorized Unauthorized: For authentication failures or missing credentials
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, unauthorizedResponse(message))
}

// unauthorizedResponse creates an APIResponse with HTTP 401 Unauthorized status and a message.
func unauthorizedResponse(message string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	}
}

// Forbidden Forbidden: For access denied due to insufficient permissions
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, forbiddenResponse(message))
}

// forbiddenResponse creates an APIResponse with HTTP 403 Forbidden status and a message.
func forbiddenResponse(message string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusForbidden,
		Message:    message,
	}
}

// NotFound Not Found: For when a resource is not found
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, notFoundResponse(message))
}

// notFoundResponse creates an APIResponse with HTTP 404 Not Found status and a message.
func notFoundResponse(message string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

// InternalServerError Internal Server Error: For unexpected errors
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, internalServerErrorResponse(message))
}

// internalServerErrorResponse creates an APIResponse with HTTP 500 Internal Server Error status and a message.
func internalServerErrorResponse(message string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}
