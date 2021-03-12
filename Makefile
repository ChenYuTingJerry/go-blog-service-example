up:
	docker-compose up --build
up_web:
	docker-compose up web
up_db:
	docker-compose up db
down:
	docker-compose down
prune:
	docker system prune