qtc:
	@qtc -dir=views -ext=html && qtc -dir export -ext=sql && qtc -dir export -ext=csv

watch:
	@go run cmd/fileWatcher/fileWatcher.go

sqlc:
	@sqlc generate

build:
	@go build -ldflags "-w -s" main.go

rename:
	@find . -type f -exec sed -i -e 's/github.com\/Jiang-Gianni\/htmx-go/github.com\/Jiang-Gianni\/$(name)/g' {} \;

reset:
	@go run cmd/resetSqlite/resetSqlite.go

sd:
	@sudo systemctl start docker

prune:
	@docker container prune

pull:
	@docker pull mysql && docker pull postgres

mysql:
	docker run --rm -it --name local-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_DATABASE=mydb -d mysql

km:
	docker kill $$(docker ps -a -q --filter ancestor=mysql) && docker container prune

postgres:
	docker run --rm -it --name local-postgres -p 5432:5432 -e POSTGRES_PASSWORD=my-secret-pw -e POSTGRES_USER=root -e POSTGRES_DB=mydb -d postgres

kp:
	docker kill $$(docker ps -a -q --filter ancestor=postgres) && docker container prune