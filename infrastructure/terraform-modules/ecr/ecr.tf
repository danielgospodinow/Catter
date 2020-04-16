resource "aws_ecr_repository" "container-repository" {
  name = "catter-container-repository"

  image_scanning_configuration {
    scan_on_push = true
  }
}
