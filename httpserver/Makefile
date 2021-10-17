export tag=xwjv1.0
#root:
#	export ROOT=github.com/fireman-coder/goexercise

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t httpserver:${tag} .

push: release
	echo "pushing goexercise/httpserver"
	docker tag httpserver:${tag} firemancoder/httpserver:${tag}
	docker push firemancoder/httpserver:${tag}
