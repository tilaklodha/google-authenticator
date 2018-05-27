## Google Authenticator
A small go program to generate the google authenticator code.

- Install go
    - On OSX run `brew install go`.
    - Follow instructions on https://golang.org/doc/install for other OSes.

- Setup go
    - Make sure that the executable `go` is in your shell's path.
    - Add the following in your .zshrc or .bashrc: (where `<workspace_dir>` is the directory in
        which you'll checkout your code)

```
GOPATH=<workspace_dir>
export GOPATH
PATH="${PATH}:${GOPATH}/bin"
export PATH
```

- Clone the repo in your GOPATH
- Provide your 16-digit secret token in secrets.pem file
- Change the input file in `google-authenticator.go` from `dummy_secret.pem` to `secret.pem`

The auth code works on the secret token and the current time. The time on your local machine should be in sync according to NTP.
- Run `sudo ntpdate time.nist.gov` to sync time
- Run `go run google_authenticator.go`. (For MAC users, it will copy the code to clipboard also)
