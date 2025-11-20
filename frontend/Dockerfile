FROM node:20-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Build untuk production
RUN npm run build

EXPOSE 5173

CMD ["npm", "run", "preview", "--", "--host"]
