# aws-param-inject

golang script for config injection of AWS SSM parameter-store values.

install:
```sh
sudo curl -o /usr/local/bin/param-inject https://github.com/rbi13/aws-param-inject/releases/latest
sudo chmod +x /usr/local/bin/param-inject
```

example:
```sh
# Treat the binary as a passthrough script
export AWS_ENV_PATH='/prod/my-app'
param-inject node -e "console.log(process.env)"
# OR
AWS_ENV_PATH=/prod/my-app/ param-inject node -e "console.log(process.env)"
```

Use multiple namespaces by seperating them with semi-colons:
```sh
AWS_ENV_PATH=/prod/all;/prod/app1
```
