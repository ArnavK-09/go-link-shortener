package main

// imports required
import (
	"fmt"
	"gitee.com/golang-module/dongle"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
	"strings"
)

// Constant Length for code generated for URL
const LENGTH int = 10

// Link struct for body parser
type Link struct {
	Url string `json:"url"`
}

// Function to decode code for url
func getLink(code string) string {
	// Decode by base58 from string and output string
	url := dongle.Decode.FromString(code).ByBase58().ToString()
	return url
}

// Function to encode url into code
func generateUrlCode(url string) string {
	// Encode by base58 from string and output string
	code := dongle.Encode.FromString(url).BySafeURL().ByBase58().ToString()
	return code
}

// Server Manager
func main() {
	// standard html templating engine
	engine := html.New("./pages", ".html")

	// new fiber client
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// GET / landing page
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Minimal Link Shortner Written In Golang...",
		})
	})

	// GET /link/<id> Link info
	app.Get("/:link", func(c *fiber.Ctx) error {
		var code string = c.Params("link")
		url := getLink(code)
		return c.Redirect(`//`+url, fiber.StatusPermanentRedirect)
	})

	// POST /link Add Link
	app.Post("/link", func(c *fiber.Ctx) error {
		// Parsing Request Body
		L := new(Link)
		if err := c.BodyParser(L); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Error Parsing Body!")
		}

		// Formatting url
		var url_fmt string = strings.TrimSpace(L.Url)

		// Validating URL
		const notValidUrl bool = false //todo

		if notValidUrl {
			// invalid url
			var response string = `<section>Invalid Link Provided<br /><kbd>Please provide a valid url like, https://github.com/ArnavK-09</kbd></section>`
			return c.SendString(response)
		} else {
			// Valid URL
			var code = generateUrlCode(url_fmt)
			var response string = fmt.Sprintf(`<section>Successfully Link Created!<br />Follow Link Here:- <a target="_blank" href="/%s">%s</a></section>`, code, code)
			return c.SendString(response)
		}
	})

	// starting server
	log.Fatal(app.Listen(":3000"))
}
