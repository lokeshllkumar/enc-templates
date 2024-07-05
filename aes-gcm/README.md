# AES-GCM

A simple model designed in Go to implement the Advanced Encryption Standard(AES) block cipher in Galois/Counter mode(GCM).

## Usage

The `key.txt` file contains a 32-bit encryption key. A fresh key can be generated using the function`KeyGen()`. The generated key is also stored in a public variable called `Key`.

The `readKey()` function tis private and used to read the generated key from the `key.txt` file and write it to the public `Key` variable.

The `EncryptPasssword()` and `DecryptPassword()` functions are used to encrypt and decrypt data passed as a parameter using AES-GCM and the generated key, respectively. Encrypted/decrypted data is returned as a string.