#!/bin/sh

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------

SUCCESS=0
FAILURE=1

# -----------------------------------------------------------------------------
#  Functions
# -----------------------------------------------------------------------------

echoHeader() {
    echoHR
    echo " ${*}"
    echoHR
}

echoHR() {
    echo '-------------------------------------------------------------------------------'
}

echoN() {
    printf '%s' "$*"
}

runShellCheck() {
    echoN "- Running ShellCheck ... "
    result=$(shfmt -f . | xargs -n 1 shellcheck -a -x --shell=sh 2>&1) || {
        printf "NG\n%s" "$result"

        return "$FAILURE"
    }
    echo 'OK'

    return "$SUCCESS"
}

runShellSpec() {
    echoN "- Running ShellSpec ... "
    result=$(shellspec . 2>&1) || {
        printf "NG\n%s" "$result"

        return "$FAILURE"
    }
    echo 'OK'

    return "$SUCCESS"
}

runGoTestAndCoverage() {
    echoN "- Running Go tests and 100% coverage ... "

    # Run test and coverage
    result=$(go-carpet -256colors -include-vendor . 2>&1) || {
        printf "NG\n%s" "$result"

        return "$FAILURE"
    }

    # Check coverage
    echo "$result" | tail -n 1 | grep '100.0% of statements' >/dev/null || {
        printf "NG\n%s" "$result"

        return "$FAILURE"
    }
    echo 'OK'

    return "$SUCCESS"
}

# -----------------------------------------------------------------------------
#  Main
# -----------------------------------------------------------------------------

# エラー時即時終了
set -eu

runShellCheck
runShellSpec
runGoTestAndCoverage
