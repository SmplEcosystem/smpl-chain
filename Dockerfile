FROM golang:1.18-bullseye as build

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y build-essential ca-certificates curl

RUN curl https://get.ignite.com/cli! | bash

WORKDIR /app

COPY ./ ./

RUN ignite chain build -o /tmp/go/bin
RUN ignite chain init

FROM golang:1.18-bullseye

COPY --from=build /go/bin/* /go/bin/
COPY --from=build /root/.smpl-chain /root/.smpl-chain

CMD [ "/go/bin/smpl-chaind", "start" ]