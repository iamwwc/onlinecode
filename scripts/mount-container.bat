
REM 启动一个容器，并且将当前文件夹挂载到 golang 镜像下的 /go/ 目录下
REM 在开发环境下，挂在windows temp目录,否则在golang container写入的文件不能经docker daemon 传递到新的container
REM docker network create -d bridge onlinecode-net

docker run -itd -v %cd%:/go/src/onlinecode -v /var/run/docker.sock:/var/run/docker.sock -v onlinecode_onlinecodetemp:/onlinecodetemp --name test -p 8086:8086 golang tail -f /dev/null

REM docker run -itd --name mysqlcontainer -p 3306:3306 --network onlinecode-net -e MYSQL_ROOT_PASSWORD=mysqlpassword mysql:8.0
