FROM ubuntu:xenial

RUN apt-get update && apt-get install --yes software-properties-common
RUN apt-get install --yes unzip
RUN apt-get install --yes wget

# bitcoind PPA is unmaintained, so we download it from the bitcoincore website.
# This can sometimes stall, which can be fixed by killing and restarting the
# build.

RUN wget -c https://github.com/Crowndev/crown-core/releases/download/v0.14.0.2-Emerald/Crown-0.14.0.2-Linux64.zip && unzip -a Crown-0.14.0.2-Linux64
RUN mv ./Crown-0.14.0.2-Linux64 /app
RUN chmod +x /app/bin/crownd
RUN chmod +x /app/bin/crown-cli

COPY crown.conf /root/.crown/
COPY run.sh /root/
RUN chmod +x /root/run.sh

EXPOSE 9341
ENTRYPOINT ["./root/run.sh"]