# Config Module

[![Go CI](https://github.com/Snider/config/actions/workflows/ci.yml/badge.svg)](https://github.com/Snider/config/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/Snider/config/branch/main/graph/badge.svg)](https://codecov.io/gh/Snider/config)

This repository is a config module for the Core Framework. It includes a Go backend, an Angular custom element, and a full release cycle configuration.

## Getting Started

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Snider/config.git
    ```

2.  **Install the dependencies:**
    ```bash
    cd config
    go mod tidy
    cd ui
    npm install
    ```

3.  **Run the development server:**
    ```bash
    go run ./cmd/demo-cli serve
    ```
    This will start the Go backend and serve the Angular custom element.

## Building the Custom Element

To build the Angular custom element, run the following command:

```bash
cd ui
npm run build
```

This will create a single JavaScript file in the `dist` directory that you can use in any HTML page.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the EUPL-1.2 License - see the [LICENSE](LICENSE) file for details.
