run:
	go run ./main.go

watch:
	air -c watcher.conf

upd:
	docker-compose -f docker-compose.dev.yml up -d
	docker-compose logs -f

downd:
	docker-compose -f docker-compose.dev.yml down