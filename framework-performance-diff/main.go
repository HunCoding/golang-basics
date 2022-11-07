package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
)

type ObjectExample struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	log.Println("Starting frameworks")
	ch := make(chan struct{})

	go initGinGonic()
	go initGoChi()
	go initGoFiber()
	go initGorillaMux()

	<-ch
}

func initGinGonic() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.POST("/testGinGonic/:channel", func(c *gin.Context) {

		var objectExample ObjectExample

		if err := c.ShouldBindJSON(&objectExample); err != nil {
			c.Status(400)
			return
		}

		if value, ok := c.Params.Get("channel"); !ok || value != "huncoding" {
			c.Status(400)
			return
		}

		c.JSON(http.StatusOK, objectExample)
		return
	})

	router.Run(":8080")
}

func initGorillaMux() {
	router := mux.NewRouter()
	router.HandleFunc("/testGorillaMux/{channel}", func(w http.ResponseWriter, r *http.Request) {
		var objectExample ObjectExample

		if value, ok := mux.Vars(r)["channel"]; !ok || value != "huncoding" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(b, &objectExample); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		jsonValue, err := json.Marshal(objectExample)
		w.Write(jsonValue)
		return

	}).Methods("POST")

	http.ListenAndServe(":8081", router)
}

func initGoChi() {
	r := chi.NewRouter()
	r.Post("/testGoChi/{channel}", func(w http.ResponseWriter, r *http.Request) {
		var objectExample ObjectExample

		channel := chi.URLParam(r, "channel")
		if channel != "huncoding" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(b, &objectExample); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		jsonValue, err := json.Marshal(objectExample)
		w.Write(jsonValue)
		return
	})

	http.ListenAndServe(":8082", r)
}

func initGoFiber() {
	app := fiber.New()
	log.SetOutput(ioutil.Discard)
	app.Post("/testGoFiber/:channel", func(c *fiber.Ctx) error {

		var objectExample ObjectExample

		if value := c.Params("channel"); value != "huncoding" {
			return c.Status(400).Send(nil)
		}

		if err := c.BodyParser(&objectExample); err != nil {
			return c.Status(400).Send(nil)
		}

		return c.Status(http.StatusOK).JSON(objectExample)
	})

	app.Listen(":8083")
}
