# mole-devbox

Forwards every TCP port a remote dev host listens on to the same port on your laptop.
It re-scans every 10 seconds, so new services appear without editing `~/.ssh/config`.

```
mole-devbox devbox              # follow the host, rescan every 10s
mole-devbox -i 0 devbox         # scan once
mole-devbox -p 3000,4500 devbox # only these ports, no scan
mole-devbox -x 4444 devbox      # also skip 4444
mole-devbox -n devbox           # print the mole command and exit
```

How it works:

- Resolves `<ssh-host>` with `ssh -G`, so `Host` aliases, `Include` and `Match` work.
- Lists listeners with `ssh <host> ss -ltnH`. Remote host must be Linux.
- Keeps ports 1024-32767 bound to loopback or a wildcard address. Skips 5355, 9091 and 9100.
- A service bound only to `[::1]` (Vite, for example) is forwarded to `[::1]`.
- Skips a port that is already in use on the laptop and logs it.
- Restarts mole when the set of remote ports changes.

Limits:

- mole does not support `ProxyJump` or `ProxyCommand`. The script stops if the host uses either.
- mole needs a private key file. The script uses the first `IdentityFile` that exists.
  A key with a passphrase is signed by the agent in `IdentityAgent` or `SSH_AUTH_SOCK`.
- mole checks the host key against `~/.ssh/known_hosts` only.
