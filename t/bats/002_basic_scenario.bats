setup() {
  load 'test_helper/bats-support/load'
  load 'test_helper/bats-assert/load'
}

@test "ensure machines starts up" {
  run multipass-compose up

  [ "$status" -eq 0 ]
}

@test "ensure status shown" {
  run multipass-compose status

  assert_output --partial 'Name:           web-server'
  assert_output --partial 'Name:           database'
  assert_output --partial 'Name:           backend'
}

@test "ensure machines tears down" {
  run multipass-compose down

  [ "$status" -eq 0 ]
}
