docker run \
	-e POSTGRES_PASSWORD=password \
	-e POSTGRES_USER=user \
	-e POSTGRES_DB=database \
	-p 5432:5432 \
	postgres:alpine
