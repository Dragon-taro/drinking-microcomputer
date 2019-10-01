deploy:
	yarn build
	tar -czvf public.tgz ./public
	gcloud compute scp public.tgz drinking-microcomputer:~
	gcloud compute ssh drinking-microcomputer --command "tar -xzvf public.tgz && rm public.tgz"
	rm public.tgz