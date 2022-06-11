# serverstat-cli [![build](https://github.com/vikpe/serverstat-cli/actions/workflows/test.yml/badge.svg)](https://github.com/vikpe/serverstat-cli/actions/workflows/build.yml)  [![codecov](https://codecov.io/gh/vikpe/serverstat-cli/branch/main/graph/badge.svg)](https://codecov.io/gh/vikpe/serverstat-cli) [![Go Report Card](https://goreportcard.com/badge/github.com/vikpe/serverstat-cli)](https://goreportcard.com/report/github.com/vikpe/serverstat-cli)

```shell
Get info from QuakeWorld servers.

  Usage:   serverstat <address>
Example:   serverstat qw.foppa.dk:27501
```

## Example response
```json
{
  "Address": "91.206.14.17:27503",
  "Mode": "1on1",
  "Title": "1on1: KoLoB vs Zepp [dm4]",
  "Status": "Started",
  "Time": {
    "Elapsed": 4,
    "Total": 10,
    "Remaining": 6
  },
  "PlayerSlots": {
    "Used": 2,
    "Total": 2,
    "Free": 0
  },
  "Players": [
    {
      "Name": "Zepp",
      "NameColor": "wwww",
      "Team": "1",
      "TeamColor": "w",
      "Skin": "",
      "Colors": [
        12,
        12
      ],
      "Frags": 40,
      "Ping": 25,
      "Time": 5,
      "CC": "",
      "IsBot": false
    },
    {
      "Name": "KoLoB",
      "NameColor": "wwwww",
      "Team": "blue",
      "TeamColor": "wwww",
      "Skin": "",
      "Colors": [
        13,
        13
      ],
      "Frags": 1,
      "Ping": 25,
      "Time": 5,
      "CC": "",
      "IsBot": false
    }
  ],
  "Teams": [],
  "SpectatorSlots": {
    "Used": 3,
    "Total": 10,
    "Free": 7
  },
  "SpectatorNames": [
    "NL",
    "player",
    "[ServeMe]"
  ],
  "Settings": {
    "*admin": "glad@quakeworld.ru",
    "*gamedir": "qw",
    "*progs": "so",
    "*qvm": "so",
    "*version": "MVDSV 0.35-dev",
    "*z_ext": "511",
    "deathmatch": "3",
    "fpd": "222",
    "hostname": "Gladius SPB KTX #27503\ufffd",
    "hostname_parsed": "91.206.14.17:27503",
    "ktxver": "1.41-dev",
    "map": "dm4",
    "matchtag": "duel",
    "maxclients": "2",
    "maxfps": "77",
    "maxspectators": "10",
    "pm_ktjump": "1",
    "serverdemo": "duel_kolob_vs_zepp[dm4]110622-1558.mvd",
    "status": "6 min left",
    "sv_antilag": "2",
    "timelimit": "10"
  },
  "QtvStream": {
    "Title": "Gladius SPB QTV (4)",
    "Url": "4@91.206.14.17:28000",
    "Id": 4,
    "Address": "91.206.14.17:28000",
    "SpectatorNames": [
      "twitch.tv/vikpe"
    ],
    "SpectatorCount": 1
  },
  "Geo": {
    "CC": "RU",
    "Country": "Russia",
    "Region": "Europe",
    "City": "St Petersburg",
    "Coordinates": [59.9311, 30.3609]
  },
  "Type": "mvdsv"
}
```

## Download

See [releases](https://github.com/vikpe/serverstat-cli/releases) for downloads.

## See also

* [serverstat](https://github.com/vikpe/serverstat)
* [masterstat](https://github.com/vikpe/masterstat)
* [masterstat-cli](https://github.com/vikpe/masterstat-cli)
