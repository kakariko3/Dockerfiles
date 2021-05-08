FROM ruby:2.5

# 必要なパッケージのインストール
RUN apt-get update -qq && \
    apt-get install -y nodejs

WORKDIR /myapp
COPY Gemfile /myapp
COPY Gemfile.lock /myapp
RUN bundle install
COPY . /myapp

# コンテナが起動する度に実行されるスクリプトを追加
COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
EXPOSE 3000

# 起動時に実行するコマンド
CMD ["rails", "server", "-b", "0.0.0.0"]
