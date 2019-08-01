# Set up CloudWatch group and log stream and retain logs for 30 days
resource "aws_cloudwatch_log_group" "log_group" {
  name              = "/ecs/rollie-dev"
  retention_in_days = 30

  tags = {
    Name = "rollie-dev-log-group"
  }
}

resource "aws_cloudwatch_log_stream" "log_stream" {
  name           = "rollie-dev-log-stream"
  log_group_name = aws_cloudwatch_log_group.log_group.name
}
