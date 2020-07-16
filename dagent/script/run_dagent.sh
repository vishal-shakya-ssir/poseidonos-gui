#!/bin/bash
currdir=$(readlink -f $(dirname $0))
cd $currdir
parentdir="$(dirname "$currdir")"
x=`systemctl is-active dagent.service`

if [ $x = "active" ]
then
  echo DAgent service is already running!
else
	#create a soft link of the directory, to get the absolute path for starting the service
	sudo rm /usr/local/dagent
	sudo ln -s $parentdir /usr/local

	#move the service file  to the /etc/systemd/system/
	sudo cp $PWD/dagent.service /etc/systemd/system
	sudo chmod 664 /etc/systemd/system/dagent.service

	sudo systemctl daemon-reload
	sudo systemctl enable dagent.service
	sudo systemctl start dagent

	echo "Starting DAgent...."
	sleep 8s
	echo "Checking Status..."

	x=`systemctl is-active dagent.service`
	if [ $x = "active" ]
	then
  	echo DAgent Service Started Succesfully
	else
		sleep 90s
		if [ $x = "active" ]
		then
			echo DAgent Service Started Successfully
		else
                        echo DAgent Service Failed to Start!
		fi
	fi
fi
