terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  backend "s3" {
    bucket = "maidsafe-org-infra-tfstate"
    key    = "terraform-ant-supply.tfstate"
    region = "eu-west-2"
  }
}

resource "aws_cloudwatch_log_group" "ant_supply" {
  name = "/ecs/${terraform.workspace}-${var.ant_supply_log_group_name}"
}

resource "aws_cloudwatch_log_resource_policy" "ant_supply" {
  policy_name = "${terraform.workspace}-${var.ant_supply_log_group_name}"

  policy_document = <<CONFIG
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "es.amazonaws.com"
      },
      "Action": [
        "logs:PutLogEvents",
        "logs:PutLogEventsBatch",
        "logs:CreateLogStream"
      ],
      "Resource": "arn:aws:logs:*"
    }
  ]
}
CONFIG
}

module "vpc" {
  source               = "terraform-aws-modules/vpc/aws"
  version              = "5.19.0"
  name                 = "${terraform.workspace}-${var.ant_supply_vpc_name}"
  cidr                 = "10.0.0.0/16"
  azs                  = var.availability_zones
  public_subnets       = ["10.0.0.0/24", "10.0.1.0/24"]
  private_subnets      = ["10.0.2.0/24"]
  enable_nat_gateway   = true
  single_nat_gateway   = true
  enable_dns_hostnames = true
  enable_dns_support   = true
}

resource "aws_security_group" "ant_supply" {
  name        = "ant_supply-sg"
  description = "Allow inbound HTTP and HTTPS"
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "ant_supply_ecs_tasks" {
  name        = "ant_supply_ecs_tasks-sg"
  description = "Allow ALB to communicate with ECS tasks"
  vpc_id      = module.vpc.vpc_id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group_rule" "ecs_ingress_from_alb" {
  type                     = "ingress"
  from_port                = 3000
  to_port                  = 3000
  protocol                 = "tcp"
  security_group_id        = aws_security_group.ant_supply_ecs_tasks.id
  source_security_group_id = aws_security_group.ant_supply.id
}

resource "aws_lb" "ant_supply" {
  name               = "ant-supply-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.ant_supply.id]
  subnets            = module.vpc.public_subnets
}

resource "aws_lb_target_group" "ant_supply" {
  name     = "ant-supply-tg"
  port     = 3000
  protocol = "HTTP"
  vpc_id   = module.vpc.vpc_id
  target_type = "ip"

  health_check {
    path                = "/"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }
}

resource "aws_lb_listener" "ant_supply" {
  load_balancer_arn = aws_lb.ant_supply.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.ant_supply.arn
  }
}

resource "aws_ecs_cluster" "ant_supply" {
  name = "${terraform.workspace}-${var.ant_supply_ecs_cluster_name}"
}

resource "aws_ecs_task_definition" "ant_supply" {
  family                   = "${terraform.workspace}-ant-supply"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 2048
  memory                   = 4096
  execution_role_arn       = var.ecs_ant_supply_iam_role_arn
  task_role_arn            = var.ecs_ant_supply_iam_role_arn
  container_definitions    = <<TASK_DEFINITION
[
    {
        "name": "${terraform.workspace}-ant-supply",
        "image": "jacderida/ant-circulating-supply:latest",
        "cpu": 0,
        "memoryReservation": 2048,
        "portMappings": [
            {
                "containerPort": 3000,
                "hostPort": 3000,
                "protocol": "tcp"
            }
        ],
        "essential": true,
        "entryPoint": [],
        "command": [],
        "environment": [],
        "volumesFrom": [],
        "logConfiguration": {
            "logDriver": "awslogs",
            "options": {
                "awslogs-group": "/ecs/${terraform.workspace}-${var.ant_supply_log_group_name}",
                "awslogs-region": "eu-west-2",
                "awslogs-stream-prefix": "ecs"
            }
        }
    }
]
TASK_DEFINITION

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }
}

resource "aws_ecs_service" "ant_supply" {
  name            = "${terraform.workspace}-ant-supply"
  cluster         = aws_ecs_cluster.ant_supply.id
  task_definition = aws_ecs_task_definition.ant_supply.arn
  launch_type     = "FARGATE"

  network_configuration {
    subnets         = module.vpc.public_subnets
    security_groups = [aws_security_group.ant_supply_ecs_tasks.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.ant_supply.arn
    container_name   = "${terraform.workspace}-ant-supply"
    container_port   = 3000
  }

  desired_count = 2
}
