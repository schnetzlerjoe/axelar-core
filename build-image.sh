docker build -f Dockerfile.rocksdb -t axelard-rocksdb-v0.13.6 .
docker tag axelard-rocksdb-v0.13.6 kalidux/axelard-rocksdb:v0.13.62
docker push kalidux/axelard-rocksdb:v0.13.62

