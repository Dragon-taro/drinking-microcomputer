deploy:
	tar -czvf public.tgz ./public
	gsutil cp ./public.tgz gs://tar-for-deploy/front/
	gsutil cp ./Makefile.prd gs://tar-for-deploy/front/Makefile
	ssh dragontaro@drinking.dragon-taro.dev "gsutil cp gs://tar-for-deploy/front/Makefile ./front/ && cd front && make run"
	rm public.tgz