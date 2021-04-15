![Gardener](img/gardener128x128.png)

# Gardener

[![Go Report Card](https://goreportcard.com/badge/github.com/devops-kung-fu/gardener)](https://goreportcard.com/report/github.com/devops-kung-fu/gardenre) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/devops-kung-fu/gardener)

A pre-commit hook that renders PlantUML diagram source in Markdown image tags. 

## Usage with Hookz

Gardner is best used with [Hookz](https://github.com/devops-kung-fu/hookz). Follow the instructions to install Hookz and then configure your ```.hookz.yaml``` file to add the following action:

__NOTE:__ Ensure that for the ```url``` value you select the appropriate Gardener release for your platform and architecture.

``` yaml
version: 2.1.1
hooks:
  - type: pre-commit
    actions:
      - name: "Generate Images for PlantUML"
        url: https://github.com/devops-kung-fu/gardener/releases/download/v0.1.4/pre-plantuml-0.1.4-linux-amd64
        args: ["deflate"]
```

## Using with pre-commit

If this is your first time using pre-commit, start [here](https://pre-commit.com/#install)

## Add a Configuration

At the root of your repo, create `.pre-commit-config.yaml` and add the following:

```yaml
repos:
- repo: https://github.com/devops-kung-fu/gardener
  rev: v0.0.26
  hooks:
  - id: pre-plantuml
```

## Install pre-plantuml

```shell
$ pre-commit install
```

## Update pre-plantuml

```shell
$ pre-commit autoupdate
```

## Usage

Easy! Say we have our diagram source in `diagrams/src/alicenbob.pu`.

```
@startuml
Alice -> Bob: Authentication Request
Bob --> Alice: Authentication Response

Alice -> Bob: Another authentication Request
Alice <-- Bob: Another authentication Response
@enduml
```

And we want to display it in our README.md.

```
# My README

Below is the image tag that pre-plantuml will update on pre-commit!

![diagrams/src/alicenbob.pu]()

Next!
```

You just insert the path to the diagram source in the **alt** text (the part in brackets) of the image. pre-plantuml takes care of finding diagram source files, markdown files, and image source to update (the part in parentheses).

## Examples

Here are some examples from processed from [diagrams/src/example1.pu](diagrams/src/example1.pu) and [diagrams/src/example2.pu](diagrams/src/example2.pu).

`diagrams/src/example1.pu`

![diagrams/src/example1.pu](http://www.plantuml.com/plantuml/png/1C3HZSCW40JGVwgO1cZ0Ebd19RW3u4PY9RARcA7_lDTIVRJVCvLfdSWdhcW7ojQWotgLXUFcTtCfNT6GyuaohVD0sHfqMQ-oSDnSd_35bCgqJkGJLxG3nKE33-hMeCjwbONZvdTpAPLfdVZB6LUq0yL3Wm_grg3BUfM5u-RwX2-c5_r_lsVw0G00__y1003__m==)

`diagrams/src/example2.pu`

![diagrams/src/example2.pu](http://www.plantuml.com/plantuml/png/1C3HZSCW40JGVwgO1cZ0Ebd19RW3u4PY9RARcA7_lDTIVRJVCvLfdSWdhcW7ojQWotgLXUFcTtCfNT6GyuaohVD0sHfqMQ-oSDnSd_35bCgqJkGJLxG3nKE33-hMeCjwbONZvdTpAPLfdVZB6LUq0yL3Wm_grg3BUfM5u-RwX2-c5_r_lsVw0G00__y1003__m==)

## Supported Files

### PlantUML

- .pu
- .puml
- .plantuml
- .iuml
- .wsd

### Markdown

- .md

## Credits

A big thank-you to our friends at [Freepik](https://www.freepik.com) for the Gardener logo.