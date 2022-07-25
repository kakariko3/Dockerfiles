FROM python:3.10.5

# 環境変数を定義（Dockerfileとコンテナ内で参照可）
ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

WORKDIR /code
COPY requirements.txt /code
RUN python -m pip install --upgrade pip && \
    pip install -r requirements.txt

COPY . /code
