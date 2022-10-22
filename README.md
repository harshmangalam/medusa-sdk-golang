<p align="center">
  <a href="https://www.medusajs.com">
    <img alt="Medusa" src="https://user-images.githubusercontent.com/7554214/153162406-bf8fd16f-aa98-4604-b87b-e13ab4baf604.png" width="100" />
  </a>
   <a href="https://go.dev/">
    <img alt="Medusa" src="https://miro.medium.com/max/1200/1*i2skbfmDsHayHhqPfwt6pA.png" width="100" />
  </a>
</p>
<h1 align="center">
  Medusa  SDK Golang
</h1>

<h4 align="center">
  <a href="#">Documentation</a> |
  <a href="https://www.medusajs.com">Website</a>
</h4>

<p align="center">
An open source medusa sdk for golang
</p>
<p align="center">
  <a href="https://github.com/harshmangalam/medusa-sdk-golang/blob/main/LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="Medusa is released under the MIT license." />
  </a>
  <a href="https://github.com/harshmangalam/medusa-sdk-golang/blob/main/CONTRIBUTING.md">
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat" alt="PRs welcome!" />
  </a>

</p>

## Contents <!-- omit in toc -->

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Configuration](#config)
- [Auth](#auth)
  - [Customer Login](#customer-login)
  - [Customer Log out](#customer-log-out)

## Getting Started

You can install Medusa by either following our [Quickstart guide](https://docs.medusajs.com/quickstart/quick-start) or the following steps:

1. **Install Medusa CLI**

    ```bash
    npm install -g @medusajs/medusa-cli
    ```

2. **Create a new Medusa project**

    ```bash
    medusa new my-medusa-store --seed
    ```

3. **Start your Medusa engine**

    ```bash
    cd my-medusa-store
    medusa develop
    ```

## Installation

### required go v1.18+

```bash
go get github.com/harshmangalam/medusa-sdk-golang@latest
```


## Configuration

```go

config := medusa.NewConfig().
		SetMaxRetries(3).
		SetBaseUrl("http://localhost:9000")

```

## Auth
Auth endpoints that allow authorization of customers and manages their sessions

### Customer Login
Logs a Customer in and authorizes them to view their details. Successful authentication will set a session cookie in the Customer's browser.

#### Example

```go

resp, err := auth.NewAuth().
		SetEmail("harsh@gmail.com").
		SetPassword("123456").
		Authenticate(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)


```

### Customer Log out

Destroys a Customer's authenticated session.

#### Example

```go

data, err := auth.DeleteSession(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Data)
	fmt.Println(data.Error)
	fmt.Println(data.Errors)

```

## License

Licensed under the [MIT License](https://github.com/harshmangalam/medusa-sdk-golang/blob/main/LICENSE)
