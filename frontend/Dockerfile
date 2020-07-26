# stage 0: run npm build
FROM node:latest as builder

# creating building workdir
WORKDIR /usr/src/app

# copy depedencies to workdir
COPY package*.json ./
# install necessary packages
RUN npm i --production

# copy angular project code to workdir
COPY . .
# build angular project

# ARG CONFIG=production

RUN npm run build -- --output-path=./dist/out
# --configuration $CONFIG --source-map=false --aot=false

# stage 1: run nginx web server

# load nginx latest image
FROM nginx:latest

# copy nginx config and certificate
COPY nginx /etc/nginx/

# create and copy ng output to nginx html directoy
WORKDIR /usr/share/nginx/html
COPY --from=builder /usr/src/app/dist/out /usr/share/nginx/html
