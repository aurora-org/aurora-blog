#!/bin/sh

CONFIG_FILE=./config/deploy.config
LOG_FILE=./aurora.log

nohup ./aurora -conf ${CONFIG_FILE} > ${LOG_FILE}