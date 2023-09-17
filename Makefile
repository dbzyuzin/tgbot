run:
	@ while read line; do export $line; done < dev.env
	@ go run cmd/main.go