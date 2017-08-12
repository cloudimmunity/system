here="$(dirname "$BASH_SOURCE")"
cd $here

#eval "$(docker-machine env default)"
docker build -f Dockerfile.ubuntu -t my/system-example-ubuntu .