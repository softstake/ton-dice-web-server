FROM golang:1.14-alpine as builder

ARG GITHUB_TOKEN
ARG WEB_LISTEN_PORT=9999
ARG RPC_LISTEN_PORT=5300

# RUN mkdir /root/.ssh/
# RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
# RUN chmod 400 /root/.ssh/id_rsa

# RUN apk add --no-cache openssh
RUN apk add --no-cache git

# RUN touch /root/.ssh/known_hosts
# RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

# RUN git clone git@github.com:tonradar/ton-api.git
RUN git config --global url."https://${GITHUB_TOKEN}:@github.com/".insteadOf "https://github.com/"

RUN git clone https://github.com/tonradar/ton-api.git

WORKDIR /go/src/build
ADD . .
# RUN git config --global url."git@github.com:".insteadOf "https://github.com/"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dice-server ./cmd


FROM scratch
COPY --from=builder /go/src/build/dice-server /app/
WORKDIR /app
EXPOSE $WEB_LISTEN_PORT $RPC_LISTEN_PORT

CMD ["./dice-server"]