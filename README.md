# todo-go
basic todo app with go

## demo video
./local-demo.mp4
## set up locally
1. Make sure you have go and git installed
2. run `git clone https://github.com/brilliant-ember/todo-go`
3. run `cd todo-go`
4. install dependencies by running `go get github.com/google/uuid`
5. run `go run main.go`
6. open a browser and navigate to localhost:8000

## Set up with docker/podman
1. run `podman build -f Dockerfile` to build an image from the dockerfile
2. create a container from the image
3. run the container, make sure to expose port 8000
4. follow steps from section "set up locally"
5. navigate to localhost:8000
