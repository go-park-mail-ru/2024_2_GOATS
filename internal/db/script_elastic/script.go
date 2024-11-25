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

func delIndex(indexName string) {
	url := fmt.Sprintf("http://localhost:9200/%s", indexName)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer([]byte("")))
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
	delIndex("movies")
	delIndex("actors")
	time.Sleep(2 * time.Second)

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
		{
			Id:        "14",
			Title:     "Перевозчик",
			CardURL:   "/static/movies/the-transporter/card.png",
			AlbumURL:  "/static/movies/the-transporter/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "15",
			Title:     "Трансформеры",
			CardURL:   "/static/movies/transformers/card.png",
			AlbumURL:  "/static/movies/transformers/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "16",
			Title:     "Волк с Уолл-Стрит",
			CardURL:   "/static/movies/wolf-of-wall-street/card.png",
			AlbumURL:  "/static/movies/wolf-of-wall-street/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "17",
			Title:     "Мстители",
			CardURL:   "/static/movies/avengers/card.png",
			AlbumURL:  "/static/movies/avengers/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "18",
			Title:     "Еще по одной",
			CardURL:   "/static/movies/drunk/card.png",
			AlbumURL:  "/static/movies/drunk/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "19",
			Title:     "Форд против Феррари",
			CardURL:   "/static/movies/ford-v-ferrari/card.png",
			AlbumURL:  "/static/movies/ford-v-ferrari/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "20",
			Title:     "Зеленая книга",
			CardURL:   "/static/movies/greenbook/card.png",
			AlbumURL:  "/static/movies/greenbook/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "21",
			Title:     "Однажды в Голливуде",
			CardURL:   "/static/movies/once-in-hollywood/card.png",
			AlbumURL:  "/static/movies/once-in-hollywood/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "22",
			Title:     "Ламборгини",
			CardURL:   "/static/movies/lamborgini/card.png",
			AlbumURL:  "/static/movies/lamborgini/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "23",
			Title:     "Легенда",
			CardURL:   "/static/movies/legend/card.png",
			AlbumURL:  "/static/movies/legend/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "24",
			Title:     "Пеле: Рождение легенды",
			CardURL:   "/static/movies/pele/card.png",
			AlbumURL:  "/static/movies/pele/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "25",
			Title:     "Стрельцов",
			CardURL:   "/static/movies/streltsov/card.png",
			AlbumURL:  "/static/movies/streltsov/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "26",
			Title:     "Гнев человеческий",
			CardURL:   "/static/movies/wrath-of-man/card.png",
			AlbumURL:  "/static/movies/wrath-of-man/poster.png",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			Id:        "27",
			Title:     "Брат 2",
			CardURL:   "/static/movies/brother2/card.png",
			AlbumURL:  "/static/movies/brother2/poster.png",
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
		{
			Id:          "11",
			FullName:    "Урсула Корберо Дельгадо",
			PhotoURL:    "/static/actors/corbero/small.png",
			PhotoBigURL: "/static/actors/corbero/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "12",
			FullName:    "Марион Котийяр",
			PhotoURL:    "/static/actors/alonso/small.png",
			PhotoBigURL: "/static/actors/cotillard/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "13",
			FullName:    "Леонардо ДиКаприо",
			PhotoURL:    "/static/actors/dicaprio/small.png",
			PhotoBigURL: "/static/actors/dicaprio/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "14",
			FullName:    "Фредерик Дифенталь",
			PhotoURL:    "/static/actors/diefenthal/small.png",
			PhotoBigURL: "/static/actors/diefenthal/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "15",
			FullName:    "Роберт Дауни-младший",
			PhotoURL:    "/static/actors/downey-jr/small.png",
			PhotoBigURL: "/static/actors/downey-jr/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "16",
			FullName:    "Кристофер Экклстон",
			PhotoURL:    "/static/actors/eccleston/small.png",
			PhotoBigURL: "/static/actors/eccleston/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "17",
			FullName:    "Джесси Айзенберг",
			PhotoURL:    "/static/actors/eisenberg/small.png",
			PhotoBigURL: "/static/actors/eisenberg/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			Id:          "18",
			FullName:    "Махершала Али",
			PhotoURL:    "/static/actors/mahershala-ali/small.png",
			PhotoBigURL: "/static/actors/mahershala-ali/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "19",
			FullName:    "Владимир Меньшов",
			PhotoURL:    "/static/actors/menshov/small.png",
			PhotoBigURL: "/static/actors/menshov/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "20",
			FullName:    "Вуди Харрельсон",
			PhotoURL:    "/static/actors/harrelson/small.png",
			PhotoBigURL: "/static/actors/harrelson/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "21",
			FullName:    "Фрэнк Грилло",
			PhotoURL:    "/static/actors/grillo/small.png",
			PhotoBigURL: "/static/actors/grillo/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "22",
			FullName:    "Бен Кингсли",
			PhotoURL:    "/static/actors/kingsley/small.png",
			PhotoBigURL: "/static/actors/kingsley/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "23",
			FullName:    "Сами Насери",
			PhotoURL:    "/static/actors/samy-nacery/small.png",
			PhotoBigURL: "/static/actors/samy-nacery/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "24",
			FullName:    "Кристоф Вальц",
			PhotoURL:    "/static/actors/waltz/small.png",
			PhotoBigURL: "/static/actors/waltz/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			Id:          "25",
			FullName:    "Сэм Уоррингтон",
			PhotoURL:    "/static/actors/sam-worthington/small.png",
			PhotoBigURL: "/static/actors/sam-worthington/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "26",
			FullName:    "Дмитрий Юрьевич Назаров",
			PhotoURL:    "/static/actors/nazarov/small.png",
			PhotoBigURL: "/static/actors/nazarov/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "27",
			FullName:    "Энн Хэтэуэй",
			PhotoURL:    "/static/actors/hathaway/small.png",
			PhotoBigURL: "/static/actors/hathaway/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "28",
			FullName:    "Меган Фокс",
			PhotoURL:    "/static/actors/fox/small.png",
			PhotoBigURL: "/static/actors/fox/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "29",
			FullName:    "Хайме Лоренте Лопес",
			PhotoURL:    "/static/actors/lorente/small.png",
			PhotoBigURL: "/static/actors/lorente/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "30",
			FullName:    "Данила Козловский",
			PhotoURL:    "/static/actors/kozlovsky/small.png",
			PhotoBigURL: "/static/actors/kozlovsky/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "31",
			FullName:    "Вигго Мортенсен",
			PhotoURL:    "/static/actors/mortensen/small.png",
			PhotoBigURL: "/static/actors/mortensen/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			Id:          "32",
			FullName:    "Стася Милославская",
			PhotoURL:    "/static/actors/miloslavskaya/small.png",
			PhotoBigURL: "/static/actors/miloslavskaya/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "33",
			FullName:    "Сергей Васильевич Маковецкий",
			PhotoURL:    "/static/actors/makovetskiy/small.png",
			PhotoBigURL: "/static/actors/makovetskiy/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "34",
			FullName:    "Омар Си",
			PhotoURL:    "/static/actors/omar-sy/small.png",
			PhotoBigURL: "/static/actors/omar-sy/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "35",
			FullName:    "Сеу Жоржи",
			PhotoURL:    "/static/actors/jorge/small.png",
			PhotoBigURL: "/static/actors/jorge/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "36",
			FullName:    "Александр Яценко",
			PhotoURL:    "/static/actors/yatsenko/small.png",
			PhotoBigURL: "/static/actors/yatsenko/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "37",
			FullName:    "Мэттью Макконахи",
			PhotoURL:    "/static/actors/mcconaughey/small.png",
			PhotoBigURL: "/static/actors/mcconaughey/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "38",
			FullName:    "Айла Фишер",
			PhotoURL:    "/static/actors/isla-fisher/small.png",
			PhotoBigURL: "/static/actors/isla-fisher/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "39",
			FullName:    "Том Харди",
			PhotoURL:    "/static/actors/tom-hardy/small.png",
			PhotoBigURL: "/static/actors/tom-hardy/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "40",
			FullName:    "Джейсон Стэйтем",
			PhotoURL:    "/static/actors/jason-statham/small.png",
			PhotoBigURL: "/static/actors/jason-statham/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			Id:          "41",
			FullName:    "Мишель Родригес",
			PhotoURL:    "/static/actors/michelle-rodriguez/small.png",
			PhotoBigURL: "/static/actors/michelle-rodriguez/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "42",
			FullName:    "Шая ЛаБаф",
			PhotoURL:    "/static/actors/labeouf/small.png",
			PhotoBigURL: "/static/actors/labeouf/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "43",
			FullName:    "Джона Хилл",
			PhotoURL:    "/static/actors/jonah-hill/small.png",
			PhotoBigURL: "/static/actors/jonah-hill/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "44",
			FullName:    "Марго Робби",
			PhotoURL:    "/static/actors/margo-robbie/small.png",
			PhotoBigURL: "/static/actors/margo-robbie/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "45",
			FullName:    "Мадс Миккельсен",
			PhotoURL:    "/static/actors/mikkelsen/small.png",
			PhotoBigURL: "/static/actors/mikkelsen/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "46",
			FullName:    "Дмитрий Геннадьевич Нагиев",
			PhotoURL:    "/static/actors/nagiev/small.png",
			PhotoBigURL: "/static/actors/nagiev/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "47",
			FullName:    "Елена Подкаминская",
			PhotoURL:    "/static/actors/podkaminskaya/small.png",
			PhotoBigURL: "/static/actors/podkaminskaya/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "48",
			FullName:    "Зои Салдана",
			PhotoURL:    "/static/actors/saldana/small.png",
			PhotoBigURL: "/static/actors/saldana/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "49",
			FullName:    "Олег Меньшиков",
			PhotoURL:    "/static/actors/menshikov/small.png",
			PhotoBigURL: "/static/actors/menshikov/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "50",
			FullName:    "Шу Ци",
			PhotoURL:    "/static/actors/shu/small.png",
			PhotoBigURL: "/static/actors/shu/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "51",
			FullName:    "Альваро Антонио Гарсия",
			PhotoURL:    "/static/actors/morte/small.png",
			PhotoBigURL: "/static/actors/morte/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "52",
			FullName:    "Эмили Браунинг",
			PhotoURL:    "/static/actors/emily-browning/small.png",
			PhotoBigURL: "/static/actors/emily-browning/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "53",
			FullName:    "Александр Наумов",
			PhotoURL:    "/static/actors/naumov/small.png",
			PhotoBigURL: "/static/actors/naumov/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "54",
			FullName:    "Мэтт Дэймон",
			PhotoURL:    "/static/actors/matt-damon/small.png",
			PhotoBigURL: "/static/actors/matt-damon/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "55",
			FullName:    "Александр Петров",
			PhotoURL:    "/static/actors/petrov/small.png",
			PhotoBigURL: "/static/actors/petrov/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "56",
			FullName:    "Холт МакКэллани",
			PhotoURL:    "/static/actors/mccallany/small.png",
			PhotoBigURL: "/static/actors/mccallany/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "57",
			FullName:    "Виктор Сухоруков",
			PhotoURL:    "/static/actors/suhorukov/small.png",
			PhotoBigURL: "/static/actors/suhorukov/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "58",
			FullName:    "Ли Джонджэ",
			PhotoURL:    "/static/actors/jae/small.png",
			PhotoBigURL: "/static/actors/jae/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "59",
			FullName:    "Уильям Брэдли Питт",
			PhotoURL:    "/static/actors/pitt/small.png",
			PhotoBigURL: "/static/actors/pitt/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "60",
			FullName:    "Пак Хэ-Су",
			PhotoURL:    "/static/actors/soo/small.png",
			PhotoBigURL: "/static/actors/soo/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "61",
			FullName:    "Чон Хо-ён",
			PhotoURL:    "/static/actors/yeon/small.png",
			PhotoBigURL: "/static/actors/yeon/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "62",
			FullName:    "Диана Крюгер",
			PhotoURL:    "/static/actors/kruger/small.png",
			PhotoBigURL: "/static/actors/kruger/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "63",
			FullName:    "Мигель Анхель Гарсия",
			PhotoURL:    "/static/actors/herran/small.png",
			PhotoBigURL: "/static/actors/herran/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "64",
			FullName:    "Виктор Викторович Хориняк",
			PhotoURL:    "/static/actors/horinyak/small.png",
			PhotoBigURL: "/static/actors/horinyak/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "65",
			FullName:    "Алессандро Нивола",
			PhotoURL:    "/static/actors/nivola/small.png",
			PhotoBigURL: "/static/actors/nivola/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			Id:          "66",
			FullName:    "Тайриз Дарнелл Гибсон",
			PhotoURL:    "/static/actors/gibson/small.png",
			PhotoBigURL: "/static/actors/gibson/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "67",
			FullName:    "Марк Руффало",
			PhotoURL:    "/static/actors/mark-ruffalo/small.png",
			PhotoBigURL: "/static/actors/mark-ruffalo/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			Id:          "68",
			FullName:    "Ольга Николаевна Кузьмина",
			PhotoURL:    "/static/actors/kuzmina/small.png",
			PhotoBigURL: "/static/actors/kuzmina/big.png",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			Id:          "69",
			FullName:    "Франсуа Клюзе",
			PhotoURL:    "/static/actors/kluzet/small.png",
			PhotoBigURL: "/static/actors/kluzet/big.png",
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
