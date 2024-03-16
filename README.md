<h1 align="center">Welcome to GoX 👋</h1>
<p>
  <a href="https://github.com/JesseKoldewijn/GoX/blob/main/LICENCE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

> A simple GoLang based http server that serves htmx content.

## Install

```sh
go mod download
```

## Usage

`[port]` is the port you want to run the server on.

```sh
go run main.go -port [port]
```

# Docker

There also is a Dockerfile included in the project. You can build the image with the following command:

```sh
docker build -t gox .
```

You can run the image with the following command:

```sh
docker run -p [port]:[port] gox
```

or pull the image from the registry:

```sh
docker pull jessekoldewijn/gox:latest
```

## Author

👤 **Jesse Koldewijn**

-   Website: https://jereko.dev
-   Github: [@JesseKoldewijn](https://github.com/JesseKoldewijn)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/JesseKoldewijn/GoX/issues).

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2024 [Jesse Koldewijn](https://github.com/JesseKoldewijn).<br />
This project is [MIT](https://github.com/JesseKoldewijn/GoX/blob/main/LICENCE) licensed.
