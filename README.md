# mini-AdEx

## Запуск

Сначала создайте `.env` из `.env.example`.

Windows PowerShell:

```powershell
Copy-Item .env.example .env
```

Linux / WSL:

```bash
cp .env.example .env
```

После этого запустите проект:

```bash
docker compose up --build
```

После запуска сервис доступен на:

```text
http://localhost:8080
```

## Пример запроса

### Windows PowerShell

```powershell
curl.exe -X POST "http://localhost:8080/auction" `
  -H "Content-Type: application/json" `
  --data-raw '{\"request_id\":\"request-1\",\"country\":\"RU\",\"device_type\":\"mobile\",\"bid_floor\":1.5,\"categories\":[\"news\",\"sport\"]}'
```

### Linux / WSL

```bash
curl -X POST "http://localhost:8080/auction" \
  -H "Content-Type: application/json" \
  --data-raw '{"request_id":"request-1","country":"RU","device_type":"mobile","bid_floor":1.5,"categories":["news","sport"]}'
```

Для этого запроса фильтрацию проходят `dsp-alpha` и `dsp-gamma`.

Пример ответа:

```json
{
  "request_id": "request-1",
  "matched_dsps": ["dsp-alpha", "dsp-gamma"],
  "sent": 2,
  "succeeded": 1,
  "duration_ms": 202
}
```

Значение `succeeded` может отличаться между запросами, так как используется fake DSP-клиент.

## DSP-партнёры

Партнёры хранятся в PostgreSQL.

Начальные партнёры добавляются в миграции:

```text
migrations/000002_seed_partners.up.sql
```

Чтобы добавить нового партнёра, нужно создать новую миграцию с добавлением записи в таблицу `partners`.

## Что бы я доделал

Если бы было больше времени:

* добавил метрики;
* расширил тесты для timeout и ошибок зависимостей;
* добавил реальный HTTP-клиент для DSP вместо fake-реализации.
