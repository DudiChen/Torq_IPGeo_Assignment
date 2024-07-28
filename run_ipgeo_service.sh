#!/bin/bash

WORK_DIR=.
#CSV_DB_PATH=$WORK_DIR/csv/data.csv
APP_NAME=ipgeo

cd $WORK_DIR

sudo apt update -y
sudo apt upgrade -y

sudo docker build -t $APP_NAME .

sudo docker run -it $APP_NAME