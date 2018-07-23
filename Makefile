$(BINARY): check-env
	cd $(BINARY) && docker build -t kdobmayer/$(BINARY):v2 .
	docker push kdobmayer/$(BINARY):v2

distclean:
	docker rm -f `docker ps -aq`
	docker image rm kdobmayer/backend
	docker image rm kdobmayer/frontend
	docker image prune --force

check-env:
ifndef BINARY
  $(error BINARY is undefined)
endif

.PHONY: check-env
