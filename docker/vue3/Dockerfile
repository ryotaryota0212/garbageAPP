FROM node:16.15.1-alpine
 
CMD ["/bin/sh"]
 
WORKDIR /app
 
RUN apk update 
 
ADD ./front/package.json /app
 
RUN npm install