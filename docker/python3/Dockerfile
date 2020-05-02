FROM python:3.8.2

RUN apt-get update -y && apt-get upgrade -y && apt-get install -y build-essential libblas-dev liblapack-dev gfortran
RUN pip install numpy==1.18.2
RUN pip install scipy==1.4.1
RUN pip install numba==0.48.0
RUN pip install networkx==2.4
RUN pip install sklearn==0.22.2.post1

CMD ["python"]
