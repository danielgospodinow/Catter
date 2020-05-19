resource "aws_ecr_repository" "catter-facts-service-repository" {
  name = "catter-facts-service"
}

resource "aws_ecr_repository" "catter-rating-service-repository" {
  name = "catter-rating-service"
}

resource "aws_ecr_repository" "catter-account-service-repository" {
  name = "catter-account-service"
}
