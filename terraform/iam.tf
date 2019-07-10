data "aws_iam_policy_document" "kms_parameter_store" {
  statement {
    sid = "Enable IAM User Permissions"
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
    }
    actions = ["kms:*"]
    resources = ["*"]
  }
  statement {
    sid = "Allow user of the key"
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/waltzofpearls"]
    }
    actions = ["kms:Decrypt"]
    resources = ["*"]
  }
}

resource "aws_kms_key" "kms_parameter_store" {
  description             = "Parameter store kms master key"
  policy                  = "${data.aws_iam_policy_document.kms_parameter_store.json}"
  deletion_window_in_days = 10
  enable_key_rotation     = true
}

resource "aws_kms_alias" "parameter_store_alias" {
  name          = "alias/parameter_store_key"
  target_key_id = "${aws_kms_key.kms_parameter_store.id}"
}

data "aws_iam_policy_document" "ecs_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ecs_execution_role" {
  name               = "app-ecs-execution-role"
  assume_role_policy = "${data.aws_iam_policy_document.ecs_assume_role.json}"
}

resource "aws_iam_policy_attachment" "ecs_execution_role" {
  name       = "app-ecs-execution-role"
  roles      = ["${aws_iam_role.ecs_execution_role.name}"]
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

data "aws_iam_policy_document" "parameter_store" {
  statement {
    actions = [
      "ssm:GetParameters",
      "kms:Decrypt"
    ]
    resources = [
      "arn:aws:ssm:${var.aws_region}:${data.aws_caller_identity.current.account_id}:parameter/rolli3net/*",
      "${aws_kms_key.kms_parameter_store.arn}"
    ]
  }
}

resource "aws_iam_policy" "parameter_store" {
  name   = "app-parameter-store"
  policy = "${data.aws_iam_policy_document.parameter_store.json}"
}


resource "aws_iam_policy_attachment" "parameter_store" {
  name       = "app-parameter-store"
  roles      = ["${aws_iam_role.ecs_execution_role.name}"]
  policy_arn = "${aws_iam_policy.parameter_store.arn}"
}
