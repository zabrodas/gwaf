set PWD=%CD%
cd %~dp0

call docker-params.bat

docker build -f %PWD%\Dockerfile -t %IMAGE%:%TAG% .
@if errorlevel 1 goto err

set excode=0
@echo Success
goto fin

:err
set excode=1
@echo ERROR

:fin
cd %PWD%
exit /b %excode%
