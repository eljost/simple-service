# simple-service

Simple demo golang http server that can be built as docker image and pushed to the
Github container registry.

```bash
nix run .#push-docker-image
```

A docker image is built and pushed to the Github container registry, everytime a new
tag is pushed to Github.