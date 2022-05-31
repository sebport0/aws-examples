import logging

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def lambda_handler(event, context):
    logger.info({"state": "PENDING", "failed": False})
    logger.info({"state": "INITIATED", "failed": False})
    logger.info({"state": "RUNNING", "failed": False, "progress": 0})
    logger.info({"state": "RUNNING", "failed": False, "progress": 25})
    logger.info({"state": "RUNNING", "failed": False, "progress": 50})
    logger.error({"state": "TERMINATING", "failed": True, "progress": 75})
    logger.info(
        {
            "state": "TERMINATED",
            "failed": True,
            "progress": 75,
            "message": "Task failed at 75%. Unable to upload images.",
        }
    )
