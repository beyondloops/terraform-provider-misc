ephemeral "utilities_preserve_sensitivity" "example" {
  input = {
    key = "key"
    val = sensitive("secret")
  }
}
