import logging
import random

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def lambda_handler(event, context):
    for i in range(1, 11):
        latency = random.choice((10, 15, 25, 30, 50, 999))
        logger.info({"packet": i, "latency": latency})
