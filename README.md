#### 开发

使用Goland IDE打开，idea下会有配置文件，里面有个golang-delve启动文件
打开`docker-compose.yml`,注释掉web service, 运行`docker-compose up -d`，然后就可以正常使用了，调试的时候选中`golang-delve`启动选项，里面已经配置好了在运行之前会先启动
调试container

#### 部署的时候，打开`docker-compose.yml`文件，注释掉`debug-service`，取消`web`的注释，然后使用travis就可以部署了

##### 注意
挂载目录的问题，docker-compose.yml文件需要放到onlinecode，因为docker-compose会为创建的volume和network添加目录前缀，
会导致创建执行容器的时候temp卷挂载不上去

####
使用Docker Container对每一份代码进行了环境隔离，进行了权限限制

