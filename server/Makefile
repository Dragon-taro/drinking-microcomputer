run:
	JSON_PATH=credential/firestore.json go run main.go

dist/server:
	GOOS=linux \
	GOARCH=amd64 \
	go build -o dist/server main.go

build:
	go build -o dist/server main.go
	cp -r credential/ dist/credential
	cp -r .env dist/.env

deploy: dist/server
	tar -czvf dist.tgz dist/
	gsutil cp ./dist.tgz gs://tar-for-deploy
	gsutil cp ./Makefile.prd gs://tar-for-deploy/Makefile
	gcloud compute ssh test2 --command "gsutil cp gs://tar-for-deploy/Makefile . && make run"
	rm -rf dist
	rm dist.tgz
