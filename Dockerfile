FROM python:3.7

RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app/

COPY . /usr/src/app/

# with bash
CMD ["python", "app.py"]
# without bash
ENTRYPOINT ["python", "app.py"]