build:
	bash build.sh
clean:
	sudo rm -rf bin
docker:
	docker-compose -f dev.yml run --rm dev /bin/bash
