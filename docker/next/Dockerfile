FROM node:20-slim
WORKDIR /app
# install deps first for cache
COPY ./next/package.json ./next/pnpm-lock.yaml ./
RUN corepack enable && pnpm install --frozen-lockfile
# copy source
COPY ./next .
EXPOSE 3000
CMD ["pnpm", "dev", "-H", "0.0.0.0"]