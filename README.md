Docgen
=====================

Transform your postman collection to HTML/Markdown documentation

![Task screenshot](screenshot.png)

#### Installation on Mac/Linux
```bash
curl https://raw.githubusercontent.com/thedevsaddam/docgen/v3/install.sh -o install.sh && sudo chmod +x install.sh && sudo ./install.sh
```
#### Uninstallation
```bash
curl https://raw.githubusercontent.com/thedevsaddam/docgen/v3/install.sh -o uninstall.sh && sudo chmod +x uninstall.sh && sudo ./uninstall.sh
```

#### Windows
**For windows download the binary and set environment variable so that you can access the binary from terminal**

#### Binary link
[Download binary](https://github.com/thedevsaddam/docgen-bin/tree/master/latest)

### Available features
* Live postman collection to documentation
* Build postman collection to html/markdown documentation
* Supports multi-level collection build

### Usage
* To view live HTML documentation from postman collection use `docgen server -f input-postman-collection.json -p 8000` This will open the html version of postman collection to the defined port
* To view live Markown documentation from postman collection use `docgen server -f input-postman-collection.json -p 8000 -m` This will open the markdown version of postman collection to the defined port
* To make HTML documentation use `docgen build -i input-postman-collection.json -o ~/Downloads/index.html`
* To make Markdown documentation use `docgen build -i input-postman-collection.json -o ~/Downloads/index.md -m`

***[Demo markdown API documentation](_examples/example-doc.md)***

### Author
1. [Sajib Sikder](https://github.com/mhshajib)
1. [Saddam H](https://github.com/thedevsaddam)

### Contributor
1. [Anondo](https://github.com/Anondo)

### Contribution
Your suggestions will be more than appreciated.
[Read the contribution guide here](CONTRIBUTING.md)

### See all [contributors](https://github.com/thedevsaddam/docgen/graphs/contributors)

### **License**
The **docgen** is an open-source software licensed under the [MIT License](LICENSE.md).

## Support on Beerpay
Hey dude! Help me out for a couple of :beers:!

[![Beerpay](https://beerpay.io/thedevsaddam/docgen/badge.svg?style=beer-square)](https://beerpay.io/thedevsaddam/docgen)  [![Beerpay](https://beerpay.io/thedevsaddam/docgen/make-wish.svg?style=flat-square)](https://beerpay.io/thedevsaddam/docgen?focus=wish)