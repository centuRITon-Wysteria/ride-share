<a href="https://github.com/centuRITon-Wysteria/ride-share/blob/master/LICENSE">![LICENSE](https://img.shields.io/github/license/centuRITon-Wysteria/ride-share)</a>
<a href="https://github.com/centuRITon-Wysteria/ride-share">![GitHub last commit](https://img.shields.io/github/last-commit/centuRITon-Wysteria/ride-share?label=GitHub)</a>

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/open-source.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/contains-tasty-spaghetti-code.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/powered-by-coffee.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/not-a-bug-a-feature.svg)](https://forthebadge.com)

# ride-shape [babe-chain]

#### repo on GitHub - [nPerlinNoise](https://github.com/centuRITon-Wysteria/ride-share)

### A Decentralized 

- A _powerful_ and _fast_ API for _n-dimensional_ noise.
- Easy hyper-parameters selection of _octaves_, _lacunarity_ and _persistence_
  as well as complex and customizable hyper-parameters for n-dimension
  _frequency_, _waveLength_, _warp_(interpolation) and _range_.
- Includes various helpful tools for noise generation and for procedural generation tasks
  such as customizable _Gradient_, _Color Gradients_, _Warp_ classes.
- Implements custom _PRNG_ generator for n-dimension and can be easily tuned.

**Details**:

- **Technology stack**:
  > **Status**: **`alpha`** focusing on middleware integration<br>
  > ###### _Tested on go 1.19, Windows 10_
- **Future work**:
  > P2P network on external IP<br>
  > writing **unit tests** for gpp<br>
  > expanding announcement protocols<br>
  > writing in-code docs<br>
  > writing **ReadTheDocs**<br>
  > **blogging**<br>

---

**Screenshots**:

<div align="center">

</div>

---

## Dependencies

- `go>=1.19`
- `github.com/gin-gonic/gin>=1.8.1`
- `github.com/ubiq/g0-ubiq>=3.0.1`

for production dependencies see [Requirements](gpp/go.sum)<br>

## Installation

```shell
$ go mod tidy
```

<a id="usage"></a>

## Usage

**Setup**

```go
```

**Basic usage**

1

- ###### 1

  > 1

    ```go
    ```

for detailed usage see [EXAMPLE](scripts/main.py)

## API

## How to test the software

```shell
go test blockchain
```

```shell
go test gpp
```

## Known issues

- **_`Can't connect to devices outside the network`_**

## Getting help

- Check [main](gpp/main.go) for detailed usage
- Check [Usage](#usage) Section

If you have questions, concerns, bug reports, etc.
please file an [issue](https://github.com/centuRITon-Wysteria/baby-chain/issues) in this repository's Issue Tracker or
open a [discussion](https://github.com/centuRITon-Wysteria/baby-chain/discussions) in this repository's Discussion section.

## Getting involved

<a id="contribute"></a>
- `Looking for Contributors for feature additions`
- `Looking for Contributors for optimization`
- `Looking for Contributors for unit testing`
- `Looking for Contributors for ReadTheDocs`
- `Looking for Contributors for WebApp`
- `Looking for Contributors for API `
- [Fork](https://github.com/centuRITon-Wysteria/baby-chain/fork) the repository
  and issue a [PR](https://github.com/centuRITon-Wysteria/baby-chain/pulls) to contribute

General instructions on _how_ to contribute [CONTRIBUTING](CONTRIBUTING.md)
and [CODE OF CONDUCT](CODE_OF_CONDUCT.md)

----

## Open source licensing info

1. [TERMS](TERMS.md)
2. [LICENSE](LICENSE)
3. [CFPB Source Code Policy](https://github.com/cfpb/source-code-policy/)

----

## Credits and references

1. Inspired from [The Coding Train](https://www.youtube.com/channel/UCvjgXvBlbQiydffZU7m1_aw)
   -> [perlin noise](https://thecodingtrain.com/challenges/24-perlin-noise-flow-field)
2. hash function by [xxhash](https://github.com/Cyan4973/xxHash)
   inspired the [rand3](src/nPerlinNoise/tools.py) algo
   and ultimately helped for O(1) time complexity n-dimensional random generator [NPrng](src/nPerlinNoise/tools.py)
3. [StackOverflow](https://stackoverflow.com/) for helping on various occasions throughout the development
4. [vnoise](https://github.com/plottertools/vnoise) and [opensimplex](https://github.com/lmas/opensimplex)
   for ideas for README.md
5. docs derivative from [open-source-project-template](https://github.com/cfpb/open-source-project-template)
6. packaging help from [realpython](https://realpython.com/pypi-publish-python-package/)

**Maintainer**:

|        <a href="https://github.com/Amith225"><img src="https://avatars.githubusercontent.com/u/75326634?v=4" height=250></a>        |
|:-----------------------------------------------------------------------------------------------------------------------------------:|
|                                    **[Amith M](https://www.linkedin.com/in/amith-m-17088b246/)**                                    |
| [![Instagram](https://img.shields.io/badge/Instagram-%23E4405F.svg?logo=Instagram&logoColor=white)](https://instagram.com/amithm3 ) |
