# todo-go
basic todo app with go, kubernetes, and postgres

## Set up cluster and run locally
1. build app.Dockerfile `podman build -f app.Dockerfile -t apptodo`
2. Make sure you're on the right version of postgres `docker.io/library/postgres:14.4-bullseye`
3. Setup minikube:
   1. `podman machine init --cpus 6 --memory 12288 disk-size 50`
   2. `podman machine start`
   3. `podman system connection default podman-machine-default-root` 
   4. `minikube start --driver=podman --container-runtime=cri-o`
4. Apply k8s configs - make sure to apply secrets and config maps first, then database and last the goapp. `kubectl apply -f x.yaml`
5. the cluster will expose port 30100, so go to localhost:30100 to view the webapp

## Run backend without kubernetes
1. edit `sqldb/sqldb.go` env variables and replace them with commented values, so for example replace `os.Getenv("DB_HOST")` with `"127.0.0.1"`
2. start a postgres db containe `podman run -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres` and then run it with with `podman run -d container_id`
3. run `go run main.go`
4. navigate to localhost:8000

### TODO
- use db_init.sql to init database along with building its image, and make a restricted user for the app - edit configs for k8s for that too
- add metrics and observability
