# Arbitrage steps

1. fetch the symbol details from all exchanges
2. iterate over each exchange and check for the symbol details
3. fetch the details and store in symbolDetailsMap keyed by the exchange name or identifier
4. iterate over symbolDetailsMap and compare the buy&sell prices across all exchanges.
5. identify potential arbitrage opportunities where sell price on one exchange is higher than the buy price on another.
6. calculate the potential profit considering transaction fees, transfer times, and slippage.
7. store te details of profitable opportunities for later execution.


1. implement executetradeonexchanges method
2. method should take identified arbitrage opportunities and execute trades accourdingly
3. handle order execution to maximize profit and minimize risk (consider order types, quantities, etc.)
4. implement safeguards to handle possible issues, like partial fills, rapidly changing prices.

1. log all key actions and decisions
2. monitor executed trades for success and log issues or failures.