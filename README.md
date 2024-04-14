# Banner service
Сервис баннеров предназначен для динамического отображения информационных баннеров пользователям, основываясь на их принадлежности  к определенным группам или интересам. Это достигается за счет ассоциации баннеров с различными тегами и фичами, что позволяет предлагать контент для соответствующих сегментов аудитории.

Используемые технологии:
- Go
- Swagger
- PostgreSQL
- Docker

## Начало работы

Склонируйте репозиторий:
```shell
git clone https://github.com/DrozdovRoman/avito-tech-banner-service.git
```

Скопируйте и заполните конфигурации в файле .env
```shell
cp .env.example .env
```

Установите зависимости:
```shell
make install-deps
```

Выполните запуск базы данных и установку миграций:
```shell
make up-db
make local-migration-up
```

Запустите приложение:
```shell
make run
```


## Примеры запросов в приложении

### Получение токена

- Токен обычного пользователя (Для получения токена необходимо ввести любой username и password)
```shell
curl --location --request POST 'http://localhost:8000/login' \       
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"dannromm",
    "password":"Qwerty123!"
}'
```

- Токен админа (Для получения токена необходимо ввести любой username и пароль "avitoDeveloper")
```shell
curl --location --request POST 'http://localhost:8000/login' \       
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"dannromm",
    "password":"avitoDeveloper"
}
```


## Методы взаимодействия с баннером

## Admin
### Создание баннера

```shell
curl -X POST http://localhost:8000/banner \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer your_admin_token" \
    -d '{
          "feature_id": 6,
          "tag_ids": [2, 3],
          "content": "{\"Example\": \"Hello world!\"}"
      }'
```

### Удаление баннера

```shell
curl -X DELETE http://localhost:8000/banner/54 \
    -H "Authorization: Bearer your_admin_token"
```

### Изменение баннера

```shell
curl -X PATCH http://localhost:8000/banner/54 \
    -H "Authorization: Bearer your_admin_token"
    -d '{
          {
            "is_active": true,
            "feature_id": 2,
            "tag_ids": [1, 2, 3]
            "content": "{\"hello world\":\"abc\"}"
          }
    }
```

### Получение баннеров

```shell
curl -X GET http://localhost:8000/banner \
    -H "Authorization: Bearer your_admin_token"
    }
```

## User

### Полученния баннера для пользователя 

```shell
curl -X GET "http://localhost:8000/user_banner?tag_id=10&feature_id=20&use_last_revision=false" \
    -H "Authorization: Bearer your_user_toke"
```

# Results
- Я постарался реализовать максимально приближенное к production приложение, для этого были использованы основные инструменты в современной разработке. Данное приложение построено с применением ключевых принцепов Domain Driven Design, оно было разделено на 4 слоя (presentation, infrastructure, domain, application). В основе приложения был использован DI контейнер от uber/FX, что в дальнейшем бы позволило значительно проще масштабировать приложение и внедрять в него различные зависимости. Огромная работа была проведена, для того, что бы приложение сделать максимально масштабируемым в дальнейшем. Здесь очень легко добавлять новые модули, например cache, который был реализован in-memory в дальнейшем можно было бы заменить на redis. Я очень надеюсь, что человек который это будет просматривать, хотя бы дочитает до этого момента, а я могу пожелать только великолепонго дня и буду надеяться, что 40 часов моей жизни прошли не зря <3.

