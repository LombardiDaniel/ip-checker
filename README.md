<!-- PROJECT LOGO -->

<br />
<p align="center">
  <!-- <a href="LOGO.png">
    <img src="LOGO.png" alt="Logo" width="230" height="auto">
  </a> -->

  <h3 align="center">Home IP Checker</h3>

  <p align="center">
    A simple script to notify you when your home ip changes.
    <br />
    <a href="https://github.com/LombardiDaniel/ip-checker"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/LombardiDaniel/ip-checker/issues">Report Bug</a>
    ·
    <a href="https://github.com/LombardiDaniel/ip-checker/issues">Request Feature</a>
  </p>
</p>

## Table of Contents

- [Table of Contents](#table-of-contents)
  - [About the Project](#about-the-project)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Instalation](#instalation)
  - [Usage](#usage)
  - [License](#license)

### About the Project

The idea behind this repo is a simple way to guarantee that you'll always able to connect to you [PiVPN](https://www.pivpn.io) even if your internet provider changes you public IP. By sending a notification via e-mail and/or Telegram whenever the IP of you Raspberry Pi changes, you'll never be left out of your home internet!

The system is a very simple python script that uses [amazon AWS](https://checkip.amazonaws.com) to get the current IPv4 Address of the device (must be running in the same machine as the VPN) it's running in.

### Getting Started

#### Prerequisites

To maximize compatbility, you will only need [docker](https://www.docker.com/) and docker-compose to be installed on your computer. But if you wish to not run it as a container, feel free to check the code and run it simply on the CLI.

For the notifications, you will need an e-mail to be used as the "notificator" as well as a configured Telegram Bot. Remember to allow less secure apps on the e-mail end. For the Telegram Bot, you'll need 2 things: the Chat ID of the people that will receive the
notification, as well as the bot ID. You can learn about them [here](https://core.telegram.org/bots).

#### Instalation

Since the script is containerized, just run `git pull https://github.com/LombardiDaniel/ip-checker`, `cd ip-checker`. Create a `.env` file with the variables described below.

You can also just check out the docker image at [docker-hub](https://hub.docker.com/repository/docker/lombardi/ip-checker/general).

### Usage

The only configuration needed is to set the environment variables in the `.env` file. Configure your email and/or Telegram key in the `.env` file (separate multiple addresses by using `,`) and run `docker-compose up -d`. It logs to `stdout`, so you can easily check the logs via [portainer](https://www.portainer.io).

|    Variable Name    | Meaning                                     | Separator |
| :-----------------: | ------------------------------------------- | :-------: |
|      MAIL_LIST      | List of emails to send notifications to     |    `,`    |
| TELEGRAM_CHAT_LIST  | List of chat keys to send notifications to  |    `,`    |
|  TELEGRAM_BOT_KEY   | Bot API Keys to send notifications to       |     -     |
|   EMAIL_HOST_USER   | E-mail that will be used to send out emails |     -     |
| EMAIL_HOST_PASSWORD | Password of Host e-mail                     |     -     |

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
