package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type StatusText string

const (
	StatusSuccess StatusText = "success"
	StatusError   StatusText = "error"
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

func (r *ResponseSchema) SetCtx(ctx *fiber.Ctx) *ResponseSchema {
	r.Ctx = ctx
	return r
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

func (r *ResponseSchema) Return() error {
	return r.Ctx.Status(r.StatusCode).JSON(fiber.Map{
		"status":  r.StatusText,
		"message": r.Message,
		"data":    r.Data,
	})
}

func CustomResponse(c *fiber.Ctx, statusCode int, statusText StatusText, message string, data fiber.Map) error {
	return NewResponseSchema().
		SetCtx(c).
		SetStatusCode(statusCode).
		SetStatusText(statusText).
		SetMessage(message).
		SetData(data).
		Return()

}

func InternalServerErrorResponse(c *fiber.Ctx, err error) error {
	log.Println(err)
	return CustomResponse(c, fiber.StatusInternalServerError, StatusError, err.Error(), nil)
}

func BadRequestErrorResponse(c *fiber.Ctx, err error, data fiber.Map) error {
	log.Println(err)
	return CustomResponse(c, fiber.StatusBadRequest, StatusError, err.Error(), data)
}

func NotFoundErrorResponse(c *fiber.Ctx) error {
	return CustomResponse(c, fiber.StatusNotFound, StatusError, "Not found", nil)
}

func OkResponse(c *fiber.Ctx, message string, data fiber.Map) error {
	return CustomResponse(c, fiber.StatusOK, StatusSuccess, message, data)
}

func CreatedResponse(c *fiber.Ctx, message string, data fiber.Map) error {
	return CustomResponse(c, fiber.StatusCreated, StatusSuccess, message, data)
}
