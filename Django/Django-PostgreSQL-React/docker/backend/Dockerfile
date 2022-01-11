FROM python:3.7

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

WORKDIR /app
COPY ./backend/requirements.txt /app
RUN python -m pip install --upgrade pip
RUN pip install -r requirements.txt
COPY ./backend /app
