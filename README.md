# Password Manager

## This is a draft version of the program, note that password are only encoded in Base64. AES encryption is coming shortly.


A secure and efficient tool for managing your passwords.

## Features

- **Secure Storage**: Encrypts passwords and stores them safely.
- **Easy Retrieval**: Quickly access and manage your passwords.
- **User-Friendly Interface**: Simple and intuitive UI for ease of use.

## Extra Security Mesures (For Windows Users) :

**Encrypt the Application Using BitLocker:**

- Locate the folder containing your application.
- Right-click the folder and select Properties.
- Go to the General tab and click Advanced.
- Check the box for Encrypt contents to secure data and click OK.
- Click Apply and choose to encrypt the folder and all its contents.

**Secure Access with Windows Hello:**

- BitLocker encryption ties access to your user account credentials.
- Ensure your account is set up with Windows Hello PIN. To enable it:
- Go to Settings > Accounts > Sign-in options.
- Under PIN (Windows Hello), click Add PIN and set a PIN.
- When someone attempts to access the encrypted file or folder, they will need your PIN.

## Installation

To install the Password Manager, follow these steps:

1. **Download the latest release** from the [Releases page](https://github.com/maxBRT/password-manager/releases).
