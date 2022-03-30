APPS_FOLDER=/var/lib
APP_NAME=food
APP_VERSION=v0.1.0;
DOMAIN=${DOMAIN:?"DOMAIN is required as env variable (e.g. mondomain.com)"}

apt install awscli sqlite3

mkdir -p ${APPS_FOLDER}/${APP_NAME}/${APP_VERSION} ${APPS_FOLDER}/${APP_NAME}/tmp/sessions ${APPS_FOLDER}/${APP_NAME}/tmp/uploads;

chown -R app:app ${APPS_FOLDER}/${APP_NAME}

cat <<EOF | curl -X POST -H 'Content-Type: application/json' -d@- localhost:2019/config/apps/http/servers/srv0/routes/
{
  "@id": "food",
  "match": [{ "host": ["${DOMAIN}"] }],
  "handle": [
    {
      "handler": "headers",
      "response": {
        "set": {
          "Referrer-Policy": ["same-origin"],
          "X-Content-Type-Options": ["nosniff"],
          "X-Frame-Options": ["DENY"],
          "X-Xss-Protection": ["1; mode=block"]
        }
      }
    },
    {
      "handler": "reverse_proxy",
      "upstreams": [{"dial": "localhost:11001"}]
    }
  ]
}
EOF

echo "Copy/Paste env variables. Press <C-d> when done"
cat > ${APPS_FOLDER}/${APP_NAME}/.env;

cat <<CODE > /usr/bin/food
#!/usr/bin/env bash

set -o errexit
set -o pipefail

source ${APPS_FOLDER}/${APP_NAME}/.env

case "\$1" in
  start)
    ${APPS_FOLDER}/${APP_NAME}/food
    ;;
  dump)
    cat <<EOF | sqlite3 ${APPS_FOLDER}/${APP_NAME}/food.sqlite | bzip2 --best | env AWS_ACCESS_KEY_ID=\${FOOD_BACKUP_AWS_ACCESS_KEY_ID} AWS_SECRET_ACCESS_KEY=\${FOOD_BACKUP_AWS_SECRET_ACCESS_KEY} AWS_REGION=\${FOOD_BACKUP_AWS_REGION} aws s3 cp - s3://\${FOOD_BACKUP_AWS_BUCKET}/dumps/food/\$(date '+%Y-%m-%d').sql.bzip
.output stdout
.dump
.exit
EOF
    ;;
  *)
    >&2 echo "command does not exist. possible commands: start, dump"
    exit 1;
esac
CODE

chmod u+x /usr/bin/food
chown app:app /usr/bin/food

cat <<EOF > /etc/systemd/system/food.service
[Unit]
Description=Sport
Documentation=https://github.com/lonepeon/food
After=network.target network-online.target
Requires=network-online.target

[Service]
Type=simple
User=app
Group=app
ExecStart=/usr/bin/food start
Restart=always
TimeoutStopSec=5s
LimitNOFILE=1048576
LimitNPROC=512
PrivateTmp=true
ProtectSystem=full

[Install]
WantedBy=multi-user.target
EOF

cat <<EOF > /etc/systemd/system/food-backup.service
[Unit]
Description=Sport Backup
Wants=food-backup.timer
After=network.target network-online.target
Requires=network-online.target

[Service]
Type=oneshot
User=app
Group=app
ExecStart=/usr/bin/food dump

[Install]
WantedBy=multi-user.target
EOF

cat <<EOF > /etc/systemd/system/food-backup.timer
[Unit]
Description=Dump food database every week
Requires=food-backup.service

[Timer]
Unit=food-backup.service
OnCalendar=Sun *-*-* 14:00:00

[Install]
WantedBy=timers.target
EOF


systemctl enable --now food-backup.service
systemctl enable --now food-backup.timer
systemctl enable --now food.service
