#!/usr/bin/env bash

echo -e '\033[0;34m==>\033[0m Packaging microservice in a container...'

docker build -t plush-bot . -f ./Dockerfile

echo -e '\033[0;32m==>\033[0m Done! The image is tagged as plush-bot'

exit 0
