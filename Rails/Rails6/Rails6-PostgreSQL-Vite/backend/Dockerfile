FROM ruby:2.7.5-alpine

# Dockerfile内で使用する変数を定義
ARG RUNTIME_PACKAGES='nodejs tzdata postgresql-dev postgresql git' \
    DEV_PACKAGES='build-base curl-dev'

# 環境変数を定義（Dockerfileとコンテナ内で参照可）
ENV LANG=C.UTF-8 \
    TZ=Asia/Tokyo

WORKDIR /app
COPY Gemfile* /app

# パッケージのアップデート&インストール
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ${RUNTIME_PACKAGES} && \
    apk add --virtual build-dependencies --no-cache ${DEV_PACKAGES} && \
    bundle install -j4 && \
    apk del build-dependencies

COPY . /app

# コンテナ起動時に実行するコマンド
CMD ["rails", "server", "-b", "0.0.0.0"]
