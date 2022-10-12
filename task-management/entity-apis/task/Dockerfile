FROM node:18-alpine

# Create app directory
WORKDIR /usr/src/app
RUN npm install -g typescript ts-node
# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./
COPY tsconfig.json ./


RUN npm install

# Bundle app source
COPY . .

# Set a non-root user
USER 10014

EXPOSE 8080

CMD [ "ts-node", "src/index.ts" ]

