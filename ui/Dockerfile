FROM node:12-slim as builder

ARG UC_URL

WORKDIR /app

COPY package.json /app/
RUN npm install
COPY . /app
RUN npm run build


FROM nginx:1-alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]