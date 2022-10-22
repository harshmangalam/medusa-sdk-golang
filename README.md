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
  - [Get Current Customer](#get-current-customer)
  - [Check if email exists](#check-if-email-exists)

- [Cart](#cart)
  - [Add a Shipping Method](#add-a-shipping-method)
  - [Complete a Cart](#complete-a-cart)
  - [Create Payment Sessions](#create-payment-sessions)
  - [Create a Cart](#create-a-cart)
  - [Remove Discount](#remove-discount)
  - [Delete a Payment Session](#delete-a-payment-session)
  - [Refresh a Payment Session](#refresh-a-payment-session)
  - [Get a Cart](#get-a-cart)

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

### Get Current Customer

Gets the currently logged in Customer.

#### Example

```go

data, err := auth.GetSession(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Data)
	fmt.Println(data.Error)
	fmt.Println(data.Errors)

```

### Check if email exists

Checks if a Customer with the given email has signed up.

#### Example

```go
data, err := auth.Exists("harsh@gmail.com", config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Data)
	fmt.Println(data.Error)
	fmt.Println(data.Errors)

```

## Cart

Cart allow handling carts in Medusa.

### Add a Shipping Method

Adds a Shipping Method to the Cart.

#### Example

```go

resp, err := carts.NewShippingMethod().
		SetOptionId(optionId).
		SetData(data).
		Add(cartId, config)

  if err != nil {
		fmt.Println(err)
	}
  
	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Complete a Cart

Completes a cart. The following steps will be performed. Payment authorization is attempted and if more work is required, we simply return the cart for further updates. If payment is authorized and order is not yet created, we make sure to do so. The completion of a cart can be performed idempotently with a provided header Idempotency-Key. If not provided, we will generate one for the request.


```go
resp, err := carts.Complete(cartId, config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Create Payment Sessions
Creates Payment Sessions for each of the available Payment Providers in the Cart's Region.

```go

	resp, err := carts.CreatePaymentSession(cartId, config)
	if err != nil {
		fmt.Println(err)
	}

	
	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Create a Cart

Creates a Cart within the given region and with the initial items. If no region_id is provided the cart will be associated with the first Region available. If no items are provided the cart will be empty after creation. If a user is logged in the cart's customer id and email will be set.

```go

resp, err := carts.NewCreateCart().
		SetCountryCode(countryCode).
		SetItems(items).
		SetRegionId(regionId).
		SetSalesChannelId(salesChannelId).
		SetContext(cartContext).
		Create(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data.Cart)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```


### Remove Discount

Removes a Discount from a Cart.

```go

resp, err := carts.DeleteDiscount(cartId, code, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Delete a Payment Session

Deletes a Payment Session on a Cart. May be useful if a payment has failed.


```go
	resp, err := carts.DeletePaymentSession(cartId, providerId, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Refresh a Payment Session

Refreshes a Payment Session to ensure that it is in sync with the Cart - this is usually not necessary.


```go

	resp, err := carts.RefreshPaymentSession(cartId, providerId, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Get a Cart

Retrieves a Cart.

```go
resp, err := carts.Retrieve(cartId, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```
## License

Licensed under the [MIT License](https://github.com/harshmangalam/medusa-sdk-golang/blob/main/LICENSE)
