here="$(dirname "$BASH_SOURCE")"
cd $here

eval "$(docker-machine env default)"
docker run --rm -it --name example my/system-example