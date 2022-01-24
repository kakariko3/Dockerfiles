FROM ruby:2.7.5-alpine

# イメージのビルド時に使用する一時的な変数を設定
ARG RUNTIME_PACKAGES='nodejs tzdata postgresql-dev postgresql git' \
    DEV_PACKAGES='build-base curl-dev'

# 環境変数を設定
ENV LANG=C.UTF-8 \
    TZ=Asia/Tokyo

WORKDIR /app
COPY ./backend/Gemfile* /app

# パッケージのアップデート&インストール
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ${RUNTIME_PACKAGES} && \
    apk add --virtual build-dependencies --no-cache ${DEV_PACKAGES} && \
    bundle install -j4 && \
    apk del build-dependencies

COPY ./backend /app

# コンテナが起動する度に実行されるスクリプトを追加
COPY ./backend/entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
EXPOSE 3000

# 起動時に実行するコマンド
CMD ["rails", "server", "-b", "0.0.0.0"]
