# emojipacks

[![CircleCI](https://circleci.com/gh/Code-Hex/go-emojipacks.svg?style=svg)](https://circleci.com/gh/Code-Hex/go-emojipacks)

> CLI to bulk upload emojis to your Slack!

## Install

### binary

You can get binary from [releases](https://github.com/Code-Hex/go-emojipacks/releases)

### go get

```bash
$ go get github.com/Code-Hex/go-emojipacks/cmd/emojipacks
```

## Usage

There is only one command:

```bash
$ emojipacks
```

It'll ask you a few questions:

```bash
Slack subdomain: helloworld
Email address login: code-hex@codehex.dev
Password: *********
2FA Code: 123456  #  if 2FA is enabled
Path or URL of Emoji yaml file: ./packs/futurama.yaml
```

Almost this command usage the same as https://github.com/lambtron/emojipacks

## Optionally Pass Command Line Parameters

This will allow for easier batch uploading of multiple yaml files

```bash
$ emojipacks -s <subdomain> -e <email> -p <password> -y <yaml_file>
```

## Emoji Yaml File

Also note that the yaml file must be indented properly and formatted as such:

```yaml
title: food
emojis:
  - name: apple
    src: http://i.imgur.com/Rw0Vlda.png
  - name: applepie
    src: http://i.imgur.com/g4RU1fM.png
```

..with the `src` pointing to an image file. According to Slack:

- Square images work best
- Image can't be larger than 128px in width or height
- Image must be smaller than 64K in file size

### Emoji Aliases

It hasn't supported yet.

## Emoji packs

[See here](https://github.com/lambtron/emojipacks/blob/master/README.md#emoji-packs)
