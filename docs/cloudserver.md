# Cloud Server
The cloud server allows for organizations and users to turn Tensorgo into a single API endpoint. This allows for developers and other users to call the endpoints to create tensors, neural networks, and things such as running an activation function. This can all be done without having to install Tensorgo on the local machine. 

## Setup Pangea
Tensorgo uses Pangea APIs for security and logging. If you want to use it, you can export your token and domain to the machine you are running the server on. 

## Running Cloud Server
### Web Server
```golang
package main

import (
	"github.com/mac-lawson/tensorgo/cloudserver"
)

func main() {
	cloudserver.RunAPIServer()
}
```
In your browser, navigate to [127.0.0.1:8080](localhost:8080)