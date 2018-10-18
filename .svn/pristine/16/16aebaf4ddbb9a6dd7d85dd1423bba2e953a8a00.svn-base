#!/bin/bash

case $1 in 
	start)
		nohup ./wadmin 2>&1 >> info.log 2>&1 /dev/null &
		echo "wadmin started..."
		sleep 1
	;;
	stop)
		killall wadmin
		echo "wadmin stoped..."
		sleep 1
	;;
	restart)
		killall wadmin
		sleep 1
		nohup ./wadmin 2>&1 >> info.log 2>&1 /dev/null &
		echo "wadmin restarted..."
		sleep 1
	;;
	*) 
		echo "$0 {start|stop|restart}"
		exit 4
	;;
esac
