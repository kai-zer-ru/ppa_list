#!/bin/sh
### BEGIN INIT INFO
# Provides:          ppa_list
# Required-Start:    $local_fs $network $named $time $syslog
# Required-Stop:     $local_fs $network $named $time $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Description:       Repository manager
### END INIT INFO

NAME=ppa_list
SCRIPT=/opt/ppalist/ppa_list
RUNAS=root

PIDFILE=/var/run/ppa_list.pid
LOGFILE=/var/log/ppa_list.log

start() {
  if [ -f /var/run/$PIDNAME ] && kill -0 $(cat /var/run/$PIDNAME); then
    echo 'Service ppa_list already running' >&2
    return 1
  fi
  echo 'Starting service ppa_list' >&2
  local CMD="$SCRIPT &> \"$LOGFILE\" & echo \$!"
  su -c "$CMD" $RUNAS > "$PIDFILE"
  echo 'Service ppa_list started' >&2
}

stop() {
  if [ ! -f "$PIDFILE" ] || ! kill -0 $(cat "$PIDFILE"); then
    echo 'Service ppa_list not running' >&2
    return 1
  fi
  echo 'Stopping service ppa_list' >&2
  kill -15 $(cat "$PIDFILE") && rm -f "$PIDFILE"
  echo 'Service ppa_list stopped' >&2
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  retart)
    stop
    start
    ;;
  *)
    echo "Usage: $0 {start|stop|restart}"
esac