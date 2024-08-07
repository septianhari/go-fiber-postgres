package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/syndtr/goleveldb/leveldb/storage"
	"golang.org/x/mod/sumdb/storage"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := Book{}

	err := c.BodyParser(&book)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})

		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has ben added"})
	return nil
}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(&bookModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched successfully",
		"data":    bookModels,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_book", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books", r.GetBooksByID)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")

}
