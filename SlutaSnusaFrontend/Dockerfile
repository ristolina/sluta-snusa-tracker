FROM python:alpine
COPY . /webapp/
EXPOSE 8000
WORKDIR /webapp
RUN pip install -r requirements.txt
CMD ["python",  "manage.py", "runserver", "0.0.0.0:8000"]