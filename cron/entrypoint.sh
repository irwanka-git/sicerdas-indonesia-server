#!/bin/sh
# start cron
chmod 755 /home/script.sh
chmod 755 /app/utils/skoring/main
echo "Cronjob Register....!"
/usr/bin/crontab /home/crontab.txt
echo "Cronjob Starter....!"
/usr/sbin/crond -f -l 8

