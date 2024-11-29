FROM python:3.12-slim

WORKDIR /app

COPY python/requirments.txt .

RUN pip3 install --no-cache-dir -r requirments

COPY python/ .

CMD ["python", "sort_pruducts.py"]