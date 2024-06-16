FROM node:20.11.0
WORKDIR /usr/docker/JourniCal/frontend/
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
CMD ["npm", "run", "dev", "--host"]