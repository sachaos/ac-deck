FROM python:3.4.3-slim

RUN apt-get update -y && apt-get upgrade -y && apt-get install -y build-essential libblas-dev liblapack-dev gfortran
RUN pip install numpy==1.8.2
RUN pip install scipy==0.13.3

CMD ["python"]
