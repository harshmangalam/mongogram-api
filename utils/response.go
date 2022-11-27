package utils

import "github.com/gofiber/fiber/v2"

type StatusText string

const (
	Success StatusText = "success"
	Error   StatusText = "error"
)

type ResponseSchema struct {
	Ctx        *fiber.Ctx `json:"-"`
	StatusText StatusText `json:"status"`
	StatusCode int        `json:"-"`
	Message    string     `json:"message"`
	Data       fiber.Map  `json:"data"`
}

func NewResponseSchema() *ResponseSchema {
	return &ResponseSchema{}
}

func (r *ResponseSchema) SetStatusText(statusText StatusText) *ResponseSchema {
	r.StatusText = statusText
	return r
}
func (r *ResponseSchema) SetStatusCode(statusCode int) *ResponseSchema {
	r.StatusCode = statusCode
	return r
}
func (r *ResponseSchema) SetMessage(message string) *ResponseSchema {
	r.Message = message
	return r
}

func (r *ResponseSchema) SetData(data fiber.Map) *ResponseSchema {
	r.Data = data
	return r
}

func CustomResponse(c *fiber.Ctx, statusText StatusText, statusCode int, message string, data any) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusText,
		"message": message,
		"data":    data,
	})
}

func InternalServerErrorResponse(c *fiber.Ctx, err error) {
	return CustomResponse(c, Error, fiber.StatusInternalServerError, "", nil)
}
