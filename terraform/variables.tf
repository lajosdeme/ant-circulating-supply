variable "ant_supply_vpc_name" {
  default = "ant_supply"
  description = "The name of the ant supply VPC"
}

variable "availability_zones" {
  default = ["eu-west-2a", "eu-west-2b", "eu-west-2c"]
  description = "The availability zones to use"
}

variable "ant_supply_log_group_name" {
  default = "ant-supply"
  description = "The name of the ant supply logging group"
}

variable "ant_supply_ecs_cluster_name" {
  default = "ant-supply"
  description = "The name of the ECS cluster"
}

# At the moment this is using a role I had already defined with the necessary policies.
# A new role specific to this deployment can be created later.
variable "ecs_ant_supply_iam_role_arn" {
  default = "arn:aws:iam::389640522532:role/ecs_testnet_infra"
  description = "The ARN of the role for executing the Elastic-related services on ECS"
}
