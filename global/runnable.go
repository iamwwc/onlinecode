package global

import "context"

//Runnable return image name and cmd to exec
type Runnable interface{
	Run(ctx context.Context)(context.Context,string,string)
}
