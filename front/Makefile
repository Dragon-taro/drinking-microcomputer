deploy:
	tar -czvf public.tgz ./public
	scp public.tgz dragontaro@drinking.dragon-taro.dev:~
	ssh dragontaro@drinking.dragon-taro.dev "tar -xzvf public.tgz && rm public.tgz"
	rm public.tgz