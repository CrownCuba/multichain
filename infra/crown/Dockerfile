FROM ubuntu:xenial

RUN apt-get update --fix-missing && apt-get install --yes software-properties-common
RUN apt-get install --yes curl python2.7 wget unzip git
RUN wget https://github.com/Crowndev/crown-core/releases/download/v0.14.0.2-Emerald/Crown-0.14.0.2-Linux64.zip
RUN unzip -a Crown-0.14.0.2-Linux64
RUN git clone https://gitlab.crownplatform.com/crown/crown-core.git
RUN mv ./Crown-0.14.0.2-Linux64 /app
RUN chmod +x /app/bin/crownd
RUN chmod +x /app/bin/crown-cli
RUN mkdir /app/sandbox

COPY crown.conf /root/.crown/
COPY run.sh /root/
RUN chmod +x /root/run.sh
RUN /root/run.sh prepare
RUN /root/run.sh start

EXPOSE 18443

ENTRYPOINT ["./root/run.sh sim"]