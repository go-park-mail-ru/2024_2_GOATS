#!/bin/bash

start_compose_service() {
    local service_path=$1

    echo "Запуск docker-compose для $service_path"

    # Проверяем, существует ли docker-compose.yml
    if [ ! -f "$service_path/docker-compose.yml" ]; then
        echo "docker-compose.yml не найден в $service_path. Пропуск..."
        return
    fi

    pushd "$service_path" > /dev/null
    docker-compose down -v
    docker-compose up --build
    popd > /dev/null
}

echo "Сборка Docker образа для корня проекта"
docker-compose down -v
docker-compose up --build

services=("user_service" "auth_service" "movie_service" "payment_service")
for service in "${services[@]}"; do
    start_compose_service "$service"
done

echo "Проверяем статус всех контейнеров..."
docker ps

echo "Все сервисы запущены!"
