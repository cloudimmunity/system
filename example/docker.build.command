here="$(dirname "$BASH_SOURCE")"
cd $here

eval "$(docker-machine env default)"
docker build -t my/system-example .