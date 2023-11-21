# Test task for "Wildberries"

```bash
docker compose --file docker-compose-postgresql.yml up --detach --build
```

```bash
docker compose --file docker-compose-nats-streaming.yml up --detach --build
```

```bash
docker compose --file docker-compose-microservices.yml up --detach --build
```

###### Get Employee By Id

```bash
curl \
    --location '192.168.50.100:3001/v1/employees/df6c64ec-4548-467b-a5b2-5e5f9aa442ca' \
    --header 'authentication-token: token'
```

###### Create Employee

```bash
curl \
    --location '192.168.50.100:3001/v1/employees' \
    --header 'Content-Type: application/json' \
    --header 'authentication-token: token' \
    --data '{
        "firstname": "firstname",
        "lastname": "lastname",
        "email": "email",
        "phone_number": "phone_number",
        "address": "address",
        "position": "position",
        "department": "department",
        "date_of_birth": "2023-11-21T08:17:13.599Z",
        "hire_date": "2023-11-21T08:17:13.599Z"
    }'
```

