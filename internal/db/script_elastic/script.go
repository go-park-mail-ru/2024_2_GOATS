package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/db/script_elastic/models"
)

func delIndex(indexName string) {
	url := fmt.Sprintf("http://83.166.232.3:9200/%s", indexName)
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
	defer func() {
		if clErr := resp.Body.Close(); clErr != nil {
			fmt.Printf("cannot close delIndex body: %v", clErr)
		}
	}()

	body := make([]byte, 1024)
	_, err = resp.Body.Read(body)

	if err != nil {
		fmt.Printf("error deleting index: %v", err)
		return
	}
	fmt.Printf("Index %s creation status: %s\n", indexName, resp.Status)
	fmt.Printf("Response body: %s\n", body)
}

func createIndex(indexName, mapping string) {
	url := fmt.Sprintf("http://83.166.232.3:9200/%s", indexName)
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
	defer func() {
		if clErr := resp.Body.Close(); clErr != nil {
			fmt.Printf("cannot close createIndex body: %v", clErr)
		}
	}()

	body := make([]byte, 1024)
	_, err = resp.Body.Read(body)

	if err != nil {
		fmt.Printf("error creating index: %v", err)
		return
	}

	fmt.Printf("Index %s creation status: %s\n", indexName, resp.Status)
	fmt.Printf("Response body: %s\n", body)
}

func addMovie(id int, movie models.Movie) {
	data, err := movie.MarshalJSON()
	if err != nil {
		log.Fatalf("Error marshaling movie_service data: %v", err)
	}

	url := fmt.Sprintf("http://83.166.232.3:9200/movies/_doc/%d", id)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}

	defer func() {
		if clErr := resp.Body.Close(); clErr != nil {
			fmt.Printf("cannot close addMovie body: %v", clErr)
		}
	}()

	body := make([]byte, 1024)
	_, err = resp.Body.Read(body)

	if err != nil {
		fmt.Println(fmt.Errorf("error adding movies: %w", err))
		return
	}
	fmt.Printf("Movie with ID %d added with status: %s\n", id, resp.Status)
	fmt.Printf("Response body: %s\n", body)
}

