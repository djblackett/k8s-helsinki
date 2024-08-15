FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:alpine

WORKDIR /usr/src/app

COPY ./ ./

RUN apk update && apk add --no-cache postgresql-client \
    && chmod +x ./backup.sh

ENV URL=something

CMD [ "./backup.sh" ]
