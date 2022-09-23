package services

import (
	"errors"

	"github.com/abhay676/school-managment/services/gatekeeper/app/dal"
	"github.com/abhay676/school-managment/services/gatekeeper/app/types"
	"github.com/abhay676/school-managment/services/gatekeeper/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateEntity(ctx *fiber.Ctx) error {
	b := new(types.CreateEntityDTO)
	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return ctx.JSON(err)
	}

	e := &dal.Entity{
		Name:     b.Name,
		Email:    b.Email,
		Password: b.Password,
		Role:     b.Role,
	}

	if err := dal.CreateEntity(e).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON((&types.EntityCreateResponse{
		Entity: &types.EntityResponse{
			EID:       e.EID,
			Name:      e.Name,
			Email:     e.Email,
			IsActive:  e.IsActive,
			CreatedAt: e.CreatedAt,
		},
	}))

}

func Login(ctx *fiber.Ctx) error {
	b := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return ctx.JSON(err)
	}

	e := &types.EntityResponse{}

	err := dal.FindUserByEmail(e, b.Email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.JSON(fiber.NewError(fiber.StatusUnauthorized, "Invalid email"))
	}

	if err := utils.CompareHashPassword(e.Password, b.Password); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.NewError(fiber.StatusUnauthorized, "Invalid Password"))
	}

	t := utils.GenerateJWT(&utils.Payload{
		Email:    e.Email,
		IsActive: e.IsActive,
		EID:      e.EID,
	})
	return ctx.JSON(&types.LoginResponse{
		Entity: e,
		Auth: &types.AccessResponse{
			Token: t,
		},
	})

}
