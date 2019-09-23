deploy:
	yarn build
	tar -czvf public.tgz ./public
	gcloud compute scp public.tgz test2:~
	gcloud compute ssh test2 --command "tar -xzvf public.tgz && rm public.tgz"
	rm public.tgz