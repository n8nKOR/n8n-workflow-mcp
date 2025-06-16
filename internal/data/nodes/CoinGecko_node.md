# CoinGecko Node

## Overview

Consume CoinGecko API

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: coin

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Coin Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Options | collection | Whether to include community data | No |  |

#### Operation: candlestick

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Base Currency Name or ID | options | The first currency in the pair. For BTC:ETH this is BTC. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: market

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Base Currency Name or ID | options | The first currency in the pair. For BTC:ETH this is BTC. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Options | collection | Filter results by comma-separated list of coin ID | No |  |

#### Operation: price

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Base Currency Names or IDs | multiOptions | The first currency in the pair. For BTC:ETH this is BTC. Choose from the list, or specify IDs using an <a href= | Yes |  |
| Contract Addresses | string | The contract address of tokens, comma-separated | Yes |  |
| Quote Currency Names or IDs | multiOptions | The second currency in the pair. For BTC:ETH this is ETH. Choose from the list, or specify IDs using an <a href= | Yes |  |
| Options | collection |  | No |  |

#### Operation: marketChart

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Base Currency Name or ID | options | The first currency in the pair. For BTC:ETH this is BTC. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: history

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Date | dateTime | The date of data snapshot | Yes |  |
| Options | collection | Whether to exclude localized languages in response | No |  |

#### Operation: ticker

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Options | collection | Filter results by exchange IDs. Choose from the list, or specify IDs using an <a href= | No |  |

### Resource: event

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Country code of event. Choose from the list, or specify an ID using an <a href= | No |  |

## UseCases

- **Cryptocurrency Portfolio Tracking**: Monitor real-time prices and market data for cryptocurrency portfolios and investment tracking
- **Trading Bot Integration**: Feed live price data, OHLC candlestick data, and market information to automated trading systems
- **Market Analysis and Research**: Gather comprehensive market data, historical prices, and trading volumes for cryptocurrency research
- **Price Alert Systems**: Create automated alerts when cryptocurrencies reach specific price thresholds or market conditions
- **DeFi Protocol Monitoring**: Track token prices and market caps for decentralized finance protocols and yield farming
- **Arbitrage Opportunity Detection**: Compare prices across different exchanges to identify arbitrage opportunities
- **Tax Reporting and Compliance**: Gather historical price data for cryptocurrency tax calculations and regulatory reporting
- **Market Intelligence Dashboards**: Build comprehensive dashboards with real-time crypto market data and trending information
- **Risk Management Systems**: Monitor market volatility, trading volumes, and price movements for risk assessment
- **Cryptocurrency News and Events**: Stay updated with crypto events, announcements, and market-moving news for informed trading decisions

