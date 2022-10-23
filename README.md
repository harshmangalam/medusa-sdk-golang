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
  <a href="https://pkg.go.dev/github.com/harshmangalam/medusa-sdk-golang@v0.0.0-20221023101217-804ec82d21a9">Documentation</a> |
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

## Author
- Harsh Mangalam


## Resources
- [Medusa’s GitHub repository](https://github.com/medusajs/medusa
)
- [How to Create Services](https://docs.medusajs.com/advanced/backend/services/create-service)

- [Medusa SDK Golang docs](https://pkg.go.dev/github.com/harshmangalam/medusa-sdk-golang@v0.0)

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
  - [Select a Payment Session](#select-a-payment-session)
  - [Update a Payment Session](#update-a-payment-session)
  - [Update a Cart](#update-a-cart)

- [Collection](#collection)

  - [Get a Collection](#get-a-collection)
  - [List Collections](#list-collections)

- [Customer](#customer)

  - [Create a Customer](#create-a-customer)
  - [List Orders](#list-orders)
  - [Request Password Reset](#request-password-reset)
  - [Reset Password](#reset-password)
  - [Get a Customer](#get-a-customer)
  - [Update Customer](#update-customer)

- [Gift Card](#gift-card)

  - [Get Gift Card by Code](#get-gift-card-by-code)

- [Order](#order)

  - [Get by Cart ID](#get-by-cart-id)
  - [Look Up an Order](#look-up-an-order)
  - [Get an Order](#get-an-order)

- [OrderEdit](#order-edit)

  - [Completes an OrderEdit](#completes-an-order-edit)
  - [Decline an OrderEdit](#decline-an-order-edit)
  - [Retrieve an OrderEdit](#retrieve-an-orderedit)

- [Product](#product)

  - [Get a Product](#get-a-product)
  - [List Products](#list-products)
  - [Search Products](#search-products)
  - [Get a Product Variant](#get-a-product-variant)
  - [Get Product Variants](#get-product-variants)

- [Region](#region)

  - [Get a Region](#get-a-region)
  - [List Regions](#list-regions)

- [Return](#return)

  - [Create Return](#create-return)

- [Return Reason](#return-reason)

  - [Get a Return Reason](#get-a-return-reason)
  - [List Return Reasons](#list-return-reasons)

- [Shipping Option](#shipping-options)

	- [Get Shipping Options](#get-shipping-options)
	- [List for Cart](#list-for-cart)

- [Swap](#swap)

	- [Create a Swap](#create-a-swap)
	- [Get by Cart ID](#get-by-cart-id)



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

### Select a Payment Session

Selects a Payment Session as the session intended to be used towards the completion of the Cart.

```go
resp, err := carts.NewSelectPaymentSession().
		SetProviderId(providerId).
		Select(crtId, config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Update a Payment Session

Updates a Payment Session with additional data.

```go
resp, err := carts.NewUpdatePaymentSession().
		SetData(data).
		Update(cartId, providerId, confi)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Update a Cart

Updates a Cart.

```go
resp, err := carts.NewUpdateCart().
		SetBillingAddress(billingAddress).
		SetContext(context).
		SetCountryCode(countryCode).
		SetCustomerId(customerId).
		SetDiscounts(discount).
		SetEmail(email).
		SetGiftCards(giftCards).
		SetRegionId(regionId).
		SetSalesChannelId(salesChannelId).
		SetShippingAddress(shippingAddress).
		Update(cartId, config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

## Collection

Collection allow handling collections in Medusa.

### Get a Collection

Retrieves a Product Collection.

```go

	resp, err := collections.Retrieve(id, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)


```

### List Collections

Retrieve a list of Product Collection.

```go
createdAt := common.NewDateComparison().
		SetLte(time).
		SetGt(time).
		SetGte(time).
		SetLt(time)

	updatedAt := common.NewDateComparison().
		SetLte(time).
		SetGt(time).
		SetGte(time).
		SetLt(time)

	resp, err := collections.
		NewCollectionsQuery().
		SetLimit(4).SetOffset(2).
		SetCreatedAt(createdAt).
		SetUpdatedAt(updatedAt).
		List(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)
```

## Customer

Customer endpoints that allow handling customers in Medusa.

### Create a Customer

Creates a Customer account.

```go
resp, err := customers.
		NewCreateCustomer().
		SetFirstName(firstname).
		SetLastName(lastname).
		SetEmail(email).
		SetPassword(password).
		Create(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)
```

### List Orders

Retrieves a list of a Customer's Orders.

```go

	resp, err := customers.NewListOrderQuery().
		SetCancledAt(dateComparison).
		SetCartId(cartId).
		SetCreatedAt(dateComparison).
		SetCurrencyCode(currencyCode).
		SetEmail(email).
		SetExpand(expand).
		SetFields(fields).
		SetFulfillmentStatus(fulfillmentStatus).
		SetId(id).
		SetLimit(limit).
		SetOffset(offset).
		SetPaymentStatus(paymentStatus).
		SetQ(q).
		SetRegionId(regionId).
		SetStatus(status).
		SetTaxRate(taxRate).
		SetUpdatedAt(dateComparison).
		List(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Request Password Reset

Creates a reset password token to be used in a subsequent /reset-password request. The password token should be sent out of band e.g. via email and will not be returned.

```go

resp, err := customers.NewRequestPasswordReset().
		SetEmail(email).
		RequestReset(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)


```

### Reset Password

Resets a Customer's password using a password token created by a previous /password-token request.

```go
	resp, err := customers.
		NewResetPassword().
		SetEmail(email).
		SetPassword(password).
		SetToken(token).
		Reset(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)
```

### Get a Customer

Retrieves a Customer - the Customer must be logged in to retrieve their details.

```go
resp, err := customers.Retrieve(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Update Customer

Updates a Customer's saved details.

```go

	resp, err := customers.
		NewUpdateCustomer().
		SetFirstName(firstname).
		SetLastName(lastname).
		SetBillingAddress(address).
		SetMetadata(metaData).
		SetPassword(password).
		SetPhone(phone).
		SetEmail(email).
		Update(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

## Gift Card

Gift Card endpoints that allow handling gift cards in Medusa.

### Get Gift Card by Code

Retrieves a Gift Card by its associated unqiue code.

```go
	resp, err := giftcards.Retrieve(code, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

## Order

Order allow handling orders in Medusa.

### Get by Cart ID

Retrieves an Order by the id of the Cart that was used to create the Order.

```go

	resp, err := orders.RetrieveByCartId(cartId, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Look Up an Order

Look up an order using filters.

```go
	resp, err := orders.NewLookup().
		SetDisplayId(displayId).
		SetEmail(email).
		SetShippingAddress(address).
		Lookup(cartId, config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)
```

### Get an Order

Retrieves an Order

```go
	resp, err := orders.Retrieve(id, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

## OrderEdit

### Completes an OrderEdit

Completes an OrderEdit.

```go
	resp, err := orderedits.Complete(id, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Decline an OrderEdit

Declines an OrderEdit.

```go

resp, err := orderedits.NewDeclineOrderEdit().
		SetDeclineReason(declineReason).
		Decline(id, config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Retrieve an OrderEdit

Retrieves a OrderEdit.

```go
	resp, err := orderedits.Retrieve(id, config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)
```

## Product

### Get a Product

Retrieves a Product.

```go
	resp, err := products.Retrieve(id, config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)
```

### List Products

Retrieves a list of Products.

```go

resp, err := products.NewListProduct().
	SetCollectionIds(collectionsId).
	SetCreatedAt(createdAt).
	SetDescription(desc).
	SetExpand(expand).
	SetFields(fields).
	SetHandle(handle).
	SetIds(ids).
	SetIsGiftcard(isGiftCard).
	SetLimit(limit).
	SetOffset(offset).
	SetQ(q).
	SetTags(tags).
	SetTitle(title).
	SetType(type).
	SetUpdatedAt(updatedAt).
	List(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Search Products

Run a search query on products using the search engine installed on Medusa

```go
	resp, err := products.NewSearchProduct().
		SetLimit(limit).
		SetOffset(offset).
		SetQ(q).
		Search(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)
```

### Get a Product Variant

Retrieves a Product Variant by id

```go
resp, err := products.RetrieveVariant(variantId,config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Retrieves a list of Product Variants

Get Product Variants

```go
resp, err := products.NewListProuductVariant().
		SetExpand(expand).
		SetIds(ids).
		SetInventoryQuantity(invQty).
		SetLimit(limit).
		SetOffset(offset).
		SetTitle(title).
		List(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

## Region

### Get a Region

```go
	resp, err := regions.Retrieve(regionId,config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### List Regions

```go
	resp, err := regions.NewListRegion().
		SetCreatedAt(createdAt).
		SetLimit(limit).
		SetOffset(offset).
		SetUpdatedAt(updatedAt).
		List(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

## Return

### Create Return

```go

resp,err := returns.NewCreateRetun().
		SetItems(items).
		SetOrderId(orderId).
		SetReturnShipping(returnShippings).
		Create(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)
```

## Return Reason

### Get a Return Reason

```go

	resp, err := returnreasons.Retrieve(id,config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### List Return Reasons

```go

	resp, err := returnreasons.List(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

## Shipping Option

### Get Shipping Options


```go

	resp, err := shippingoptions.NewListShippingOption().
		SetIsreturn(isReturn).
		SetProductIds(productIds).
		SetRegionId(regionId).
		List(config)
		
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### List for Cart

```go

	resp, err := shippingoptions.ListCartOptions(cartId,config)
		
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

## Swap

### Create a Swap

```go

	resp, err := swaps.NewCreateSwap().
		SetAdditionalItems(additionalItems).
		SetOrderId(orderId).
		SetReturnItems(returnItems).
		SetReturnShippingOption(returnShippingOptions).
		Create(config)
		
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```

### Get by Cart ID

```go

	resp, err := swaps.RetrieveByCartId(cartId,config)
		
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)
	fmt.Println(resp.Error)
	fmt.Println(resp.Errors)

```
## License

Licensed under the [MIT License](https://github.com/harshmangalam/medusa-sdk-golang/blob/main/LICENSE)
