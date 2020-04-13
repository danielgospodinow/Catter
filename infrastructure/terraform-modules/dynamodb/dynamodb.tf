resource "aws_dynamodb_table" "basic-dynamodb-table" {
  name           = "facts"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "id"

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "content"
    type = "S"
  }

  attribute {
    name = "authorId"
    type = "S"
  }
}
