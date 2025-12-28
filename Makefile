NAME := email-sender
DOMAIN := vox.buldev.com

.PHONY: clear build run

clear:
	@clear

build:
	@echo "=== Building ==="
	go build

run:
	@echo "=== Running ==="
	@eval "$$(sed -n '/^[^#]/ s/^/export /p' .env)" && ./$(NAME)
