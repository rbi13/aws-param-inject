# aws-param-inject

golang script for config injection of AWS SSM parameter-store values.

example:
```sh
export AWS_ENV_PATH='/prod/my-app'
eval $(./aws-env) && node -e "console.log(process.env)"
# OR
eval $(AWS_ENV_PATH=/prod/my-app/ ./aws-env) && node -e "console.log(process.env)"
```

Use multiple namespaces by seperating them with semi-colons:
```sh
AWS_ENV_PATH=/prod/all;/prod/app1
```
