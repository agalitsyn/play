# golang prometheus to influx

## start etcd
go get github.com/coreos/etcd 
cd $GOPATH/github.com/coreos/etcd
git checkout v2.3.7
goreman start

## start influx
docker run -p 8083:8083 -p 8086:8086 -v $PWD:/var/lib/influxdb influxdb:0.13.0

## create db

## run script
