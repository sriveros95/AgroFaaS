FROM node:10.16-slim as ship

## Todo update versions
ENV APP_ROOT /home/app

RUN apt-get update; \
    apt-get install -y ca-certificates; \
    rm -rf /var/lib/apt/lists/*

## use/enable node user
RUN mkdir -p ${APP_ROOT}

RUN chown node:node -R ${APP_ROOT}
USER node

# Turn down the verbosity to default level.
ENV NPM_CONFIG_LOGLEVEL info

# Wrapper/boot-strapper
WORKDIR ${APP_ROOT}
COPY --chown=node:node app/ ./
WORKDIR ${APP_ROOT}/app

RUN rm -rf node_modules/ ; rm package-lock.json; npm i

RUN npm run build

EXPOSE 3333

CMD ["npm", "start"]