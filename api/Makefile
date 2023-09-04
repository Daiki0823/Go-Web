help:
	@echo "Available targets:"
	@echo "  build         Build the Docker image"
	@echo "  run           Run the API with hot reload"
	@echo "  run-docker    Run the Docker container"
	@echo "  start-db      Start the MySQL container"
	@echo "  stop-db       Stop and remove the MySQL container"

build:
		docker build -t myapi .

run:
	STAGE=local air

run-docker:
		docker run -p 1323:1323 myapi

# Start MySQL container
start-db:
	docker run --name todo-db -e MYSQL_ROOT_PASSWORD=rootpassword -e MYSQL_DATABASE=tododb -p 3306:3306 -d mysql:latest

# Stop and remove MySQL container
stop-db:
	docker stop todo-db
	docker rm todo-db