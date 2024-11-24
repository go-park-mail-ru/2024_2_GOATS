package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Movie struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	CardURL     string  `json:"card_url"`
	AlbumURL    string  `json:"album_url"`
	Rating      float64 `json:"rating"`
	ReleaseDate string  `json:"release_date"`
	MovieType   string  `json:"movie_type"`
	Country     string  `json:"country"`
}

type Actor struct {
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	PhotoURL    string `json:"photo_url"`
	PhotoBigURL string `json:"photo_big_url"`
	Biography   string `json:"biography"`
	Country     string `json:"country"`
	BirthDate   string `json:"birth_date"`
}

func createIndex(indexName, mapping string) {
	url := fmt.Sprintf("http://localhost:9200/%s", indexName)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(mapping)))
	if err != nil {
		log.Fatalf("Error creating request for index %s: %v", indexName, err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request to create index %s: %v", indexName, err)
	}
	defer resp.Body.Close()

	body := make([]byte, 1024)
	resp.Body.Read(body)
	fmt.Printf("Index %s creation status: %s\n", indexName, resp.Status)
	fmt.Printf("Response body: %s\n", body)
}

func addMovie(id int, movie Movie) {
	data, err := json.Marshal(movie)
	if err != nil {
		log.Fatalf("Error marshaling movie data: %v", err)
	}

	url := fmt.Sprintf("http://localhost:9200/movies/_doc/%d", id)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body := make([]byte, 1024)
	resp.Body.Read(body)
	fmt.Printf("Movie with ID %d added with status: %s\n", id, resp.Status)
	fmt.Printf("Response body: %s\n", body)
}

func addActor(id int, actor Actor) {
	data, err := json.Marshal(actor)
	if err != nil {
		log.Fatalf("Error marshaling actor data: %v", err)
	}

	url := fmt.Sprintf("http://localhost:9200/actors/_doc/%d", id)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body := make([]byte, 1024)
	resp.Body.Read(body)
	fmt.Printf("Actor with ID %d added with status: %s\n", id, resp.Status)
	fmt.Printf("Response body: %s\n", body)
}

func main() {
	// Индекс для актеров
	actorMapping := `{
		"mappings": {
			"properties": {
				"full_name": { "type": "text" },
				"photo_url": { "type": "keyword" },
				"photo_big_url": { "type": "keyword" },
				"biography": { "type": "text" },
				"country": { "type": "keyword" },
				"birth_date": { "type": "date", "ignore_malformed": true }
			}
		}
	}`
	createIndex("actors", actorMapping)

	// Индекс для фильмов
	movieMapping := `{
		"mappings": {
			"properties": {
				"id": { "type": "integer" },
				"title": { "type": "text" },
				"card_url": { "type": "text" },
				"album_url": { "type": "text" },
				"rating": { "type": "float" },
				"release_date": { 
					"type": "date", 
					"format": "yyyy-MM-dd'T'HH:mm:ss'Z'", 
					"ignore_malformed": true 
				},
				"movie_type": { "type": "keyword" },
				"country": { "type": "text" }
			}
		}
	}`
	createIndex("movies", movieMapping)

	// Подождем, чтобы индексы были созданы
	time.Sleep(2 * time.Second)

	movies := []Movie{
		{
			Id:        "4",
			Title:     "1+1",
			CardURL:   "/static/movies/1+1/card.png",
			AlbumURL:  "/static/movies/1+1/poster.png",
			Rating:    8.8,
			MovieType: "film",
			Country:   "Франция",
		},
		{
			Id:        "5",
			Title:     "Аватар",
			CardURL:   "/static/movies/avatar/card.png",
			AlbumURL:  "/static/movies/avatar/poster.png",
			Rating:    8.0,
			MovieType: "film",
			Country:   "США",
		},
	}

	for i, movie := range movies {
		addMovie(i+1, movie)
	}

	actors := []Actor{
		{
			Id:          "1",
			FullName:    "Педро Гонсалес Алонсо",
			PhotoURL:    "/static/actors/alonso/small.png",
			PhotoBigURL: "/static/actors/alonso/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
	}

	for i, actor := range actors {
		addActor(i+1, actor)
	}

	time.Sleep(2 * time.Second)
}
