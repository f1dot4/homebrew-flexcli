#!/bin/sh
# Writes a crontab from FLEXCLI_IMPORT_SCHEDULE and starts crond.
# Logs are forwarded to stdout/stderr for Docker to capture.
set -e

SCHEDULE="${FLEXCLI_IMPORT_SCHEDULE:-0 */2 * * *}"

echo "Scheduling activity import: $SCHEDULE"

# Forward cron output to container stdout/stderr
echo "$SCHEDULE /usr/local/bin/import-activities >> /proc/1/fd/1 2>> /proc/1/fd/2" | crontab -

# Run once immediately on startup so the first import doesn't wait until 3am
echo "Running initial import..."
/usr/local/bin/import-activities || echo "Initial import failed, will retry on schedule"

exec crond -f -l 2
