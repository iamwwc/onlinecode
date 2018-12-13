### Chinese Language
### [中文文档](https://github.com/iamwwc/onlinecode/docs/zh-cn.md)

-----------------------------------------------------------

**My server not have enough memory, so run function maybe not works.**

### Demo: https://onlinecode.chaochaogege.com

#### Features
- Use Docker to restrict untrusted codes behaviors.

- Every codes run in single container, container will be destroyed after some seconds.

- Share your workspace, anyone can access your codes if you share your `share url`.

*Only support `share`, `run` features*


#### For development

`onlinecode` use `docker-compose` start two containers.
1. `onlinecodeapp`:  Run codes and serve files
2. `onlinecodedb`: Store all settings of workspace.

By default, all docker images have been compiled and store into docker hub `iamwwc/onlinecode`, tag is compile count of travis.

So in order to make all things works, You should use latest images

There are three Dockerfile.

1. `Dockerfile`: a final version, all necessary files have been store into it.
2. `DOckerfile-debug-server-run-automatic`: remote debug images used by golang codes. When container start,  a delve debug process started automatic. run `delve debug` 
    - export `8088` for debug
    - export `8086` for serve static resources containers `js`,`css`
3. `Dockerfile-debug-server-manual-start`: Difference between `automatic` and `manual` is when container started, `manual` only run `tail -f /dev/null` which make sure container will not exit, you need to control `delve`
behaviour. 

*For example, use `delve exec` to debug production version `onlinecode executable`*

View `Dockerfile***` for more information.

`onlinecodeapp` and `onlinecodedb` use docker subnet called `onlinecode_onlinecode-net`, only `onlinecodeapp` can access it.

`onlinecodedb` mount a volume called `onlinecode_onlinecode-database`, if you want to migrate your databases, it's very useful

**When create volume and network, docker-compose will add dir name to your volume and network, if you use another name called newName instead of onlinecode, volume and network names will have `newName` prefix**

**Tips**

I am use `Docker-for-Windows`, because of different file system structure between Windows and Linux, some codes cannot run Windows directly, So I am no choice but use remote debug.
If you development in Linux, You don't do follow steps.


#### Remote Debug
In `docker-compose.yml`, these are three services but one has been committed.
When in development. uncommit `debug-server` service and commit `onlinecodeapp`, `debug-service` will mount workspace to container, so that `delve` can compile, used for remote debug.

If you want compile codes by yourself. Your need `nodejs` `golang` environment.

I compile a image, you can use it directly if you need.

```
docker pull iamwwc/node-golang:latest
```


