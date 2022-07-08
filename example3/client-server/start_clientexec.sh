sleep 10
while true ; do

    echo start test

    ./clientexec

    echo end test $!

    sleep 10
done
