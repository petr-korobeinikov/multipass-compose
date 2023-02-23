setup() {
  load 'test_helper/bats-support/load'
  load 'test_helper/bats-assert/load'

  cd "$(mktemp -d)" || exit
}

@test "ensure default spec file matches expected" {
  run multipass-compose init
  [ "$status" -eq 0 ]

  run diff <(yq -P 'sort_keys(..)' "${BATS_TEST_DIRNAME}/testdata/003_expected_multipass_compose.yaml") <(yq -P 'sort_keys(..)' multipass-compose.yaml)
  [ "$status" -eq 0 ]
}
