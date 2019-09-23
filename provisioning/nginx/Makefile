deploy:
	gcloud compute ssh test2 --command "mkdir conf.d"
	gcloud compute scp ./conf.d/*.conf test2:~/conf.d/
	gcloud compute ssh test2 --command "sudo cp -r conf.d/ /etc/nginx/ && rm -rf conf.d && sudo systemctl restart nginx"