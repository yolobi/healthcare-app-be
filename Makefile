dbuser := postgres
dbname := healthcare_db

run:
	@go run ./cmd/healthcare/cmd/ -config ./files/yaml/configs/healthcare.local.yml

build:
	@go build -o ./main ./cmd/healthcare/cmd/main.go
	./main -config=./files/yaml/configs/healthcare.local.yml

automigrate:
	@go run ./cmd/healthcare/database/migrate -config ./files/yaml/configs/healthcare.local.yml	

seeding:
	@go run ./cmd/healthcare/database/seed -config ./files/yaml/configs/healthcare.local.yml

drop:
	@go run ./cmd/healthcare/database/drop -config ./files/yaml/configs/healthcare.local.yml

refresh: drop automigrate seeding

dockerbuildstart:
	docker compose up --build

dockerstart:
	docker compose start

dockerstop:
	docker compose stop

dockerdown:
	docker compose down

refreshpgdata:
	sudo chmod 666 db/pg_data
	sudo rm -rf db/pg_data

dockerpg:
	docker exec -it postgres_healthcare psql -U postgres healthcare_db

dockerseeding:
	docker exec -it healthcare-be-healthcare-api-1 sh -c "./seed -config ./files/yaml/configs/healthcare.docker.yml"

dbbackup:
	pg_dump -U ${dbuser} -h localhost -p 5432 healthcare_db > healthcare_db_backup.sql

dropdb:
	sudo dropdb -h localhost -U ${dbuser} -p 5432 ${dbname}

dockerlogin:
	sudo docker login

dockertag:
	sudo docker tag healthcare-api varmaseaapp/healthcare-api

dockerpush:
	sudo docker push varmaseaapp/healthcare-api