package repository

import (
	"net/http"

	"github.com/JuanPabloSerna/jps-alert-crud/database/migrations"
	"github.com/JuanPabloSerna/jps-alert-crud/database/models"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(alert models.Alert) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(alert)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (r *Repository) CreateAlert(context *fiber.Ctx) error {
	alert := models.Alert{}
	err := context.BodyParser(&alert)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}
	errors := ValidateStruct(alert)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := r.DB.Create(&alert).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't create alert", "data": err})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Alert has been added", "data": alert})
	return nil
}
func (r *Repository) UpdateAlert(context *fiber.Ctx) error {
	alert := models.Alert{}
	err := context.BodyParser(&alert)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}
	errors := ValidateStruct(alert)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := r.DB
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}
	if db.Model(&alert).Where("id = ?", id).Updates(&alert).RowsAffected == 0 {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get Alert with given id"})
	}

	return context.JSON(fiber.Map{"status": "success", "message": "Alert successfully updated"})
}

func (r *Repository) DeleteAlert(context *fiber.Ctx) error {
	alertModel := migrations.Alerts{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := r.DB.Delete(alertModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete boo"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Alert delete successfully"})
	return nil
}
func (r *Repository) GetAlerts(context *fiber.Ctx) error {
	db := r.DB
	model := db.Model(&migrations.Alerts{})

	pg := paginate.New(&paginate.Config{
		DefaultSize:        20,
		CustomParamEnabled: true,
	})

	page := pg.With(model).Request(context.Request()).Response(&[]migrations.Alerts{})

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"data": page,
	})
	return nil
}

func (r *Repository) GetAlertByID(context *fiber.Ctx) error {
	id := context.Params("id")
	alertModel := &migrations.Alerts{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(alertModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get the alert"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Alert id fetched successfully", "data": alertModel})
	return nil
}
