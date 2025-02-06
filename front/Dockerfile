ARG NODE_VERSION=22.8.0

FROM node:${NODE_VERSION}-slim as base

ARG PORT=3003

WORKDIR /app

# Build
FROM base as build

COPY --link package.json ./
RUN npm install

COPY --link . .

RUN npm run build

#run
FROM base

ENV PORT=$PORT
ENV NODE_ENV=production

COPY --from=build /app/.output /app/.output
# Optional, only needed if you rely on unbundled dependencies
# COPY --from=build /src/node_modules /src/node_modules

CMD [ "node", ".output/server/index.mjs" ]