FROM golang:1.18-bullseye as build

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y build-essential ca-certificates curl

RUN curl https://get.starport.network/starport! | bash

WORKDIR /app

COPY ./ ./

RUN starport chain build -o /go/bin
RUN starport chain init

RUN ls -l /go/bin

FROM golang:1.18-bullseye

COPY --from=build /go/bin/* /go/bin/
COPY --from=build /root/.smpl-chain /root/.smpl-chain


CMD [ "/go/bin/smpl-chaind", "start" ]