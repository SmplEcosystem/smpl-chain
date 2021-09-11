FROM golang:1.16

RUN adduser --system --disabled-password appuser
RUN curl https://get.starport.network/starport! | bash
#RUN git clone https://github.com/Smpl-Finance/smpl-chain.git

WORKDIR /home/appuser
COPY . .
RUN chown -R appuser .

USER appuser

RUN starport chain build

RUN smpl-chaind init devNode

COPY smpl-testnet.genesis.json /home/appuser/.smpl-chain/config/genesis.json

CMD ["smpl-chaind", "start", "--p2p.seeds", "2c6c0ec4d7576ef5061741b51416f1611f40b8d3@smpl-validator2.s56.net:26656"]