# serverstat-cli [![build](https://github.com/vikpe/serverstat-cli/actions/workflows/build.yml/badge.svg)](https://github.com/vikpe/serverstat-cli/actions/workflows/build.yml)  [![codecov](https://codecov.io/gh/vikpe/serverstat-cli/branch/main/graph/badge.svg)](https://codecov.io/gh/vikpe/serverstat-cli) [![Go Report Card](https://goreportcard.com/badge/github.com/vikpe/serverstat-cli)](https://goreportcard.com/report/github.com/vikpe/serverstat-cli)

```shell
Get info from QuakeWorld servers.

  Usage:   serverstat <address>
Example:   serverstat qw.foppa.dk:27501
```

## Download

See [releases](https://github.com/vikpe/serverstat-cli/releases) for downloads.

## Example response

```json
{
  "address": "qw.irc.ax:28503",
  "mode": "2on2",
  "title": "2on2: 1 (HoLy., NinJaA) vs oeks (nig, trl) [dm4]",
  "status": {
    "name": "Started",
    "description": "8 min left"
  },
  "time": {
    "elapsed": 2,
    "total": 10,
    "remaining": 8
  },
  "player_slots": {
    "used": 4,
    "total": 4,
    "free": 0
  },
  "players": [
    {
      "name": "HoLy.",
      "name_color": "wwwww",
      "team": "1",
      "team_color": "w",
      "skin": "",
      "colors": [
        4,
        4
      ],
      "frags": 18,
      "ping": 25,
      "time": 4,
      "cc": "",
      "is_bot": false
    },
    {
      "name": "NinJaA",
      "name_color": "wwwwww",
      "team": "1",
      "team_color": "w",
      "skin": "",
      "colors": [
        4,
        4
      ],
      "frags": 10,
      "ping": 42,
      "time": 4,
      "cc": "",
      "is_bot": false
    },
    {
      "name": "trl.........axe",
      "name_color": "wwwwwwwwwwwwbbb",
      "team": "oeks",
      "team_color": "wwww",
      "skin": "oeks_trl",
      "colors": [
        0,
        0
      ],
      "frags": 18,
      "ping": 13,
      "time": 4,
      "cc": "",
      "is_bot": false
    },
    {
      "name": "nig.........axe",
      "name_color": "wwwwwwwwwwwwbbb",
      "team": "oeks",
      "team_color": "wwww",
      "skin": "oeks_nig",
      "colors": [
        0,
        1
      ],
      "frags": 8,
      "ping": 12,
      "time": 4,
      "cc": "",
      "is_bot": false
    }
  ],
  "teams": [
    {
      "name": "1",
      "name_color": "w",
      "frags": 28,
      "colors": [
        4,
        4
      ],
      "players": [
        {
          "name": "HoLy.",
          "name_color": "wwwww",
          "team": "1",
          "team_color": "w",
          "skin": "",
          "colors": [
            4,
            4
          ],
          "frags": 18,
          "ping": 25,
          "time": 4,
          "cc": "",
          "is_bot": false
        },
        {
          "name": "NinJaA",
          "name_color": "wwwwww",
          "team": "1",
          "team_color": "w",
          "skin": "",
          "colors": [
            4,
            4
          ],
          "frags": 10,
          "ping": 42,
          "time": 4,
          "cc": "",
          "is_bot": false
        }
      ]
    },
    {
      "name": "oeks",
      "name_color": "wwww",
      "frags": 26,
      "colors": [
        0,
        0
      ],
      "players": [
        {
          "name": "trl.........axe",
          "name_color": "wwwwwwwwwwwwbbb",
          "team": "oeks",
          "team_color": "wwww",
          "skin": "oeks_trl",
          "colors": [
            0,
            0
          ],
          "frags": 18,
          "ping": 13,
          "time": 4,
          "cc": "",
          "is_bot": false
        },
        {
          "name": "nig.........axe",
          "name_color": "wwwwwwwwwwwwbbb",
          "team": "oeks",
          "team_color": "wwww",
          "skin": "oeks_nig",
          "colors": [
            0,
            1
          ],
          "frags": 8,
          "ping": 12,
          "time": 4,
          "cc": "",
          "is_bot": false
        }
      ]
    }
  ],
  "spectator_slots": {
    "used": 3,
    "total": 6,
    "free": 3
  },
  "spectator_names": [
    "[ServeMe]",
    "bass",
    "myz"
  ],
  "settings": {
    "*admin": "suom1 \u003csuom1@irc.ax\u003e",
    "*gamedir": "qw",
    "*progs": "so",
    "*qvm": "so",
    "*version": "MVDSV 0.35-dev",
    "*z_ext": "511",
    "deathmatch": "3",
    "fpd": "206",
    "hostname": "QW.IRC.AX KTX:28503 (oeks vs. 1)\ufffd",
    "hostname_parsed": "qw.irc.ax:28503",
    "ktxver": "1.41-dev",
    "map": "dm4",
    "maxclients": "4",
    "maxfps": "77",
    "maxspectators": "6",
    "pm_ktjump": "1",
    "serverdemo": "2on2_oeks_vs_1[dm4]220612-1307.mvd",
    "status": "8 min left",
    "sv_antilag": "2",
    "teamplay": "2",
    "timelimit": "10"
  },
  "qtv_stream": {
    "title": "QW.IRC.AX KTX Qtv (3)",
    "url": "3@46.227.68.148:28000",
    "id": 3,
    "address": "46.227.68.148:28000",
    "spectator_names": [],
    "spectator_count": 0
  },
  "geo": {
    "cc": "SE",
    "country": "Sweden",
    "region": "Europe",
    "city": "Hagersten",
    "coordinates": [
      59.2885,
      17.9612
    ]
  },
  "type": "mvdsv"
}
```

## See also

* [serverstat](https://github.com/vikpe/serverstat)
* [masterstat](https://github.com/vikpe/masterstat)
* [masterstat-cli](https://github.com/vikpe/masterstat-cli)
