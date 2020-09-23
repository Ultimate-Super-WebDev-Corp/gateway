FROM docker.elastic.co/logstash/logstash:7.9.1

RUN mkdir -p /usr/share/logstash/logstash-core/lib/jars
RUN curl  https://jdbc.postgresql.org/download/postgresql-42.2.16.jre7.jar -o /usr/share/logstash/logstash-core/lib/jars/postgresql.jar
RUN /usr/share/logstash/bin/logstash-plugin install logstash-output-jdbc
