package main

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"text/template"
)

func main() {
	type Menu struct {
		Name string
		Icon string
	}
	type User struct {
		Name  string
		Menus []Menu
	}
	menus := []Menu{
		Menu{
			Name: "Profile",
			Icon: "<svg aria-hidden=\"true\" class=\"mr-2 w-5 h-5 text-gray-400 group-hover:text-gray-500 dark:text-gray-500 dark:group-hover:text-gray-300\" fill=\"currentColor\" viewBox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" d=\"M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-6-3a2 2 0 11-4 0 2 2 0 014 0zm-2 4a5 5 0 00-4.546 2.916A5.986 5.986 0 0010 16a5.986 5.986 0 004.546-2.084A5 5 0 0010 11z\" clip-rule=\"evenodd\"></path></svg>",
		},
		Menu{
			Name: "Dashboard",
			Icon: "<svg aria-hidden=\"true\" class=\"mr-2 w-5 h-5 text-blue-600 dark:text-blue-500\" fill=\"currentColor\" viewBox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M5 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2H5zM5 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2H5zM11 5a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V5zM11 13a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z\"></path></svg>",
		},
		Menu{
			Name: "Settings",
			Icon: "<svg aria-hidden=\"true\" class=\"mr-2 w-5 h-5 text-gray-400 group-hover:text-gray-500 dark:text-gray-500 dark:group-hover:text-gray-300\" fill=\"currentColor\" viewBox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M5 4a1 1 0 00-2 0v7.268a2 2 0 000 3.464V16a1 1 0 102 0v-1.268a2 2 0 000-3.464V4zM11 4a1 1 0 10-2 0v1.268a2 2 0 000 3.464V16a1 1 0 102 0V8.732a2 2 0 000-3.464V4zM16 3a1 1 0 011 1v7.268a2 2 0 010 3.464V16a1 1 0 11-2 0v-1.268a2 2 0 010-3.464V4a1 1 0 011-1z\"></path></svg>",
		},
	}
	user := User{
		Name:  "Firman Agam",
		Menus: menus,
	}
	app := fiber.New()
	// for logger in cmd
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		t, err := template.ParseFiles("./app/pages/pages.html")
		if err != nil {
			panic(err)
		}
		buf := new(bytes.Buffer)
		if err = t.Execute(buf, user); err != nil {
			panic(err)
		}
		return c.SendString(buf.String())
	})

	log.Fatal(app.Listen(":9900"))
}
