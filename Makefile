build:
	go build -o nate .
	chmod +x nate
	sudo mv nate /usr/local/bin 