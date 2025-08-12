data "misc_echo" "example" {
  input = {
    key = "key"
    val = sensitive("secret")
  }
}
