# Calculate Pi using the Leibniz formula

## Leibniz formula for π

From Wikipedia, the free encyclopedia:

> In mathematics, the Leibniz formula for π, named after Gottfried Wilhelm Leibniz, states that
> 1 − 1/3 + 1/5 − 1/7 + 1/9 − ... = π/4, an alternating series.

## Usage

### Python

```shell
python main.py 0 1000
```

`out: 3.140592653839794`

The command above returns the sum of series members from 1'st to 1000'th

### Docker

```shell
docker build . -t pi
docker run pi:latest python main.py 0 1000
```

`out: 3.140592653839794`

## Links

- [Wikipedia, the free encyclopedia: Leibniz formula for π](https://en.wikipedia.org/wiki/Leibniz_formula_for_pi)
