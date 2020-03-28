# ton-dice-web-server
rest-api && storage

TODO:
    - change ssh key to access token in Dockerfile
    - remove hardcode (domains, ports)
    - 
## build 
```MY_KEY=$(cat ~/.ssh/id_rsa)```

```docker build --build-arg SSH_PRIVATE_KEY="$MY_KEY" -t dice-server .```

## run
```docker run --name dice-server --network dice-network -e PG_HOST=pg-docker -e PG_PORT=5432 -e PG_NAME=postgres -e PG_USER=postgres -e PG_PWD=docker -e TON_API_HOST=ton-api -e TON_API_PORT=5400 -d -p 8080:9999 dice-server```
