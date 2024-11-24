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
			Id:        "1",
			Title:     "Игра в кальмара",
			CardURL:   "/static/movies/squid-game/card.png",
			AlbumURL:  "/static/movies/squid-game/poster.png",
			Rating:    7.6,
			MovieType: "film",
			Country:   "Франция",
		},
		{
			Id:        "2",
			Title:     "Бумажный дом",
			CardURL:   "/static/movies/paper_house/card.png",
			AlbumURL:  "/static/movies/paper_house/poster.png",
			Rating:    8.2,
			MovieType: "film",
			Country:   "Франция",
		},
		{
			Id:        "3",
			Title:     "Кухня",
			CardURL:   "/static/movies/kitcnen/card.png",
			AlbumURL:  "/static/movies/kitcnen/poster.png",
			Rating:    8.2,
			MovieType: "film",
			Country:   "Франция",
		},
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
		{
			Id:        "6",
			Title:     "Иллюзия обмана",
			CardURL:   "/static/movies/how-you-see-me/card.png",
			AlbumURL:  "/static/movies/how-you-see-me/poster.png",
			Rating:    7.7,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "7",
			Title:     "Бесславные ублюдки",
			CardURL:   "/static/movies/inglourious-basterds/card.png",
			AlbumURL:  "/static/movies/inglourious-basterds/poster.png",
			Rating:    8.0,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "8",
			Title:     "Интерстеллар",
			CardURL:   "/static/movies/interstellar/card.png",
			AlbumURL:  "/static/movies/interstellar/poster.png",
			Rating:    8.6,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "9",
			Title:     "Легенда №17",
			CardURL:   "/static/movies/legend17/card.png",
			AlbumURL:  "/static/movies/legend17/poster.png",
			Rating:    8.0,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "10",
			Title:     "Человек, который изменил все",
			CardURL:   "/static/movies/moneyball/card.png",
			AlbumURL:  "/static/movies/moneyball/poster.png",
			Rating:    7.7,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "11",
			Title:     "Остров проклятых",
			CardURL:   "/static/movies/shutter-island/card.png",
			AlbumURL:  "/static/movies/shutter-island/poster.png",
			Rating:    8.5,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "12",
			Title:     "Такси 2",
			CardURL:   "/static/movies/taxi2/card.png",
			AlbumURL:  "/static/movies/taxi2/poster.png",
			Rating:    7.7,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "13",
			Title:     "Множественные святые Ньюагка",
			CardURL:   "/static/movies/the-many-saints-of-newark/card.png",
			AlbumURL:  "/static/movies/the-many-saints-of-newark/poster.png",
			Rating:    5.9,
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
		{
			Id:          "2",
			FullName:    "Анн Ле Ни",
			PhotoURL:    "/static/actors/alonso/small.png",
			PhotoBigURL: "/static/actors/ann-le-ni/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "3",
			FullName:    "Франсуа Берлеан",
			PhotoURL:    "/static/actors/berleand/big.png",
			PhotoBigURL: "/static/actors/berleand/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "4",
			FullName:    "Сергей Бодров",
			PhotoURL:    "/static/actors/bodrov/big.png",
			PhotoBigURL: "/static/actors/bodrov/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "5",
			FullName:    "Марк Богатырёв",
			PhotoURL:    "/static/actors/bogatyrev/big.png",
			PhotoBigURL: "/static/actors/bogatyrev/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "6",
			FullName:    "Мария Бонневи",
			PhotoURL:    "/static/actors/bonnevie/big.png",
			PhotoBigURL: "/static/actors/bonnevie/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "7",
			FullName:    "Линда Эдна Карделлини",
			PhotoURL:    "/static/actors/cardellini/big.png",
			PhotoBigURL: "/static/actors/cardellini/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			Id:          "8",
			FullName:    "Джессика Честейн",
			PhotoURL:    "/static/actors/alonso/small.png",
			PhotoBigURL: "/static/actors/chastain/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "9",
			FullName:    "Крис Эванс",
			PhotoURL:    "/static/actors/alonso/small.png",
			PhotoBigURL: "/static/actors/chris-evans/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "10",
			FullName:    "Кристиан Бэйл",
			PhotoURL:    "/static/actors/alonso/small.png",
			PhotoBigURL: "/static/actors/christian-bale/big.png",
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
