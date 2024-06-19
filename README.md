# EndPointer
EndPointer is a command-line tool designed to crawl a given URL and list all discovered endpoints. Utilizing the Colly web scraping library, EndPointer efficiently navigates through websites to provide a comprehensive list of endpoints, allowing for deeper insight into the structure and available resources of a site.
This was made with a bounty in mind 💰

# Prerequisite 🚀
| Prerequisite | Version |
|--------------|---------|
| Go           |  <=1.22 |
```
apt install golang-go || brew install go
```

# Install 💻
```
git clone https://github.com/CharlesTheGreat77/EndPointer
cd EndPointer
go mod init EndPointer
go mod tidy
```

# Build 👷‍♂️
```
go build -o endpointer main.go
```

# Usage 🦠
To use EndPointer, run the compiled binary with the desired flags:
```
./endpointer -url https://example.com [options]
```

# Available Flags 🏳️

# Features
•	URL Crawling: Specify a target URL to begin crawling.
•	Custom User Agent: Option to set a custom user-agent for requests.
•	Custom Headers: Load custom headers from a file to include in requests. (proxy rotation)
•	Proxy Support: Use a list of proxies (HTTP/HTTPS/SOCKS5) to send requests.
•	Thread Control: Set the number of concurrent threads for crawling.
•	Depth Control: Define the maximum depth for crawling.
•	Timeout Setting: Specify a timeout for requests to handle slow responses.


# Example
```
Example

./endpointer -url https://example.com -user-agent "MyCustomAgent/1.0" -custom-headers headers.txt -proxies proxies.txt -threads 5 -depth 3 -timeout 5
```


# Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

# Acknowledgments
•	Colly - Elegant Scraper and Crawler Framework for Golang.
