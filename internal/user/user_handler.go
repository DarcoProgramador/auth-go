package user

import (
	"strconv"
	"time"

	"github.com/DarcoProgramador/auth-go/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Register(c *fiber.Ctx) error {

	var u UserCreate

	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": util.ErrBadRequest.Error(),
		})
	}

	if err := h.Service.Create(&u); err != nil {
		if err == util.ErrUserExists {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": util.ErrUserExists.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": util.ErrCreateUser.Error(),
		})
	}

	userResponse := &UserResponse{
		Username: u.Username,
		Email:    u.Email,
	}
	return c.Status(fiber.StatusCreated).JSON(userResponse)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var u UserLogin

	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": util.ErrBadRequest.Error(),
		})
	}

	token, err := h.Service.Login(&u)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": util.ErrLogin.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "AccessToken",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login success",
	})
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "AcessToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logout success",
	})
}

func (h *Handler) Me(c *fiber.Ctx) error {
	cookie := c.Cookies("AcessToken")

	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": util.ErrUnauthorized.Error(),
		})
	}

	token, err := jwt.ParseWithClaims(cookie, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": util.ErrTokenInternalServerError.Error(),
		})
	}

	claims := token.Claims.(*UserClaim)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		return err
	}

	result, err := h.Service.FindByID(uint(issuer))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": util.ErrNotFoundUser.Error(),
		})
	}

	u := &UserResponse{
		Username: result.Username,
		Email:    result.Email,
	}

	return c.Status(fiber.StatusOK).JSON(u)
}
