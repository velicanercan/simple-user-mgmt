FROM node:lts-alpine

WORKDIR /app

COPY package*.json /app

RUN npm ci

COPY . .
RUN npm run build

EXPOSE ${FRONTEND_PORT}

CMD ["npm", "run", "serve"]
