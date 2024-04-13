FROM node:latest

WORKDIR /home/project/

COPY . .

RUN npm install

EXPOSE 5173
EXPOSE 5000
EXPOSE 8080

ENV HOST=0.0.0.0

CMD npm run dev -- --host