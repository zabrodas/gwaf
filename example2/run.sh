./server &
SPID=$!
sleep 3
./client
kill $SPID