func addActor(id int, actor models.Actor) {
	data, err := actor.MarshalJSON()
	if err != nil {
		log.Fatalf("Error marshaling actor data: %v", err)
	}

	url := fmt.Sprintf("http://83.166.232.3:9200/actors/_doc/%d", id)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}

	defer func() {
		if clErr := resp.Body.Close(); clErr != nil {
			fmt.Printf("cannot close addActor body: %v", clErr)
		}
	}()

	body := make([]byte, 1024)
	_, err = resp.Body.Read(body)

	if err != nil {
		fmt.Println(fmt.Errorf("error adding actors: %w", err))
		return
	}
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

	movies := []models.Movie{
		{
			ID:        "1",
			Title:     "Игра в кальмара",
			CardURL:   "/static/movies_all/squid-game/card.webp",
			AlbumURL:  "/static/movies_all/squid-game/poster.webp",
			Rating:    7.6,
			MovieType: "film",
			Country:   "Франция",
		},
		{
			ID:        "2",
			Title:     "Бумажный дом",
			CardURL:   "/static/movies_all/paper_house/card.webp",
			AlbumURL:  "/static/movies_all/paper_house/poster.webp",
			Rating:    8.2,
			MovieType: "film",
			Country:   "Франция",
		},
		{
			ID:        "3",
			Title:     "Кухня",
			CardURL:   "/static/movies_all/kitcnen/card.webp",
			AlbumURL:  "/static/movies_all/kitcnen/poster.webp",
			Rating:    8.2,
			MovieType: "film",
			Country:   "Франция",
		},
		{
			ID:        "4",
			Title:     "1+1",
			CardURL:   "/static/movies_all/1_plus_1/card.webp",
			AlbumURL:  "/static/movies_all/1_plus_1/poster.webp",
			Rating:    8.8,
			MovieType: "film",
			Country:   "Франция",
		},
		{
			ID:        "5",
			Title:     "Аватар",
			CardURL:   "/static/movies_all/avatar/card.webp",
			AlbumURL:  "/static/movies_all/avatar/poster.webp",
			Rating:    8.0,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "6",
			Title:     "Иллюзия обмана",
			CardURL:   "/static/movies_all/how-you-see-me/card.webp",
			AlbumURL:  "/static/movies_all/how-you-see-me/poster.webp",
			Rating:    7.7,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "7",
			Title:     "Бесславные ублюдки",
			CardURL:   "/static/movies_all/inglourious-basterds/card.webp",
			AlbumURL:  "/static/movies_all/inglourious-basterds/poster.webp",
			Rating:    8.0,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "8",
			Title:     "Интерстеллар",
			CardURL:   "/static/movies_all/interstellar/card.webp",
			AlbumURL:  "/static/movies_all/interstellar/poster.webp",
			Rating:    8.6,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "9",
			Title:     "Легенда №17",
			CardURL:   "/static/movies_all/legend17/card.webp",
			AlbumURL:  "/static/movies_all/legend17/poster.webp",
			Rating:    8.0,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "10",
			Title:     "Человек, который изменил все",
			CardURL:   "/static/movies_all/moneyball/card.webp",
			AlbumURL:  "/static/movies_all/moneyball/poster.webp",
			Rating:    7.7,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "11",
			Title:     "Остров проклятых",
			CardURL:   "/static/movies_all/shutter-island/card.webp",
			AlbumURL:  "/static/movies_all/shutter-island/poster.webp",
			Rating:    8.5,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "12",
			Title:     "Такси 2",
			CardURL:   "/static/movies_all/taxi2/card.webp",
			AlbumURL:  "/static/movies_all/taxi2/poster.webp",
			Rating:    7.7,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "13",
			Title:     "Множественные святые Ньюагка",
			CardURL:   "/static/movies_all/the-many-saints-of-newark/card.webp",
			AlbumURL:  "/static/movies_all/the-many-saints-of-newark/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "14",
			Title:     "Перевозчик",
			CardURL:   "/static/movies_all/the-transporter/card.webp",
			AlbumURL:  "/static/movies_all/the-transporter/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "15",
			Title:     "Трансформеры",
			CardURL:   "/static/movies_all/transformers/card.webp",
			AlbumURL:  "/static/movies_all/transformers/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "16",
			Title:     "Волк с Уолл-Стрит",
			CardURL:   "/static/movies_all/wolf-of-wall-street/card.webp",
			AlbumURL:  "/static/movies_all/wolf-of-wall-street/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "17",
			Title:     "Мстители",
			CardURL:   "/static/movies_all/avengers/card.webp",
			AlbumURL:  "/static/movies_all/avengers/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "18",
			Title:     "Еще по одной",
			CardURL:   "/static/movies_all/drunk/card.webp",
			AlbumURL:  "/static/movies_all/drunk/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "19",
			Title:     "Форд против Феррари",
			CardURL:   "/static/movies_all/ford-v-ferrari/card.webp",
			AlbumURL:  "/static/movies_all/ford-v-ferrari/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "20",
			Title:     "Зеленая книга",
			CardURL:   "/static/movies_all/greenbook/card.webp",
			AlbumURL:  "/static/movies_all/greenbook/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "21",
			Title:     "Однажды в Голливуде",
			CardURL:   "/static/movies_all/once-in-hollywood/card.webp",
			AlbumURL:  "/static/movies_all/once-in-hollywood/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "22",
			Title:     "Ламборгини",
			CardURL:   "/static/movies_all/lamborgini/card.webp",
			AlbumURL:  "/static/movies_all/lamborgini/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "23",
			Title:     "Легенда",
			CardURL:   "/static/movies_all/legend/card.webp",
			AlbumURL:  "/static/movies_all/legend/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "24",
			Title:     "Пеле: Рождение легенды",
			CardURL:   "/static/movies_all/pele/card.webp",
			AlbumURL:  "/static/movies_all/pele/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "25",
			Title:     "Стрельцов",
			CardURL:   "/static/movies_all/streltsov/card.webp",
			AlbumURL:  "/static/movies_all/streltsov/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "26",
			Title:     "Гнев человеческий",
			CardURL:   "/static/movies_all/wrath-of-man/card.webp",
			AlbumURL:  "/static/movies_all/wrath-of-man/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
		{
			ID:        "27",
			Title:     "Брат 2",
			CardURL:   "/static/movies_all/brother2/card.webp",
			AlbumURL:  "/static/movies_all/brother2/poster.webp",
			Rating:    5.9,
			MovieType: "film",
			Country:   "США",
		},
	}

	for i, movie := range movies {
		addMovie(i+1, movie)
	}

	actors := []models.Actor{
		{
			ID:          "1",
			FullName:    "Педро Гонсалес Алонсо",
			PhotoURL:    "/static/actors/alonso/small.webp",
			PhotoBigURL: "/static/actors/alonso/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "2",
			FullName:    "Анн Ле Ни",
			PhotoURL:    "/static/actors/alonso/small.webp",
			PhotoBigURL: "/static/actors/ann-le-ni/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "3",
			FullName:    "Франсуа Берлеан",
			PhotoURL:    "/static/actors/berleand/big.webp",
			PhotoBigURL: "/static/actors/berleand/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "4",
			FullName:    "Сергей Бодров",
			PhotoURL:    "/static/actors/bodrov/big.webp",
			PhotoBigURL: "/static/actors/bodrov/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "5",
			FullName:    "Марк Богатырёв",
			PhotoURL:    "/static/actors/bogatyrev/big.webp",
			PhotoBigURL: "/static/actors/bogatyrev/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "6",
			FullName:    "Мария Бонневи",
			PhotoURL:    "/static/actors/bonnevie/big.webp",
			PhotoBigURL: "/static/actors/bonnevie/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "7",
			FullName:    "Линда Эдна Карделлини",
			PhotoURL:    "/static/actors/cardellini/big.webp",
			PhotoBigURL: "/static/actors/cardellini/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			ID:          "8",
			FullName:    "Джессика Честейн",
			PhotoURL:    "/static/actors/alonso/small.webp",
			PhotoBigURL: "/static/actors/chastain/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "9",
			FullName:    "Крис Эванс",
			PhotoURL:    "/static/actors/alonso/small.webp",
			PhotoBigURL: "/static/actors/chris-evans/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "10",
			FullName:    "Кристиан Бэйл",
			PhotoURL:    "/static/actors/alonso/small.webp",
			PhotoBigURL: "/static/actors/christian-bale/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "11",
			FullName:    "Урсула Корберо Дельгадо",
			PhotoURL:    "/static/actors/corbero/small.webp",
			PhotoBigURL: "/static/actors/corbero/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "12",
			FullName:    "Марион Котийяр",
			PhotoURL:    "/static/actors/alonso/small.webp",
			PhotoBigURL: "/static/actors/cotillard/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "13",
			FullName:    "Леонардо ДиКаприо",
			PhotoURL:    "/static/actors/dicaprio/small.webp",
			PhotoBigURL: "/static/actors/dicaprio/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "14",
			FullName:    "Фредерик Дифенталь",
			PhotoURL:    "/static/actors/diefenthal/small.webp",
			PhotoBigURL: "/static/actors/diefenthal/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "15",
			FullName:    "Роберт Дауни-младший",
			PhotoURL:    "/static/actors/downey-jr/small.webp",
			PhotoBigURL: "/static/actors/downey-jr/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "16",
			FullName:    "Кристофер Экклстон",
			PhotoURL:    "/static/actors/eccleston/small.webp",
			PhotoBigURL: "/static/actors/eccleston/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "17",
			FullName:    "Джесси Айзенберг",
			PhotoURL:    "/static/actors/eisenberg/small.webp",
			PhotoBigURL: "/static/actors/eisenberg/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			ID:          "18",
			FullName:    "Махершала Али",
			PhotoURL:    "/static/actors/mahershala-ali/small.webp",
			PhotoBigURL: "/static/actors/mahershala-ali/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "19",
			FullName:    "Владимир Меньшов",
			PhotoURL:    "/static/actors/menshov/small.webp",
			PhotoBigURL: "/static/actors/menshov/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "20",
			FullName:    "Вуди Харрельсон",
			PhotoURL:    "/static/actors/harrelson/small.webp",
			PhotoBigURL: "/static/actors/harrelson/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "21",
			FullName:    "Фрэнк Грилло",
			PhotoURL:    "/static/actors/grillo/small.webp",
			PhotoBigURL: "/static/actors/grillo/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "22",
			FullName:    "Бен Кингсли",
			PhotoURL:    "/static/actors/kingsley/small.webp",
			PhotoBigURL: "/static/actors/kingsley/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "23",
			FullName:    "Сами Насери",
			PhotoURL:    "/static/actors/samy-nacery/small.webp",
			PhotoBigURL: "/static/actors/samy-nacery/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "24",
			FullName:    "Кристоф Вальц",
			PhotoURL:    "/static/actors/waltz/small.webp",
			PhotoBigURL: "/static/actors/waltz/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			ID:          "25",
			FullName:    "Сэм Уоррингтон",
			PhotoURL:    "/static/actors/sam-worthington/small.webp",
			PhotoBigURL: "/static/actors/sam-worthington/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "26",
			FullName:    "Дмитрий Юрьевич Назаров",
			PhotoURL:    "/static/actors/nazarov/small.webp",
			PhotoBigURL: "/static/actors/nazarov/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "27",
			FullName:    "Энн Хэтэуэй",
			PhotoURL:    "/static/actors/hathaway/small.webp",
			PhotoBigURL: "/static/actors/hathaway/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "28",
			FullName:    "Меган Фокс",
			PhotoURL:    "/static/actors/fox/small.webp",
			PhotoBigURL: "/static/actors/fox/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "29",
			FullName:    "Хайме Лоренте Лопес",
			PhotoURL:    "/static/actors/lorente/small.webp",
			PhotoBigURL: "/static/actors/lorente/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "30",
			FullName:    "Данила Козловский",
			PhotoURL:    "/static/actors/kozlovsky/small.webp",
			PhotoBigURL: "/static/actors/kozlovsky/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "31",
			FullName:    "Вигго Мортенсен",
			PhotoURL:    "/static/actors/mortensen/small.webp",
			PhotoBigURL: "/static/actors/mortensen/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			ID:          "32",
			FullName:    "Стася Милославская",
			PhotoURL:    "/static/actors/miloslavskaya/small.webp",
			PhotoBigURL: "/static/actors/miloslavskaya/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "33",
			FullName:    "Сергей Васильевич Маковецкий",
			PhotoURL:    "/static/actors/makovetskiy/small.webp",
			PhotoBigURL: "/static/actors/makovetskiy/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "34",
			FullName:    "Омар Си",
			PhotoURL:    "/static/actors/omar-sy/small.webp",
			PhotoBigURL: "/static/actors/omar-sy/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "35",
			FullName:    "Сеу Жоржи",
			PhotoURL:    "/static/actors/jorge/small.webp",
			PhotoBigURL: "/static/actors/jorge/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "36",
			FullName:    "Александр Яценко",
			PhotoURL:    "/static/actors/yatsenko/small.webp",
			PhotoBigURL: "/static/actors/yatsenko/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "37",
			FullName:    "Мэттью Макконахи",
			PhotoURL:    "/static/actors/mcconaughey/small.webp",
			PhotoBigURL: "/static/actors/mcconaughey/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "38",
			FullName:    "Айла Фишер",
			PhotoURL:    "/static/actors/isla-fisher/small.webp",
			PhotoBigURL: "/static/actors/isla-fisher/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "39",
			FullName:    "Том Харди",
			PhotoURL:    "/static/actors/tom-hardy/small.webp",
			PhotoBigURL: "/static/actors/tom-hardy/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "40",
			FullName:    "Джейсон Стэйтем",
			PhotoURL:    "/static/actors/jason-statham/small.webp",
			PhotoBigURL: "/static/actors/jason-statham/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			ID:          "41",
			FullName:    "Мишель Родригес",
			PhotoURL:    "/static/actors/michelle-rodriguez/small.webp",
			PhotoBigURL: "/static/actors/michelle-rodriguez/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "42",
			FullName:    "Шая ЛаБаф",
			PhotoURL:    "/static/actors/labeouf/small.webp",
			PhotoBigURL: "/static/actors/labeouf/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "43",
			FullName:    "Джона Хилл",
			PhotoURL:    "/static/actors/jonah-hill/small.webp",
			PhotoBigURL: "/static/actors/jonah-hill/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "44",
			FullName:    "Марго Робби",
			PhotoURL:    "/static/actors/margo-robbie/small.webp",
			PhotoBigURL: "/static/actors/margo-robbie/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "45",
			FullName:    "Мадс Миккельсен",
			PhotoURL:    "/static/actors/mikkelsen/small.webp",
			PhotoBigURL: "/static/actors/mikkelsen/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "46",
			FullName:    "Дмитрий Геннадьевич Нагиев",
			PhotoURL:    "/static/actors/nagiev/small.webp",
			PhotoBigURL: "/static/actors/nagiev/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "47",
			FullName:    "Елена Подкаминская",
			PhotoURL:    "/static/actors/podkaminskaya/small.webp",
			PhotoBigURL: "/static/actors/podkaminskaya/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "48",
			FullName:    "Зои Салдана",
			PhotoURL:    "/static/actors/saldana/small.webp",
			PhotoBigURL: "/static/actors/saldana/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "49",
			FullName:    "Олег Меньшиков",
			PhotoURL:    "/static/actors/menshikov/small.webp",
			PhotoBigURL: "/static/actors/menshikov/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "50",
			FullName:    "Шу Ци",
			PhotoURL:    "/static/actors/shu/small.webp",
			PhotoBigURL: "/static/actors/shu/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "51",
			FullName:    "Альваро Антонио Гарсия",
			PhotoURL:    "/static/actors/morte/small.webp",
			PhotoBigURL: "/static/actors/morte/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "52",
			FullName:    "Эмили Браунинг",
			PhotoURL:    "/static/actors/emily-browning/small.webp",
			PhotoBigURL: "/static/actors/emily-browning/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "53",
			FullName:    "Александр Наумов",
			PhotoURL:    "/static/actors/naumov/small.webp",
			PhotoBigURL: "/static/actors/naumov/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "54",
			FullName:    "Мэтт Дэймон",
			PhotoURL:    "/static/actors/matt-damon/small.webp",
			PhotoBigURL: "/static/actors/matt-damon/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "55",
			FullName:    "Александр Петров",
			PhotoURL:    "/static/actors/petrov/small.webp",
			PhotoBigURL: "/static/actors/petrov/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "56",
			FullName:    "Холт МакКэллани",
			PhotoURL:    "/static/actors/mccallany/small.webp",
			PhotoBigURL: "/static/actors/mccallany/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "57",
			FullName:    "Виктор Сухоруков",
			PhotoURL:    "/static/actors/suhorukov/small.webp",
			PhotoBigURL: "/static/actors/suhorukov/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "58",
			FullName:    "Ли Джонджэ",
			PhotoURL:    "/static/actors/jae/small.webp",
			PhotoBigURL: "/static/actors/jae/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "59",
			FullName:    "Уильям Брэдли Питт",
			PhotoURL:    "/static/actors/pitt/small.webp",
			PhotoBigURL: "/static/actors/pitt/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "60",
			FullName:    "Пак Хэ-Су",
			PhotoURL:    "/static/actors/soo/small.webp",
			PhotoBigURL: "/static/actors/soo/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "61",
			FullName:    "Чон Хо-ён",
			PhotoURL:    "/static/actors/yeon/small.webp",
			PhotoBigURL: "/static/actors/yeon/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "62",
			FullName:    "Диана Крюгер",
			PhotoURL:    "/static/actors/kruger/small.webp",
			PhotoBigURL: "/static/actors/kruger/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "63",
			FullName:    "Мигель Анхель Гарсия",
			PhotoURL:    "/static/actors/herran/small.webp",
			PhotoBigURL: "/static/actors/herran/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "64",
			FullName:    "Виктор Викторович Хориняк",
			PhotoURL:    "/static/actors/horinyak/small.webp",
			PhotoBigURL: "/static/actors/horinyak/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "65",
			FullName:    "Алессандро Нивола",
			PhotoURL:    "/static/actors/nivola/small.webp",
			PhotoBigURL: "/static/actors/nivola/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			ID:          "66",
			FullName:    "Тайриз Дарнелл Гибсон",
			PhotoURL:    "/static/actors/gibson/small.webp",
			PhotoBigURL: "/static/actors/gibson/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "67",
			FullName:    "Марк Руффало",
			PhotoURL:    "/static/actors/mark-ruffalo/small.webp",
			PhotoBigURL: "/static/actors/mark-ruffalo/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},
		{
			ID:          "68",
			FullName:    "Ольга Николаевна Кузьмина",
			PhotoURL:    "/static/actors/kuzmina/small.webp",
			PhotoBigURL: "/static/actors/kuzmina/big.webp",
			Biography:   "Педро Гонсалес Алонсо родился 21 июня 1971 года в Виго, Испания. Изучал журналистику и актерское мастерство в Высшей королевской школе драматического искусства в Мадриде...",
			Country:     "Испания",
			BirthDate:   "1971-06-21",
		},

		{
			ID:          "69",
			FullName:    "Франсуа Клюзе",
			PhotoURL:    "/static/actors/kluzet/small.webp",
			PhotoBigURL: "/static/actors/kluzet/big.webp",
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
