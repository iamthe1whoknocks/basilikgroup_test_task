build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./cmd/app/app ./cmd/app/

test:
	go test -v -race ./pkg/handler

run:
	docker build -t bazilikgroup_test_task -f Dockerfile .
	docker run -it -p 8086:8086 bazilikgroup_test_task
