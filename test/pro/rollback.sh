#!/bin/bash

pwd=`pwd`
target=`basename $pwd`

# Kill running program
pid=`ps -C ${target} -o pid=`
if [ -n "$pid" ]; then

  echo "Stopping old version, PID: ${pid}"
  if [ "$1" = "-f" ]; then
    # force shutdown
    echo "Force shutdown..."
    kill $(ps -C ${target} -o pid=)
  else
    kill -s 2 $(ps -C ${target} -o pid=)
  fi

  # wait for program to stop
  pid=`ps -C ${target} -o pid=`
  while [ -n "$pid" ]; do
    sleep 1
  done
fi

# Rollback
if [ -f "${target}-backup" ]; then
  echo "Rolling back..."
  if [ -f "${target}" ]; then
    rm "${target}"
  fi
  mv ${target}-backup ${target}
  echo "Rollback Complete"
fi

# run
echo "Starting..."
./run.sh ${target}
echo "Done"
