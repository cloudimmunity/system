here="$(dirname "$BASH_SOURCE")"
cd $here

#eval "$(docker-machine env default)"
docker build -f Dockerfile.alpine -t my/system-example-alpine .