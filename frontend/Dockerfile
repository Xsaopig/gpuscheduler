# Dockerfile
FROM node:18 as builder
WORKDIR /app
# COPY package.json yarn.lock ./
COPY package.json ./
RUN yarn install
COPY . .
RUN yarn build

# Serve the static files
FROM nginx:stable-alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
