ws-server/server.bundle.js:
	yarn build

ws-server.tgz: ws-server/server.bundle.js
	tar -czvf ws-server.tgz ./ws-server/

upload/tar: ws-server.tgz
	gsutil cp ./ws-server.tgz gs://tar-for-deploy/ws/
	gsutil cp ./Makefile.prd gs://tar-for-deploy/ws/Makefile
	rm -rf ws-server
	rm ws-server.tgz

deploy: upload/tar
	ssh dragontaro@drinking.dragon-taro.dev "gsutil cp gs://tar-for-deploy/ws/Makefile ./ws/ && cd ws && make run"
