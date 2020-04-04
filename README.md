# ton-dice-web-server
rest-api && storage

TODO:
    - change ssh key to access token in Dockerfile
    - remove hardcode (domains, ports)
## build 
```MY_KEY=$(cat ~/.ssh/id_rsa)```

```docker build --build-arg SSH_PRIVATE_KEY="$MY_KEY" -t dice-server .```

## run
```docker run --name dice-server --network dice-network -e PG_HOST=pg-docker -e PG_PORT=5432 -e PG_NAME=postgres -e PG_USER=postgres -e PG_PWD=docker -e TON_API_HOST=ton-api -e TON_API_PORT=5400 -d -p 8080:9999 dice-server```

## ENV VARS
    * PG_HOST - Postgres host, default value is 'localhost'
    * PG_PORT - Postgres port, default value is `5432`.
    * PG_NAME - Postgres database name, required variable, no default value.
    * PG_USER - Postgres user, required variable, no default value.
    * PG_PWD -  Postgres password, required variable, no default value.
    * RPC_LISTEN_PORT - Listen port for GRPC server, default value is '5400'.
    * WEB_LISTEN_PORT - Listen port for web server, default value is '9999'.
    * WEB_DOMAIN - Domain name used for CORS headers, default value is 'tonbet.io'.
    * TON_API_HOST  - Host of the 'ton-api' service, requred variable, no default value.
    * TON_API_PORT - Port of the 'ton-api' service, default value is '5400'. 