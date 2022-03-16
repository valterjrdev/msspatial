# ms spatial

Calculate Spatial distance using methods:
 * Manhattan distance 

## Usage

Install Dependencies

* [docker engine](https://docs.docker.com/engine/install/ubuntu/)
* [docker-compose](https://docs.docker.com/compose/install/) 

Generate `.env` file
```bash
cp .env.example .env
```
Run using docker-compose
```bash
make run
```

Access api doc:
```
http://127.0.0.1:8000/docs/index.html
```

## Tools
Install development tools
```bash
make install
```

Unit tests
```bash
make test
```

Generate coverage
```bash
make coverage
```

Lint
```bash
make lint
```

## Help

Tools for debug and compare results:
 * https://www.omnicalculator.com/math/manhattan-distance