# Описание таблиц
## genres
Таблица жанров. Хранит информацию о жанрах фильмов и сериалов
- id
- title - название жанра
- created_at
- updated_at

{id} -> title, created_at, updated_at
{title} -> id, created_at, updated_at

### Потенциальные ключи
- id
- title

### Первичный ключ
- id

## directors
Таблица хранит имя и фамилию режиссера
- id
- first_name
- second_name

{id} -> first_name, second_name

### Потенциальные ключи
- id

### Первичный ключ
- id

## users
Таблица хранит информацию о пользователе
- id
- username
- email
- avatar_url - ссылка на аватарку
- password_hash - захэшированный пароль
- created_at
- updated_at

{id} → username, email, avatar_url, password_hash, created_at, updated_at
{username} → id, email, avatar_url, password_hash, created_at, updated_at
{email} → id, username, avatar_url, password_hash, created_at, updated_at

### Потенциальные ключи
- id
- username
- email

### Первичный ключ
- id

## countries
Таблица хранит информацию о стране, которая является доп.информацией к фильму и актеру
- id
- title - название страны
- code - кодовое название страны
- flag_url - содержит ссылку на флаг страны

{id} → title, code, flag_url
{title} → id, code, flag_url
{code} → id, title, flag_url

### Потенциальные ключи
- id
- title
- code

### Первичный ключ
- id

## movies
Таблица хранит полную информацию о фильмах
- id
- title
- short_description
- long_description
- card_url - ссылка на фотокарточку
- album_url - ссылка на постер, отображаемый на странице фильма
- title_url - ссылка на картинку с названием
- release_date - дата выпуска
- movie_type - тип(фильм или сериал)
- country_id
- director_id
- created_at
- updated_at

{id} -> title, short_description, long_description, card_url, album_url, title_url, release_date, movie_type, country_id, director_id, created_at, updated_at

### Потенциальные ключи
- id

### Первичный ключ
- id

## episodes
Таблица хранит полную информацию о сериях сериалов
- id
- title
- description
- season_number
- episode_number
- movie_id
- preview_url - ссылка на превью серии
- release_date
- created_at
- updated_at

{id} -> title, description, season_number, episode_number, movie_id, preview_url, release_date, created_at, updated_at
{season_number, episode_number, movie_id} -> id, title, description, preview_url, release_date, created_at, updated_at

### Потенциальные ключи
- id
- season_number, episode_number, movie_id

### Первичный ключ
- id

## rating
Таблица хранит полную информацию об оценках пользователей
- id
- user_id
- movie_id
- episode_id
- rating
- created_at
- updated_at

{id} -> user_id, movie_id, episode_id, rating, created_at, updated_at
{user_id, movie_id, episode_id} -> id, rating, created_at, updated_at

### Потенциальные ключи
- id
- season_number, episode_number, movie_id

### Первичный ключ
- id

## movie_qualities
Таблица, связывающая фильмы и качество видео в отношении много ко многим
- video_url
- movie_id
- quality_id

{movie_id, quality_id} -> video_url

### Потенциальные ключи
- movie_id, quality_id

### Первичный ключ
- movie_id, quality_id

## qualities
Таблица хранит виды качества фильмов
- id
- quality

{id} -> quality
{quality} -> id

### Потенциальные ключи
- id
- quality

### Первичный ключ
- id

## actors
Таблица хранит полную информацию об актерах
- id
- first_name
- second_name
- country_id
- small_photo_url - фотокарточка актера
- big_photo_url - фотка актера на странице актера
- birthdate
- biography

{id} -> - first_name, second_name, country_id, small_photo_url, big_photo_url, birthdate, biography
{first_name, second_name} -> id, country_id, small_photo_url, big_photo_url, birthdate, biography

### Потенциальные ключи
- id
- first_name, second_name

### Первичный ключ
- id

## movie_actors
Таблица для связывания актеров и фильмов в отношении много ко многим
- movie_id
- actor_id

{movie_id, actor_id} -> {}

### Потенциальные ключи
- movie_id, actor_id

### Первичный ключ
- movie_id, actor_id

## movie_genres
Таблица для связывания жанров и фильмов в отношении много ко многим
- movie_id
- genre_id

{movie_id, genre_id} -> {}

### Потенциальные ключи
- movie_id, genre_id

### Первичный ключ
- movie_id, genre_id

## collections
Таблица хранит полную информацию по подборкам кино
- id
- title
- card_url - фотокарточка подборки
- album_url - полноразмерная картинка подборки
- created_at
- updated_at

{id} -> title, card_url, album_url, created_at, updated_at
{title} -> id, card_url, album_url, created_at, updated_at

### Потенциальные ключи
- id
- title

### Первичный ключ
- id

## movie_collections
Таблица для связывания подборок и фильмов в отношении много ко многим
- movie_id
- collection_id
- created_at
- updated_at

{movie_id, collection_id} -> created_at, updated_at

### Потенциальные ключи
- movie_id, collection_id

### Первичный ключ
- movie_id, collection_id

## subscriptions
Таблица хранит полную информацию о подписках
- id
- price
- user_id
- start_date - время с которого работает подписка(Отличается от created_at, вдруг платежи упадут)
- days_counter - время действии подписки с начала start_date в днях
- created_at

{id} -> price, user_id, start_date, days_counter, created_at

### Потенциальные ключи
- id

### Первичный ключ
- id

## payments
Таблица хранит полную информацию о платежах
- id
- captured_total - списанные средства
- refunded_total - возвращенные средства
- subscription_id - id подписки, по которой совершен платеж
- payment_number - номер платежа
- status - статус платежа
- created_at
- updated_at

{id} -> captured_total, refunded_total, subscription_id, payment_number, status, created_at, updated_at
{subscription_id, payment_number} -> id, captured_total, refunded_total, status, created_at, updated_at

### Потенциальные ключи
- id
- subscription_id, payment_number

### Первичный ключ
- id

## receipts
Таблица хранит полную информацию о чеках
- id
- payment_id - id платежа, по которому выбивается чек
- receipt_type - тип чека(Приход или Возврат)
- created_at
- updated_at

{id} -> payment_id, receipt_type, created_at, updated_at

### Потенциальные ключи
- id

### Первичный ключ
- id

## line_items
Таблица хранит полную информацию о товарных позициях в чеке
- id
- receipt_id - id чека, с которым связана товарная позиция
- title - название товарной позиции
- total - цена товарной позиции
- line_item_type - тип товарной позиции (Полная оплата, частичная оплата, возврат)
- created_at
- updated_at

{id} -> receipt_id, title, total, line_item_type, created_at, updated_at
{receipt_id, title} -> id, total, line_item_type, created_at, updated_at

### Потенциальные ключи
- id
- receipt_id, title

### Первичный ключ
- id


## Доказательства соответствия нормальным формам
### 1НФ
Все атрибуты отношений являются простыми, все используемые домены содержат только скалярные значения. Повторений строк в таблицах нет.

### 2НФ
Выполнены условия 1НФ, а также каждый неключевой атрибут функционально-полно зависит от первичного ключа.

### 3НФ
Выполнены условия 2НФ, а также каждый неключевой атрибут нетранзитивно зависит от первичного ключа.

### НФБК
1. Для отношений, имеющих один первичный ключ, НФБК является 3НФ. Т.к. у всех описанных отношений первичный ключи единственный, то отношения находятся в НФБК.
2. Каждая нетривиальная и неприводимая слева функциональная зависимость обладает потенциальным ключом в качестве детерминанта.
3. Для любой нетривиальной функциональной зависимости {X} → Y, X является надключом.
