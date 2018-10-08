### Demo: https://onlinecode.chaochaogege.com

*因为没备案，现在只能用https蒙混一下*

**服务器没多少内存，所以`Run`功能用不了**

#### 如何使用？
全部的镜像都编译完成了，你只需要下载这里的`docker-compose.yml`文件，然后运行`docker-compose up -d`，
`docker`会自己`pull`所需要的镜像然后运行

#### 开发

关于本地IDE的配置，如果使用的是`Goland`
那么可以参考 https://blog.jetbrains.com/go/2018/04/30/debugging-containerized-go-applications/

打开`docker-compose.yml`,注释掉web service, 运行`docker-compose up -d`，然后就可以正常使用了，调试的时候选中`golang-delve`启动选项，里面已经配置好了在运行之前会先启动
调试`container`

预编译了两个调试镜像
`Dockerfile-debug-server-manual-start`和`Dockerfile-debug-server-run-autmatic`

其中
  1.`manual-start`: docker entrypoint 是`tail -f /dev/null`，确保不退出容器，需要拿到容器shell之后进去手动开启远程调试，如果需要线上问题代码的话，使用这个远程挂载，然后`delve exec`

  2. `run-automatic`, 自动运行`delve debug`，8088端口用于调试，8086提供静态资源文件，如果想要自定义调试过程，使用`manual-start`


#### 部署的时候，打开`docker-compose.yml`文件，注释掉`debug-service`，取消`web`的注释
运行
```
docker-compose up -d 
```

##### 注意
挂载目录的问题，docker-compose.yml文件需要放到onlinecode，因为docker-compose会为创建的volume和network添加目录前缀，
会导致创建执行容器的时候temp卷挂载不上去

---------------------------

使用Docker Container对每一份代码进行了环境隔离，进行了权限限制，每一个用于运行代码的容器进行了资源限制，但是由于不知道到底需要多少资源去运行，所以Run功能会由于没有足够的资源被`Kill`

**如果有兴趣知道更多的细节，你可以看[英文版本](https://github.com/onlinecode/docs/en.md) :)**