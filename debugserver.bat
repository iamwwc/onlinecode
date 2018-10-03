docker run -itd --network onlinecode_onlinecode-net ^
    -v %cd%/:/go/src/chaochaogege.com/onlinecode -v onlinecodetemp:/onlinecodetemp ^
    --name debugserver -p 8086:8086 -p 8088:8088 ^
     -e CODE_WORK_DIR="/go/src/chaochaogege.com/onlinecode" ^
    --security-opt=seccomp:unconfined ^
    ccr.ccs.tencentyun.com/wwc-docker-images/onlinecode:golang-delve


docker exec -it  debugserver sh

