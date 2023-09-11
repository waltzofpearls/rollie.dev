FROM node:18.17-bookworm-slim

WORKDIR /frontend

COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

CMD npm run dev -- --host 0.0.0.0 --port $PORT
