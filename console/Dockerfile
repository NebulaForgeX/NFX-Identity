# syntax=docker/dockerfile:1
FROM node:22-alpine AS deps
WORKDIR /app
ARG VITE_API_URL=http://192.168.1.64/nfx-identity
ARG VITE_WS_URL=ws://192.168.1.64/nfx-identity
ARG VITE_BUILD_ENV=prod
ARG VITE_APP_ID=nfx-console
ENV VITE_API_URL=${VITE_API_URL}
ENV VITE_WS_URL=${VITE_WS_URL}
ENV VITE_BUILD_ENV=${VITE_BUILD_ENV}
ENV VITE_APP_ID=${VITE_APP_ID}
COPY package.json package-lock.json ./
RUN npm ci
COPY . .

FROM deps AS dev
ENV DOCKER=1
EXPOSE 5173
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]

FROM deps AS builder
ENV DOCKER_BUILD=1
RUN npm run build

FROM nginx:alpine AS prod
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
