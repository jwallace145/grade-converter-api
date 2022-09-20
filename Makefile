# build the docker image given in the dockerfile
build:
	docker build -t grade-converter-api .

# authenticate to private elastic container registry for image version control and tagging
ecr_login:
	aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com

# run go-gin container locally
run:
	docker run -it --rm -p 8080:8080 --env AWS_ACCESS_KEY_ID --env AWS_SECRET_ACCESS_KEY grade-converter-api
	
# push image to aws ecr
push:	
	docker push ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/grade-converter-api:latest

# tag the latest fastapi image according to aws ecr standards
tag:
	docker tag grade-converter-api:latest ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/grade-converter-api:latest

# add missing and remove unused go modules
tidy:
	go mod tidy