# check wget exists
if ! [ -x "$(command -v wget)" ]; then
  echo 'Error: wget is not installed.' >&2
  # fallback to curl
  if ! [ -x "$(command -v curl)" ]; then
    echo 'Error: curl is not installed.' >&2
    exit 1
  fi
  curl https://data.lacity.org/api/views/2nrs-mtv8/rows.json?accessType=DOWNLOAD -o resources/la_crime.json
  curl http://api.nobelprize.org/v1/prize.json -o resources/nobel_prize.json
  curl https://reddit.com/r/all.json -o resources/reddit.json
  exit 1
fi

wget https://data.lacity.org/api/views/2nrs-mtv8/rows.json?accessType=DOWNLOAD -O resources/la_crime.json
wget http://api.nobelprize.org/v1/prize.json -O resources/nobel_prize.json
wget https://reddit.com/r/all.json -O resources/reddit.json