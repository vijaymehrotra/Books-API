package controller

import (
	"vijaymehrotra/go-deploy/database"
	"vijaymehrotra/go-deploy/models"

	"github.com/gofiber/fiber/v2"
)

type BookData struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GetBooks(c *fiber.Ctx) error {
	bookModels := &[]models.Book{}

	err := database.DB.Find(bookModels).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error fetching books",
		})
		return err
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Books fetched successfully",
		"books":    bookModels,
	})
	return nil
}

func CreateBook(c *fiber.Ctx) error {
	var bookData BookData
	if err := c.BodyParser(&bookData); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "error parsing body",
		})
		return err
	}

	// Create a new book
	book := models.Book{
		Title:  bookData.Title,
		Author: bookData.Author,
	}

	result := database.DB.Create(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error creating book",
			"error":   result.Error.Error(),
		})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": book,
		"Message": "Book has been created",
	})
	return nil
}

func UpdateBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	bookData := BookData{}
	if err := c.BodyParser(&bookData); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "error parsing body",
		})
		return err
	}

	book := models.Book{}
	result := database.DB.Where("id = ?", bookId).First(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error fetching book",
			"error":   result.Error.Error(),
		})
	}

	book.Title = bookData.Title
	book.Author = bookData.Author

	result = database.DB.Save(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error updating book",
			"error":   result.Error.Error(),
		})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book has been updated",
		"book":    book,
	})
	return nil
}

func DeleteBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	book := models.Book{}
	result := database.DB.Where("id = ?", bookId).Delete(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error deleting book",
			"error":   result.Error.Error(),
		})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book has been deleted",
	})
	return nil
}