package global

import (
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
	log "github.com/sirupsen/logrus"
	"time"
)

type DockerController struct{
	WorkingDir string
}

func(c *DockerController)forceCloseConaniner(id string){
	t := time.Second * 0
	if err := DockerClient.ContainerStop(context.Background(),id, &t); err != nil{
		log.Errorf("Error when close container %v\n",err)
	}else{
		log.Debugf("Container close success")
	}
	removeContainer(context.Background(),id, types.ContainerRemoveOptions{
		Force:true,
		RemoveVolumes:true,
	})

}

func(c *DockerController)SubmitContainer(ctx context.Context,cmd string,imageName string)(*Response){
	ctxx := context.Background()
	hostConfig := getRestrictedContainerHostConfig()
	networkConfig := getRestrictedContainerNetworkConfig()
	log.Debugf("Create container, Working Dir %s, mount target %s\n",c.WorkingDir,VolumePath)
	containerCreated, err := DockerClient.ContainerCreate(ctxx,&container.Config{
		WorkingDir:c.WorkingDir,
		Cmd:[]string{"sh","-c",cmd},
		AttachStdout:true,
		AttachStderr:true,
		AttachStdin:true,
		Image:imageName,
	},hostConfig,networkConfig,"")

	if err != nil{
		log.Errorf("Error when create container, %v",err)
		return generateResponse("Error when build\n")
	}

	if err := DockerClient.ContainerStart(ctxx,containerCreated.ID,types.ContainerStartOptions{

	}); err != nil{
		log.Errorf("Error when start container %v\n",err)
		return generateResponse("Error when build\n")
	}

	statusCh, errCh := DockerClient.ContainerWait(ctx,containerCreated.ID,container.WaitConditionNotRunning)
	select{
		case err := <- errCh:{
			if err != nil{
				log.Debugf("process too long, close container\n")
				c.forceCloseConaniner(containerCreated.ID)
				return generateResponse("process too long\n")
			}
		}

		case <- ctx.Done():{
			log.Debugf("process too long, close container\n")
			c.forceCloseConaniner(containerCreated.ID)
			return generateResponse("process too long\n")
		}

		case <- statusCh:{

		}
	}

	reader ,err := DockerClient.ContainerLogs(context.Background(),containerCreated.ID,types.ContainerLogsOptions{
		ShowStdout:true,
		ShowStderr:true,
	})

	if err != nil{
		log.Errorf("Error when collect container logs %v",err)
		return generateResponse("Error when build\n")
	}

	var stdout, stderr bytes.Buffer
	if _ ,err := stdcopy.StdCopy(&stdout,&stderr,reader); err != nil{
		log.Errorf("Error when use stdcopy to copy logs from container %v",err)
		return generateResponse("Error when build\n")
	}

	if err := removeContainer(context.Background(),containerCreated.ID,types.ContainerRemoveOptions{
		RemoveVolumes:true,
		Force:true,
	}); err != nil{
		return generateResponse("error when remove container " + err.Error())
	}


	return &Response{
		Error:"",
		Events:[]Event{
			{
				Kind:"stdout",
				Message:stdout.String(),
			},
			{
				Kind:"stderr",
				Message:stderr.String(),
			},
		},
	}
}

func removeContainer(ctx context.Context, id string, options types.ContainerRemoveOptions)error{
	if err := DockerClient.ContainerRemove(ctx,id,options); err != nil{
		return err
	}
	return nil
}

func generateResponse(msg string)*Response{
	return &Response{
		Error:msg,
		Events:nil,
	}
}