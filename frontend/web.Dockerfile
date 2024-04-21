# Lighter Node.js base image
FROM node:lts-alpine AS build-stage

WORKDIR /app

COPY package*.json ./

RUN npm ci
COPY . .
RUN npm run build

# Production stage
FROM nginx:stable-alpine AS production-stage

COPY --from=build-stage /app/dist /usr/share/nginx/html

EXPOSE ${FRONTEND_PORT}

CMD ["nginx", "-g", "daemon off;"]
