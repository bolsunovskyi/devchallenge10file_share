FROM node:7

COPY . /root/
WORKDIR /root
RUN npm install
EXPOSE 8080

CMD  ./node_modules/http-server/bin/http-server