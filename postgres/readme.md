# Setup docker-compose
```
services:
  db:
    image: postgres
    container_name: postgres-db
    restart: always
    environment:
      POSTGRES_DB: <db_name>
      POSTGRES_USER: <db_user>
      POSTGRES_PASSWORD: <db_password>
    ports:
      - "5432:5432"
    volumes:
      - pg-data:/var/lib/postgresql/data
volumes:
    pg-data:
      driver: local
      driver_opts:
        o: bind
        type: none
        device: <full path map>

```