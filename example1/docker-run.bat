set PWD=%CD%
cd %~dp0

call docker-params.bat

docker run -ti -v "%CD%\\..:/mnt/host" -w /mnt/host/example1 -p 127.0.0.1:8080:80 %IMAGE%:%TAG%
