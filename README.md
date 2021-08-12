# ecommerce-service

Basic ecommerce service written in Go for educational purposes.

# Prerequisites

You should have at least Go 1.16 installed in your system.

Optionally, if you want to have live reloading code, you must have [air](https://github.com/cosmtrek/air) installed.

# Quickstart

1. This repo is a template. [Use it](https://github.com/cgrs/ecommerce-service-starter/generate) to generate a new repo
in your personal account with the basic folder structure.

2. After it's created, clone the repository in your machine, navigate to the directory and launch it:

    ```
    cd ecommerce-service/
    go run .
    ```

3. The basic program should greet you with the URL is listening to:
    ```
    2021/08/10 09:22:42 Server is listening on http://localhost:3000
    ```
    Navigate to [that URL](http://localhost:3000) and you'll receive a JSON with the following data:
    
    ```json
    {"status": 200, "message": "OK"}
    ```

    Meanwhile, the terminal should display something like this:

    ```
    2021/08/12 23:48:58 GET / 154.257µs
    ```

# Project definition

In order to build a simple ecommerce service, we'll define some entities:

* Customer. In charge of adding items to its basket and making orders.
* Item. Products which are bought by the customers.
* Basket. The temporary storage where a customer decides which products are going to buy.
* Order. After confirming the checkout, the order displays a list of the products bought by a customer in a shopping session.

Given these entities, we can infer these behaviors:

* Customers can make a lot of orders but they have only one basket, corresponding to their current shopping session.
* Customers can modify their baskets as they wish: adding more items, increasing or decreasing some specific item's
  quantity, or even emptying their cart. However, orders are final and they cannot be modified once customers have
  made the checkout.
* Customers are allowed to see how many orders they have made and how much each order was.
* Once a customer decides to checkout, An order is made and their basket is emptied.

To make the development easier, we'll make some assumptions:

* There's an infinite stock of items (that is, we don't need to control how many items are left for customers to buy).
* Customers payment is irrelevant (let's imagine it's managed by other service we cannot configure, but it ensures
  every customer pays their orders).

# Roadmap

Given the previous definition, you'll need to:
- [ ] Develop the customers domain:
    * Basic identity management for customers (Signing up, signing in...). There's no need to get fancy here.
      Just a map of username/passwords in plain text will do. (Needless to say you shouldn't do that in a real
      environment, right? *...* [**_Right?_**](https://i.kym-cdn.com/photos/images/newsfeed/002/081/388/497.jpg)).
    * Get some kind of session management in order to get the customer's basket working. Cookies are more than enough.
- [ ] Develop the items domain:
    * Implement a basic search for customers in order to give them a way to display the items they want to buy.
    * Write some item management endpoints for the owner of the shop to create, update and delete items. (Bonus points
      if you implement a basic role access).
- [ ] Develop the basket domain:
    * Let the customers modify their basket as they want (as said before).
    * Create an order if a customer does checkout and after that, generate a new shopping session for that customer.
- [ ] Develop the orders domain:
    * Write some endpoints to display the orders of a customer.

# Solution

> :warning: If you get stuck, start a discussion in this repo to get hints instead of just copying the solution.

You can find a working solution (not the most optimal though) in the [`solution`](https://github.com/cgrs/ecommerce-service-starter/tree/solution) branch.


# Colophon

I guess that's everything for now. If you think this workshop can be improved, just feel free to create an issue
or open a pull request.