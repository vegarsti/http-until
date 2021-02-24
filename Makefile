all:
	make test

test:
	go run main.go &
	until $$(curl --output /dev/null --silent --head --fail http://localhost:8080/); do \
		printf '.' \
		sleep 0.1; \
	done
