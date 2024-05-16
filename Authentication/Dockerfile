FROM python:3.8-bullseye
WORKDIR /app
COPY . /app
RUN pip install -r requirements.txt --no-cache-dir

EXPOSE 5000

CMD ["flask", "--app" ,"src" ,"run", "--debug","--host","0.0.0.0"]