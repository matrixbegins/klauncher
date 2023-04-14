FROM busybox

WORKDIR /app
COPY ./finite_counter.sh .

CMD [ "sh", "finite_counter.sh"]
