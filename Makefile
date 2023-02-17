OUTPUT_DIR = out

build:
	mkdir -p "$(OUTPUT_DIR)"
	go build -o "$(OUTPUT_DIR)/podman-toolbox"

format:
	gofmt -s -w .

start-podman-service:
	podman system service -t 0 &