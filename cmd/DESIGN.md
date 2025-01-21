**Design Document: `plexctl` - A command-Line tool for Plex Media Server**

**1. Introduction**

* **1.1 Purpose**
    * This document outlines the design and implementation of a command-line tool named `plexctl` for interacting with a Plex Media Server using its REST API.
    * `plexctl` will allow users to perform various administrative tasks, query server information, and control media playback remotely.
* **1.2 Scope**
    * This initial version will focus on essential functionalities such as:
        * **Server Management:**
            * Get server status and information
            * Restart/Update the server
        * **Library Management:**
            * Scan libraries
            * Optimize media files
        * **Playback Control:**
            * Play, pause, and stop playback on clients
            * Seek within currently playing media
            * Get information about currently playing media
* **1.3 Target Audience**
    * This tool is intended for Plex Media Server administrators, power users, and those who prefer a command-line interface for managing their media server

**2. User Interface**

* **2.1 Commands**
    * **Authorization:**
        * `plexctl auth login`: Login into Plex
        * `plexctl auth list`: List available authorized accounts
        * `plexctl auth revoke`: Revoke Plex credentials
    * **Server Management:**
        * `plexctl server status`: Get server status (uptime, version, transcoder load)
        * `plexctl server info`: Get detailed server information (hardware, libraries, etc.)
        * `plexctl server restart`: Restart the Plex Media Server
        * `plexctl server update`: Check for and install updates
    * **Library Management:**
        * `plexctl library scan <library-name>`: Scan a specific library for new media
        * `plexctl library optimize <library-name>`: Optimize media files in a library
    * **Playback Control:**
        * `plexctl client play <client-name>`: Start playback on a specific client
        * `plexctl client pause <client-name>`: Pause playback on a specific client
        * `plexctl client stop <client-name>`: Stop playback on a specific client
        * `plexctl client seek <client-name> <position>`: Seek to a specific position in the currently playing media
        * `plexctl client get <client-name>`: Get information about the currently playing media on a client
* **2.2 Flags**
    * **Global**
      * `--verbosity <verbosity-level>`: Override the default verbosity. Can be `debug`, `info`, `warning`, `error`, `critical`, `none`
      * `--version`: Prints the version of plexctl
* **2.3 Output**
    * Output will be formatted in a human-readable and machine-parseable format (e.g., JSON, YAML)
    * Error messages will be displayed clearly and concisely

**3. API Interaction**

* **3.1 Authentication**
    * `plexctl` will use the Plex Token for authentication with the Plex Media Server API
    * Users will need to provide their Plex Token as a command-line argument or environment variable
* **3.2 API Endpoints**
    * The tool will utilize the self-hosted Plex Media Server REST API endpoints for all interactions
    * This will be user-configurable as each installation can be unique

**4. Implementation**

* **4.1 Programming Language**
    * Go will be used for it's strong typing and it's readability.
* **4.2 Libraries**
    * [`github.com/jrudio/go-plex-client`](github.com/jrudio/go-plex-client): For making HTTP requests to the Plex Media Server API
    * [`github.com/urfave/cli/v3`](https://github.com/urfave/cli): For parsing command-line arguments
    * [`github.com/dgraph-io/badger/v3`](github.com/dgraph-io/badger/v3): For storing non-sensitive configuration data
    * `encoding/json` or `github.com/goccy/go-yaml`: For handling output data formats
* **4.3 Structure**
    * The code will be organized into modules for better maintainability:
        * `plexctl/`: Main package
            * `go.mod`: Package initialization
            * `main.go`: Entrypoint for the tool
            * `commands/`: Subdirectory for individual commands
                * `server.go`: Commands for server management
                * `library.go`: Commands for library management
                * `playback.go`: Commands for playback control
            * `utils.go`: Utility functions (e.g., argument parsing, output formatting)

**5. Testing**

* **5.1 Unit Tests**
    * Unit tests will be written to verify the functionality of each command and API interaction
* **5.2 Integration Tests**
    * Integration tests will be conducted to ensure the tool interacts correctly with a running Plex Media Server

**6. Future Enhancements**

* **6.1 Features**
    * Add support for managing users and roles
    <!-- * Implement features for transcoding and media management -->
    <!-- * Add support for controlling playback on multiple clients simultaneously -->
* **6.2 User Interface**
    * Improve command-line completion and tab completion support
    * Add support for interactive prompts and configuration files

**7. Conclusion**

This document provides a high-level overview of the design for the `plexctl` command-line tool. The implementation will follow this design and be iteratively improved based on user feedback and identified needs.

**Key Considerations:**

* **Error Handling:** Implement robust error handling to gracefully handle API errors, network issues, and invalid user input
* **Security:** Ensure proper authentication and authorization mechanisms are in place to protect user data
* **Documentation:** Create comprehensive documentation for users and developers
