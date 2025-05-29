Отличный вопрос — вот основные команды для управления `docker-compose` 👇

---

## 🚀 Запуск и пересборка

```bash
docker-compose up --build
```

* Собирает образы и запускает контейнеры
* Автоматически применяет изменения в коде/конфигурации

---

## 🛑 Остановить и удалить контейнеры, **но не данные**

```bash
docker-compose down
```

* Останавливает контейнеры
* Удаляет их
* **Том с данными БД (`pgdata`) остаётся**

---

## 🔥 Удалить всё — контейнеры, образы, тома, сеть

```bash
docker-compose down --volumes --rmi all
```

Это:

* Очищает **всё** (контейнеры, **образы**, тома, **сеть**)
* Полезно, если хочешь **начать с нуля**
  (например, БД скинуть, зависимости пересобрать)

---

## 📦 Очистка вручную (если нужно)

### Удалить все неиспользуемые контейнеры, образы, тома:

```bash
docker system prune -a --volumes
```

> ⚠️ Используй с осторожностью — удаляет всё, что не используется прямо сейчас.

---

## ✏️ Пример типичного цикла работы:

```bash
docker-compose down --volumes
# (если хочешь сбросить БД)

docker-compose up --build
```

go mod tidy

### ✅ Решение: очистить кэш и пересобрать

Выполни последовательно:

```bash
docker builder prune -a
```

⚠️ Это удалит **весь кэш сборки**.

Затем пересобери проект:

```bash
docker-compose build --no-cache
docker-compose up
```

docker network ls

Смотри в терминал (docker-compose logs -f) или в вывод контейнера

docker exec -it <имя_контейнера_postgres> psql -U user -d digital_labels
docker exec -it tapprice-postgres-1 psql -U user -d digital_labels

goose -dir ./backend/migrations postgres "host=postgres port=5432 user=user password=pass dbname=digital_labels sslmode=disable" up


go mod tidy

docker-compose up --build

docker-compose down

\d categories;
SELECT * FROM categories;

\d venues;
SELECT * FROM venues;

docker-compose logs -f backend