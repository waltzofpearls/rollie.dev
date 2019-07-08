variable "aws_region" { default = "us-west-2" }

variable "lb_port" { default = "80" }
variable "app_image" { default = "waltzofpearls/rolli3.net:latest" }
variable "app_port" { default = 3000 }
variable "health_check_path" { default = "/" }

# Fargate instance CPU units to provision (1 vCPU = 1024 CPU units)
# Fargate instance memory to provision (in MiB)
variable "fargate_cpu" { default = "256" }
variable "fargate_memory" { default = "512" }

variable "env_name" { default = "production" }
