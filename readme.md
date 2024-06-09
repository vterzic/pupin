# Pupin

Pupin is a CLI tool that renders interactive CLI menu based on the provided configuration json. Something like interactive command menu.

Every tree can contain multiple options. Each option can be a subtree or a command.
If user selects subtree, subtree options will appear.
If user selects command, command will be executed.

## Usage
To run your custom menu, you need to:
- create a configuration json (see Configuration)
- run `pupin run <config-path>`

### Configuration

Nothing special really. 

Check [config-example.json](https://github.com/vterzic/pupin/blob/main/config-example.json).

Make sure that your configuration is validated against [config-schema.json](https://github.com/vterzic/pupin/blob/main/config-schema.json) (you can use [jsonschemavalidator](https://www.jsonschemavalidator.net/)).

## Might come in handy
```
mkdir ~/pupin
nano my-config.json
cp pupin my-config.json ~/pupin
#(name alias whatever you like)
echo "alias ppn='~/pupin/pupin run ~/pupin/config-example.json'" >> ~/.zshrc
source ~/.zshrc
```

## ❗️Don't trust my binaries
It is never a good idea to use a precompiled binary with sensitive data.

Instead, compile your own binaries by running `go build .`.

## TODO
- runtime json validation
- nonExit flag for commands
- colors for submenus
- tests (doubt it)
