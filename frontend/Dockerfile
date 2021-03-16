FROM node:alpine as builder
RUN mkdir /app
ADD ./frontend /app/
WORKDIR /app
RUN yarn install && yarn build

FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
COPY --from=builder /app/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80