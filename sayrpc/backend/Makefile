build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o app ./main.go
	sudo docker build -t nandotech/say .
	# rm -f app