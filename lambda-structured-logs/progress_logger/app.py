import logging

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def lambda_handler(event, context):
    logger.info({"state": "PENDING", "failed": False})
    logger.info({"state": "INITIATED", "failed": False})
    logger.info({"state": "RUNNING", "failed": False, "progress": {"at": 0, "remaining": 100}})
    logger.info({"state": "RUNNING", "failed": False, "progress": {"at": 25, "remaining": 75}})
    logger.info({"state": "RUNNING", "failed": False, "progress": {"at": 50, "remaining": 50}})
    logger.error({"state": "TERMINATING", "failed": True, "progress": {"at": 75, "remaining": 25}})
    logger.info(
        {
            "state": "TERMINATED",
            "failed": True,
            "progress": {"at": 75, "remaining": 25},
            "message": "Task failed at 75%. Unable to upload images.",
        }
    )
