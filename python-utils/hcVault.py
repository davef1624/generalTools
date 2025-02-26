import hvac 

def hcVaultClient():
    # Initialize the Vault client
    client = hvac.Client(url='http://127.0.0.1:8200', token='my-root-token')

    # Verify if the client is authenticated
    print(client.is_authenticated())  
    return client

def main():
    # Create a hcVault client
    vaultClient = hcVaultClient()
    print(vaultClient)

if __name__ == "__main__":
    main()