# check wget exists
if ! [ -x "$(command -v wget)" ]; then
  echo 'Error: wget is not installed.' >&2
  # fallback to curl
  if ! [ -x "$(command -v curl)" ]; then
    echo 'Error: curl is not installed.' >&2
    exit 1
  fi
  if [ -f resources/la_crime.json ] && [ -f resources/nobel_prize.json ] && [ -f resources/reddit.json ]; then
    echo 'Files already exist.' >&2
    exit 1
  fi
  curl https://data.lacity.org/api/views/2nrs-mtv8/rows.json?accessType=DOWNLOAD -vL -o resources/la_crime.json -H "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:10.0) Gecko/20100101 Firefox/10.0"
  curl http://api.nobelprize.org/v1/prize.json -o resources/nobel_prize.json -vL -H "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:10.0) Gecko/20100101 Firefox/10.0"
  curl https://reddit.com/r/all.json -o resources/reddit.json -vL -H "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:10.0) Gecko/20100101 Firefox/10.0"
  exit 1
fi

wget https://data.lacity.org/api/views/2nrs-mtv8/rows.json?accessType=DOWNLOAD -O resources/la_crime.json
wget http://api.nobelprize.org/v1/prize.json -O resources/nobel_prize.json
wget https://reddit.com/r/all.json -O resources/reddit.json