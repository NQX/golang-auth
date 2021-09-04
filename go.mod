module github.com/nqx/golang-auth

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/jackc/pgx/v4 v4.13.0
	github.com/pkg/errors v0.9.1 // indirect
	golang-auth/models v0.0.0-00010101000000-000000000000 // indirect
	golang-auth/routes v0.0.0-00010101000000-000000000000
)

replace golang-auth/routes => ./routes

replace golang-auth/models => ./models
