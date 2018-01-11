provider "aws" {
  region = "us-east-1"
}

data "aws_subnet" "selected" {
 filter {
    name = "tag:Name"
    values = ["${var.subnet_name}"]
  }
}

// https://www.terraform.io/docs/providers/aws/r/instance.html
resource "aws_instance" "example" {
  ami           = "ami-2757f631"
  instance_type = "t2.micro"
  subnet_id = "${data.aws_subnet.selected.id}"

  tags {
    Name = "test-gostudy-tf-test"
  }
}


//// https://www.terraform.io/docs/providers/aws/r/ecs_cluster.html
//resource "aws_ecs_cluster" "webchat" {
//  name = "test-gostudy-webchat"
//}
//
//resource "aws_ecs_task_definition" "webchat" {
//  family = "webchat"
//  network_mode = "awsvpc"
//  requires_compatibilities = ["FARGATE"]
//  cpu = "256"
//  memory = "512"
//
//  container_definitions = <<DEFINITION
//[
//  {
//    "cpu": 256,
//    "essential": true,
//    "image": "nginx:latest",
//    "memory": 512,
//    "name": "webchatdb",
//    "networkMode": "awsvpc"
//  }
//]
//DEFINITION
//}
//
//resource "aws_ecs_service" "main" {
//  name = "tf-ecs-service-1"
//  cluster = "${aws_ecs_cluster.main.id}"
//  task_definition = "${aws_ecs_task_definition.webchat.arn}"
//  desired_count = 1
//  launch_type = "FARGATE"
//  network_configuration {
//    security_groups = ["${aws_security_group.allow_all_a.id}", "${aws_security_group.allow_all_b.id}"]
//    subnets = ["${aws_subnet.main.*.id}"]
//  }
//}