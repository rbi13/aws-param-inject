# aws-param-inject

golang script for config injection of AWS SSM parameter-store values.

install:
```sh
sudo curl -L -o /usr/local/bin/param-inject https://github.com/rbi13/aws-param-inject/releases/download/0.0.3/aws-param-inject-linux-amd64
sudo chmod +x /usr/local/bin/param-inject

# Latest
INJ_RELEASE=$(curl -L -s -H 'Accept: application/json' https://github.com/account/project/releases/latest)
INJ_VERSION=$(echo $INJ_RELEASE | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
https://github.com/rbi13/aws-param-inject/releases/download/${INJ_RELEASE}/aws-param-inject-linux-amd64
sudo chmod +x /usr/local/bin/param-inject

# Docker example
...
RUN \
  INJ_RELEASE=$(curl -L -s -H 'Accept: application/json' https://github.com/account/project/releases/latest)\
  && INJ_VERSION=$(echo $INJ_RELEASE | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')\
  && curl -L -o /usr/local/bin/param-inject "https://github.com/rbi13/aws-param-inject/releases/download/$INJ_RELEASE/aws-param-inject-linux-amd64"\
  sudo chmod +x /usr/local/bin/param-inject
...
```

example:
```sh
# Treat the binary as a passthrough script
export AWS_REGION='us-east-1'
export AWS_ENV_PATH='/production/app1'
param-inject node -e "console.log(process.env)"
# OR
AWS_REGION='us-east-1' AWS_ENV_PATH='/production/app1/' param-inject node -e "console.log(process.env)"
```

Use multiple namespaces by seperating them with semi-colons:
```sh
AWS_ENV_PATH='/production/all;/production/app1'
```

## Development

A development environment can be created using docker containers:

```sh
make docker
```

This will use the `docker-compose` config in `dev.yml` giving you a shell into a
container running the official `golang` dockerhub image. modify the
`dev.yml` config to edit/augment the development environment (e.g. to set your
`AWS_ENV_PATH`). From here it is golang business as usual:

```sh
go get
go run inject.go env
```
