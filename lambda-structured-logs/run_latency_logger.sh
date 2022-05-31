#!/bin/bash

aws lambda invoke \
    --function-name lambda-structured-logs-latencyLogger \
    latency_logger_response.json
