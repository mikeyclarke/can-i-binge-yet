FROM node:19-alpine

RUN corepack enable

WORKDIR /var/app

COPY . .

RUN yarn install && yarn run webpack --config webpack.dev.js
