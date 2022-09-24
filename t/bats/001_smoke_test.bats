setup() {
  load 'test_helper/bats-support/load'
  load 'test_helper/bats-assert/load'
}

@test "ensure multipass-compose binary exacutable" {
  run multipass-compose

  assert_output --partial 'NAME:'
  assert_output --partial 'multipass-compose'
  assert_output --partial 'USAGE:'
  assert_output --partial 'COMMANDS:'
}
