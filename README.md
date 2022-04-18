# GO Example (Web Service)

Trying to implement follow [The Twelve Factor App](https://12factor.net/) and Hexagonal Architecture

## Quickstart guide

### Development
```sh
docker compose -f docker-compose.dev.yml up --build
```

### Deployment
```sh
# host platform image build with compose
docker compose build
docker compose push
# multi platform image build
docker buildx bake -f ./docker-compose.yml --push
```


Up coming...