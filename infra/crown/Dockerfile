FROM adrianhdez929/crown_devnet:latest

RUN apt-get install --yes python2.7
COPY run.sh /root/run.sh

RUN chmod +x /root/run.sh

EXPOSE 43001

ENTRYPOINT ["./root/run.sh"]
