echo start test

./serverexec &
SPID=$!
sleep 3
./clientexec
kill $SPID


echo Start serverexec
/server/serverexec
echo Exit serverexec $?
