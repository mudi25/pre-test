package controller

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mudi25/pretest/entity"
	"github.com/mudi25/pretest/errors"
)

var language []*entity.Language = []*entity.Language{}
var validate *validator.Validate = validator.New()

func NewController() *echo.Echo {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.HideBanner = true
	e.GET("/", GetHello)
	e.GET("/palindrom", GetPalindrom)
	e.GET("/languages", GetLanguages)
	language := e.Group("/language")
	language.POST("", CreateLanguage)
	language.GET("/:id", GetLanguageById)
	language.GET("", GetLanguages)
	language.PATCH("/:id", PatchLanguage)
	language.DELETE("/:id", DeleteLanguage)
	return e
}
func GetHello(ctx echo.Context) error {
	return response(ctx, "Hello Go Developers", nil)
}
func GetLanguage(ctx echo.Context) error {
	return response(ctx, entity.Language{
		Language:       "C",
		Appeared:       1972,
		Created:        []string{"Dennis Ritchie"},
		Functional:     true,
		ObjectOriented: false,
		Relation: &entity.Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	}, nil)
}
func GetPalindrom(ctx echo.Context) error {
	req := struct {
		Text string `json:"text" validate:"required"`
	}{}
	if err := validateRequest(ctx, &req); err != nil {
		return response(ctx, nil, errors.BadRequest())
	}
	req.Text = "Not Palindrome"
	if isPalindrom := entity.IsPalindrom(req.Text); isPalindrom {
		req.Text = "Palindrome"
	}
	return response(ctx, req, nil)
}
func CreateLanguage(ctx echo.Context) error {
	var req entity.Language
	if err := validateRequest(ctx, &req); err != nil {
		return response(ctx, nil, errors.BadRequest())
	}
	language = append(language, &req)
	return response(ctx, req, nil)
}
func GetLanguages(ctx echo.Context) error {
	return response(ctx, language, nil)
}

func GetLanguageById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return response(ctx, nil, errors.BadRequest())
	}
	if id+1 > len(language) || len(language) == 0 {
		return response(ctx, nil, errors.IdNotFound())
	}
	return response(ctx, language[id], nil)
}
func PatchLanguage(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return response(ctx, nil, errors.BadRequest())
	}

	var req entity.Language
	if err := validateRequest(ctx, &req); err != nil {
		return response(ctx, nil, errors.BadRequest())
	}
	if id+1 > len(language) || len(language) == 0 {
		return response(ctx, nil, errors.IdNotFound())
	}
	language[id] = &req
	return response(ctx, language[id], nil)
}
func DeleteLanguage(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return response(ctx, nil, errors.BadRequest())
	}
	if id+1 > len(language) || len(language) == 0 {
		return response(ctx, nil, errors.IdNotFound())
	}
	l := language[id]
	language = append(language[:id], language[id+1:]...)
	return response(ctx, l, nil)
}
