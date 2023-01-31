# query

Service that processes read/query operations requests

## Hooks

This repository is configured with client-side Git hooks which you need to install by running the following command:

```bash
./hooks/INSTALL
```

## Usage

To properly run this service, you will need to a setup a `.env` file. Start by creating a copy of the `.env.tpl` file and fill the variables with values appropriate for the execution context.

Then, all you need to do is to run the service with the following command:

```bash
go run cmd/query/query.go
```

## Docker

You can also run this service with Docker, but you will also need to setup the `.env` file and `.git-local-credentials`. The credentials file shall contain the git credentials config to access `damn` and `palavrapasse` private modules.

To build the service image:

```bash
docker_tag=query:latest

docker build \
    -f ./deployments/Dockerfile \
    --secret id=git-credentials,src=.local-git-credentials \
    --secret id=env,src=.env \
    . -t $docker_tag
```

To run the service container:

```bash
local_port=8080
container_port=8080

docker run -p $local_port:$container_port -t $docker_tag
```