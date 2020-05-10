# ton-dice-web-server
Implement REST API service for frontend and gRPC storage service for ton-dice-web-worker.


## build 
```
export GITHUB_TOKEN = 'token'
docker build --build-arg GITHUB_TOKEN="$GITHUB_TOKEN" -t dice-server .
```

## codegen
For sql client:
```
cd storage; salgen -destination=./client.go -package=github.com/tonradar/ton-dice-web-server/storage github.com/tonradar/ton-dice-web-server/storage Store
```

For protobug grpc:
```
cd proto; protoc -I . bets.proto --go_out=plugins=grpc:.
```


## run (develop)
```docker run --name dice-server --network dice-network -e PG_HOST=pg-docker -e PG_PORT=5432 -e PG_NAME=postgres -e PG_USER=postgres -e PG_PWD=docker -e TON_API_HOST=ton-api -e TON_API_PORT=5400 -d -p 8080:9999 dice-server```

## ENV VARS
    * PG_HOST - Postgres host, default value is 'localhost'
    * PG_PORT - Postgres port, default value is `5432`.
    * PG_NAME - Postgres database name, required variable, no default value.
    * PG_USER - Postgres user, required variable, no default value.
    * PG_PWD -  Postgres password, required variable, no default value.
    * RPC_LISTEN_PORT - Listen port for GRPC server, default value is '5300'.
    * WEB_LISTEN_PORT - Listen port for web server, default value is '9999'.
    * WEB_DOMAIN - Domain name used for CORS headers, default value is 'tonbet.io'.
    * TON_API_HOST  - Host of the 'ton-api' service, requred variable, no default value.
    * TON_API_PORT - Port of the 'ton-api' service, default value is '5400'. 