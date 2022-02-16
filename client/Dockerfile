FROM node:16-alpine as develop-stage
WORKDIR /app
COPY ../client/package*.json ./
RUN yarn install
COPY ../client .

FROM develop-stage as build-stage
RUN yarn build

FROM nginx:alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]