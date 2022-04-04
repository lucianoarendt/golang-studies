package lists

import "github.com/gofiber/fiber/v2"

type ListService interface {
	Create(c *fiber.Ctx) error
	Read(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type listServiceImpl struct {
}

func NewListService() ListService {
	return &listServiceImpl{}
}

func (l *listServiceImpl) Create(c *fiber.Ctx) error {
	return nil
}

func (l *listServiceImpl) Read(c *fiber.Ctx) error {
	return nil
}

func (l *listServiceImpl) Update(c *fiber.Ctx) error {
	return nil
}

func (t *listServiceImpl) Delete(c *fiber.Ctx) error {
	return nil
}
