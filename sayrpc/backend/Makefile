build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o app ./main.go
	sudo docker build -t gcr.io/nandotechinc/say:latest .
	rm -f app

push:
	sudo gcloud docker -- push gcr.io/nandotechinc/say:latest