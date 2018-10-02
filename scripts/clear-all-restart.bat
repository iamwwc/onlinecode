docker-compose down
./compile-linux.bar
cd ..
docker network prune -f
dcoerk volume prune -f
docker-compose up -d