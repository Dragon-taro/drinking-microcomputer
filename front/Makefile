deploy:
	tar -czvf public.tgz ./public
	gcloud compute scp public.tgz dragontaro@drinking-microcomputer:~
	gcloud compute ssh dragontaro@drinking-microcomputer --command "tar -xzvf public.tgz && rm public.tgz"
	rm public.tgz