package global

import(
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
)

func getRestrictedContainerHostConfig()*container.HostConfig{
	return &container.HostConfig{
		Mounts:[]mount.Mount{
			{
				Source:VolumeName,
				Target:VolumePath,
				Type:mount.TypeVolume,
			},
		},
		AutoRemove:true,
		Resources:container.Resources{
			MemorySwap:2e+7,
			Memory:1e+7,
			NanoCPUs:1,
		},
		NetworkMode:"none",
		RestartPolicy:container.RestartPolicy{
			Name:"no",
			MaximumRetryCount:0,
		},
		Privileged:false,
	}
}

func getRestrictedContainerNetworkConfig()*network.NetworkingConfig{
	return &network.NetworkingConfig{
		EndpointsConfig:map[string]*network.EndpointSettings{
			"none":&network.EndpointSettings{

			},
		},
	}
}