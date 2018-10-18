#!/bin/bash

case $1 in 
	start)
		nohup ./wapi 2>&1 >> info.log 2>&1 /dev/null &
		echo "wapi started..."
		sleep 1
	;;
	stop)
		killall wapi
		echo "wapi stoped..."
		sleep 1
	;;
	restart)
		killall wapi
		sleep 1
		nohup ./wapi 2>&1 >> info.log 2>&1 /dev/null &
		echo "wapi restarted..."
		sleep 1
	;;
	*) 
		echo "$0 {start|stop|restart}"
		exit 4
	;;
esac
