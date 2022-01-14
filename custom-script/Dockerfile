FROM python:3.9

WORKDIR /app

COPY . .

RUN pip3 install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple

CMD ["gunicorn",  "app:app", "-c", "gunicorn.conf.py"]