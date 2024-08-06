# github.com/darulfh/skuy_pay_be C-10


# Build docker
docker build -t projectname .
docker run -d -p 80:8080 projectname

docker tag <image ID>  <docker hub username>/<image name>:<version label or tag>
docker tag projectname username/projectname
docker login
docker push username/projectname
docker run -d -p 80:8080 username/projectname



// See all running containers on your machine
$ docker ps
// See all docker images on your machine
$ docker images
// Remove Docker-Whale images
$ docker rmi -f <image ID or image name>