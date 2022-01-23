#!/bin/bash
OUTFILE="benchmark_result.log"

for bytes in 10 20 50 100 200 1000 5000;
do
	echo "======start operate bytes size: $bytes======" >> $OUTFILE
	./src/redis-benchmark -h 127.0.0.1 -c 10 -d $bytes -n 10000 -t get >> $OUTFILE
	./src/redis-benchmark -h 127.0.0.1 -c 10 -d $bytes -n 10000 -t set >> $OUTFILE
	echo -e "================end operate=================\n" >> $OUTFILE
done
