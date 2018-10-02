package global

import(
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
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
	}
}
