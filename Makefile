include .env
export

run:
	@ while read line; do export $line; done < .env
	@ go run cmd/main.go
