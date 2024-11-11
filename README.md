# CLI Currency Converter in Golang
A Command Line tool for converting amount from one currency to another. It features a Terminal User Interface (TUI) for smooth user experience.

# Preview

https://github.com/user-attachments/assets/e3e23ac7-cae8-4f1d-b4ee-3b79f3efbb66


## Packages Used
- net/http for http requests to the currency exchange api
- github.com/charmbracelet/huh for the TUI interface form and spinner
- encoding/json in order to marshal the data for the api
- godotenv for loading the environment file containing api key

### API used
I used https://app.exchangerate-api.com/ for converting the currency.

## Supported Currencies
Currently, this project supports conversion between 4 currencies: USD, EUR, JPY and INR.

## Setting up the project in your local environment
1. Clone this repository on your machine
```
git clone https://github.com/kartik699/go-cli-currency-converter.git
```
2. Create an account on the website mentioned in 'API used' section to get your API key
3. Create an env file named `.env.local` and copy paste your API key in this file
```
API_KEY=
```

4.  Run the command `go run .` in your terminal and you are good to go!
