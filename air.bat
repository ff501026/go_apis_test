set DB_REPLICAS=0
set SOURCES_HOST=127.0.0.1
set SOURCES_PORT=5432
set SOURCES_USER=postgres
set SOURCES_PASSWORD=pms5566
set SOURCES_DATABASE=postgres
set SOURCES_SSLMODE=disable

echo %SOURCES_DATABASE%
air.exe -c .air.toml
