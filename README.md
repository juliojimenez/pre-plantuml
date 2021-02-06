# pre-plantuml

A pre-commit hook that renders PlantUML diagram source in Markdown image tags. 

## Install pre-commit

If this is your first time using pre-commit, start [here](https://pre-commit.com/#install)

## Add a Configuration

At the root of your repo, create `.pre-commit-config.yaml` and add the following:

```yaml
repos:
- repo: https://github.com/juliojimenez/pre-plantuml
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

![diagrams/src/example1.pu](http://www.plantuml.com/plantuml/png/~h407374617274756d6c0a416c696365202d3e20426f623a2041757468656e7469636174696f6e20526571756573740a426f62202d2d3e20416c6963653a2041757468656e7469636174696f6e20526573706f6e73650a0a416c696365202d3e20426f623a20416e6f746865722061757468656e7469636174696f6e20526571756573740a416c696365203c2d2d20426f623a20416e6f746865722061757468656e7469636174696f6e20526573706f6e73650a40656e64756d6c)

`diagrams/src/example2.pu`

![diagrams/src/example2.pu](http://www.plantuml.com/plantuml/png/~h407374617274756d6c0a416c696365202d3e20426f623a2041757468656e7469636174696f6e20526571756573740a426f62202d2d3e20416c6963653a2041757468656e7469636174696f6e20526573706f6e73650a0a416c696365202d3e20426f623a20416e6f746865722061757468656e7469636174696f6e20526571756573740a416c696365203c2d2d20426f623a20416e6f746865722061757468656e7469636174696f6e20526573706f6e73650a40656e64756d6c)

## Supported Files

### PlantUML

- .pu
- .puml
- .plantuml
- .iuml
- .wsd

### Markdown

- .md