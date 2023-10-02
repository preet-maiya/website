VERSION=dev
DOCKERUSER=preethammaiya


build:
	docker build -t website:dev .
	docker tag website:dev $(DOCKERUSER)/website:$(VERSION)

build-dev:
	docker build -t website:dev .

push:
	docker tag website:dev $(DOCKERUSER)/website:$(VERSION)
	docker push $(DOCKERUSER)/website:$(VERSION)
	# docker tag website $(DOCKERUSER)/website:latest
	# docker push $(DOCKERUSER)/website:latest

deploy-dev:
	docker run --rm --name website -d -p 8080:8080 -v $(pwd)/blog/resume.pdf:/resume.pdf preethammaiya/website:dev

delete:
	docker stop website