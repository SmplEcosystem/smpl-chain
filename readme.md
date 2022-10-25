# smplchain
**smplchain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite](https://github.com/ignite/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite docs](https://docs.ignite.com/).

### Launch

To launch your blockchain live on multiple nodes, use `ignite network` commands. Learn more about [Starport Network](https://docs.ignite.com/network/introduction).


## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

### Docker
`docker-compose build` will build the latest version

You can start a test chain with `docker-compose up -d`

Make account able to mint token
`docker-compose exec SmplChain  smpl-chaind tx roles add bank smpl1q28v96p6lhyac2ghjlyylsl4290tl722x9kmtg --from alice`

Mint token 
`docker-compose exec SmplChain smpl-chaind tx smplcoins mintsusdse 100 --from alice`