last_day=$(ls | grep '^day[0-9]*$' | sort -V | tail -n 1);
next_day=$(( ${last_day:3} + 1 ));
cp -r "$last_day" "day$next_day"
rm -rf "day$next_day/main2"

