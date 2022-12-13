FROM node:18.1-alpine
 
CMD ["/bin/sh"]
 
WORKDIR /app
 
RUN apk update && \
    npm install vue@latest
 
ADD ./front/package.json /app
 
RUN npm install