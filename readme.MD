# Start Project gorm + postgresql + fiber

# Run Docker
```
 cd postgres
 export PATH_DB_TEST=<your full path to map files from docker> DBNAME=<db_name> DBUSER=<db_user> DBPASSWORD=<db_password>
 docker-compose -d up

 if you have .env file on root project >> 
 docker-compose -d --env-file ../config/.env up

```

# Run Migrate
```
 go run cmd/migrate.go
```