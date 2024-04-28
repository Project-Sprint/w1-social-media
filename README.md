# w1-social-media
CatsSocial adalah aplikasi dimana pemilik kucing dapat saling menjodoh-jodohkan kucingnya


go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate create -ext sql -dir db/migration table_name

migrate -database "postgres://postgres:your_password@localhost:5432/your_database?sslmode=disable" -path db/migration up

migrate -database "postgres://postgres:your_password@localhost:5432/your_database?sslmode=disable" -path db/migration down