#!/bin/bash

aws lambda invoke \
    --function-name lambda-structured-logs-progressLogger \
    progress_logger_response.json