# Dockerfile
FROM node:18 as builder
WORKDIR /app
RUN yarn install
COPY package.json yarn.lock ./
COPY . .
RUN yarn build

# Serve the static files
FROM nginx:stable-alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
