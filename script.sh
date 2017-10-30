#!/bin/bash

case "$1" in
	start)
		nohup ./fizzbuzz &
		;;

	stop)
		pid=$(ps -ef | grep fizzbuzz | grep -v grep | awk '{print $2}' | xargs kill)
		;;
	*)
		echo "Usage: ./script.sh start|stop"
	esac
