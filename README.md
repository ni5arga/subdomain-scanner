# Subdomain Scanner
A wordlist based subdomain scanner written in golang. It scans a target domain for subdomains using a provided wordlist and saves the discovered subdomains with a status code of 200 to an output file.

## Usage

### Prerequisites

- Install [Go](https://golang.org/doc/install)

### Running the Scanner

1. Clone or download the repository.

2. Open a terminal and navigate to the directory containing `scanner.go`.

3. Run the following command:

    ```bash
    go run scanner.go -domain example.com
    ```

    Replace `example.com` with the target domain you want to scan.

### Options

- `-domain`: Specify the target domain to scan (required).

- `-wordlist`: Path to the wordlist file containing subdomains. Defaults to `subdomains.txt`.

- `-output-file`: Output file to write found subdomains. Defaults to `found-subdomains.txt`.

### Example

```bash
go run scanner.go -domain example.com -wordlist custom-wordlist.txt -output-file discovered-domains.txt
```

This command will scan `example.com` using a custom wordlist (`custom-wordlist.txt`) and save the discovered subdomains to `discovered-domains.txt`.

## Note

- The program will only write subdomains with a status code of 200 to the output file.

- Ensure that you have the necessary permissions to scan the target domain.

- Use responsibly and respect the terms of service of the target domain.

## Wordlist Credit
[dnscan](https://github.com/rbsec/dnscan)

Feel free to modify the wordlist, output file, or any other parameters based on your needs.

