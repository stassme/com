# Commands

Use vscode

### Fix vscode permission denied

sudo chown -R $(whoami):vscode /go/pkg/ 

### Systemctl

./grafana-agent-linux-amd64 --config.file=agent-config.yaml

curl -O -L "https://github.com/grafana/agent/releases/latest/download/grafana-agent-linux-amd64.zip" \
 && unzip "grafana-agent-linux-amd64.zip" \
 && chmod a+x grafana-agent-linux-amd64

 api token name
 crm-back-dev

 curl localhost:8080/metrics
 
cat << EOF > /home/tim/agent-config.yaml
metrics:
  global:
    scrape_interval: 30s
  configs:
  - name: hosted-prometheus
    scrape_configs:
      - job_name: crm-back-dev
        static_configs:
        - targets: ['localhost:8080']
    remote_write:
      - url: https://prometheus-prod-24-prod-eu-west-2.grafana.net/api/prom/push
        basic_auth:
          username: 1342595
          password: glc_eyJvIjoiMTAxNTc0NiIsIm4iOiJzdGFjay04MTY5NjctaW50ZWdyYXRpb24tZGV2LWJhY2stbWFpbiIsImsiOiJqMWhzNGgydjN1eDY5eW5tUDE2bTg5dUQiLCJtIjp7InIiOiJwcm9kLWV1LXdlc3QtMiJ9fQ==

traces:
  configs:
  - name: default
    remote_write:
      - endpoint: tempo-prod-10-prod-eu-west-2.grafana.net:443
        basic_auth:
          username: 769066
          password: glc_eyJvIjoiMTAxNTc0NiIsIm4iOiJzdGFjay04MTY5NjctaW50ZWdyYXRpb24tdGVtcG8iLCJrIjoiMTc1QURmYzc1VjA4UGQ4bzdJdDhIbVVNIiwibSI6eyJyIjoicHJvZC1ldS13ZXN0LTIifX0=

    receivers:
      jaeger:
        protocols: 
          thrift_binary:
          thrift_compact:
          thrift_http:
      zipkin:
      otlp:
        protocols:
          http:
          grpc:
      opencensus:

EOF


sudo cat << EOF > /etc/systemd/system/grafana-agent.service
[Unit]
Description=grafana-agent

[Service]
ExecStart=/home/tim/grafana-agent-linux-amd64 --config.file=/home/tim/agent-config.yaml
Restart=always

[Install]
WantedBy=multi-user.target
EOF

systemctl enable grafana-agent.service

systemctl start grafana-agent.service

## Fluent-Bit

sudo nano /etc/fluent-bit/parsers.conf 

[PARSER]
    Name         docker2
    Format       json
    Time_Keep    off
    # --
    # Since Fluent Bit v1.2, if you are parsing Docker logs and using
    # the Kubernetes filter, it's not longer required to decode the
    # 'log' key.
    #
    # Command      |  Decoder | Field | Optional Action
    # =============|==================|=================
    #Decode_Field_As    json     log


sudo nano /etc/fluent-bit/fluent-bit.conf 


[INPUT]
    Name     tail
    Path     /var/lib/docker/containers/*/*.log
    Parser   docker
    Tag      docker.*

[FILTER]
    name          parser
    Match         *
    Parser        docker2
    key_name      log
    Reserve_Data  On
    Preserve_Key  On

[OUTPUT]
    Name         loki
    Match        *
    Host         logs-prod-012.grafana.net
    port         443
    line_format  json
    HTTP_User    770327
    HTTP_Passwd  glc_eyJvIjoiMTAxNTc0NiIsIm4iOiJmbHVlbnRiaXQtZmx1ZW50Yml0IiwiayI6ImZlNzI2TDY1N1oxWXdXd1JGQjNiNkFXOSIsIm0iOnsiciI6InVzIn19
    Labels       job=fluentbit
    tls         on
    tls.verify  on



### Cli

git clone https://github.com/grafana/loki.git
cd loki
make logcli
cp cmd/logcli/logcli /usr/local/bin/logcli

...
export LOKI_ADDR=https://logs-prod-012.grafana.net
export LOKI_USERNAME=770327
export LOKI_PASSWORD=glc_eyJvIjoiMTAxNTc0NiIsIm4iOiJmbHVlbnRiaXQtZmx1ZW50Yml0IiwiayI6ImZlNzI2TDY1N1oxWXdXd1JGQjNiNkFXOSIsIm0iOnsiciI6InVzIn19
...
eval "$(logcli --completion-script-bash)"
eval "$(logcli --completion-script-zsh)"

logcli --addr=https://logs-prod-012.grafana.net --username=770327 --password=glc_eyJvIjoiMTAxNTc0NiIsIm4iOiJsb2tpLWNsaS1sb2tpLWNsaSIsImsiOiJ0TTkyOEo3MHRpbzBpaWhZRzUzbjB2RTUiLCJtIjp7InIiOiJwcm9kLWV1LXdlc3QtMiJ9fQ==  query '{job="fluentbit"}'

## DB

Max parallel workers per gather = 0
 
## DB

SELECT cron.schedule('* * * * *' , 'VACUUM');
SELECT cron.schedule('30 * * * *' , 'VACUUM ANALYZE');
SELECT cron.schedule('30 3 * * *' , 'VACUUM FULL');
SELECT * from cron.job;
SELECT * from cron.job_run_details;
-- SELECT cron.unschedule(3);

## Atlas
 
atlas schema inspect -u "postgres://default:secret@postgres:5432/main?sslmode=disable" --format '{{ sql . }}' > schema.sql
atlas schema inspect -u "postgres://default:secret@postgres:5432/main?sslmode=disable" > schema.hcl


atlas schema apply \
  -u "postgres://default:secret@postgres:5432/main?sslmode=disable" \
  --dev-url "postgres://default:secret@postgres:5432/temp?sslmode=disable" \
  --to file://schema.sql

atlas schema inspect \
  -u "postgres://default:secret@postgres:5432/main?sslmode=disable" \
  --web

  atlas migrate diff initial  \
     --dir "file://migrations2" \
    --to "file://schema.sql" \
    --format '{{ sql . }}'  \
    --dev-url "postgres://default:secret@postgres:5432/temp?sslmode=disable"
 


  atlas migrate diff  initial  \
     --dir "file://migrations2" \
    --to "file://schema.sql" \
    --format '{{ sql . }}'  \
    --dev-url "postgres://default:secret@postgres:5432/temp?sslmode=disable"
 
 atlas schema apply \
  -u "postgres://default:secret@postgres:5432/main?sslmode=disable" \
  --dev-url "postgres://default:secret@postgres:5432/temp?sslmode=disable" \
  --to "file://migrations2"  
 

  