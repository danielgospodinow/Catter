FROM node:12.16.1-stretch

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build

RUN npm install -g serve

CMD ["serve", "-s", "build"]
