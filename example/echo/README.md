## echo service for the example

** Simple Echo service as an example **

### Installation

```bash
npm run build
npm test
```

### Usage

#### Starting service

```bash
npm start
```

#### Using service

```bash
curl -sd '
{
  "role": "echo",
  "cmd": "test"
}
' localhost:3030/act
```

### License

Copyright (c) 2015 Geoffrey Clements. - All rights reserved.
