#!/bin/bash

pwd=`pwd`
target=`basename $pwd`
# kill
pid=`ps -C ${target} -o pid=`
if [ -n "$pid" ]; then
  echo "Stopping old version, PID: ${pid}"
  if [ "$1" = "-f" ]; then
    # force shutdown
    echo "Force shutdown..."
    kill -9 $(ps -C ${target} -o pid=)
  else
    kill -s 2 $(ps -C ${target} -o pid=)
  fi
  # wait for program to stop
  pid=`ps -C ${target} -o pid=`
  while [ -n "$pid" ]; do
    sleep 1
  done
fi

# upgrade
if [ -f "${target}-new" ]; then
  echo "Upgrading..."
  if [ -f "${target}-backup" ]; then
    backupdt=`date +%Y%m%d-%H`
    mv "${target}-backup" "${target}-backup-${backupdt}"
  fi

  mv ${target} ${target}-backup
  mv ${target}-new ${target}

  echo "Upgrade Complete"
fi

# run
echo "Starting..."
./run.sh ${target}
echo "Done"
