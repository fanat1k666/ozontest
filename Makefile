include .env
install-dependencies:
    go install github.com/go-chi/chi
    go install github.com/graphql-go/graphql
    go install github.com/spf13/viper
    go install github.com/lib/pq

migrate:
    cd db; \
    goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable host=localhost port=5432" up

reverse-migration:
    cd db; \
    goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable host=localhost port=5432" down