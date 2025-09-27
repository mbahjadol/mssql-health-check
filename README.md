# mssql-health-check
Tiny independent web service that check of MSSQL health check, that can be implemented in nix base container .. (especially for uptime-kuma by date of 2025-09-27 they latest release is bugs)

## It will using system environment parameter for this to be running like
- MSSQL_USER
- MSSQL_PASS
- MSSQL_HOST
- MSSQL_PORT
- MSSQL_DB

so this will running in daemon
