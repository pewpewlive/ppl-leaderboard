# `ppl-leaderboard` - The leaderboard computation code of PPL

This repository holds the code used to compute the [PPL Era 2 ranks](https://pewpew.live/era2).

## Repository structure

Files:

```text
├── cmd/
    ├── data/
    └── main.go
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── csv.go
├── scores.go
└── types.go
```

Info:

* `types.go` - contains the structures that store the leaderboards and player ranks (`LevelLeaderboard` and `PlayerRank`).
* `csv.go` - contains the utility functions for reading data from CSV.
* `scores.go` - contains the code that computes the leaderboards and the player ranks.

## Running the leaderboard computation

You need to obtain the latest [leaderboard data](https://github.com/pewpewlive/ppl-data) in cmd/data/:
```
git submodule update --init --recursive
```

You can then run the computation of the leaderboard:
```
cd cmd/
go run main.go
```

## Contributing

We welcome contributions that improve the computation of leaderboards.