FROM node:lts-alpine

WORKDIR /app

COPY package*.json /app

RUN yarn add global @vue/cli
RUN yarn install

COPY . .

EXPOSE ${FRONTEND_PORT}

CMD [ "yarn", "serve"]
