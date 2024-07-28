# `ppl-leaderboard` - The leaderboard computation code of PPL

This repository holds the necessary code that is used to compute PPL Era 2 ranks that can be found [here](https://pewpew.live/era2).

## Repository structure

Files:

```text
├── LICENSE
├── README.md
├── accounts.go
├── go.mod
├── go.sum
├── scores.go
└── types.go
```

Info:

* `types.go` - the file that has types necessary for storing player ranks and level leaderboards (`LevelLeaderboard` and `PlayerRank`).
* `accounts.go` - the file that has a few helper functions for creating account id to name maps
* `scores.go` - the file that has all of the code used to make a score map, sorting scores, and of course, computing player ranks!

## Data

The data is stored in a separate repository, [ppl-data](https://github.com/pewpewlive/ppl-data). Here you'll find a few files that are key to computing player ranks, or performing any other sort of analysis on the game leaderboards. The repository is updated with the newest weekly score data, at the time that era rank computations are being performed.

This repository has a submodule linked to `ppl-data`, so you can perform local score analysis. Make sure to pull the latest data once in a while.

## Local testing

In the `cmd` folder, you'll find a `main.go` file that you can run by going into that directory, and running `go run main.go`. Modify the file to your liking, and perform local tests on git cloned data from `ppl-data`.

## Contributing

We are more than happy to accept any contributions to this project. All of the code is used internally to compute the ranks of the Era leaderboards at [pewpew.live](https://pewpew.live). Make a Pull Request with changes that you think could help the computation be more fair, stable, or anything else for that matter, and we will try to merge as much as we can. Thank you in advance!
