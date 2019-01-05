# multitool

[![Build Status](https://travis-ci.org/delucks/multitool.svg?branch=master)](https://travis-ci.org/delucks/multitool)

This is a project composed of many small tools executed under a single binary, like `busybox`. It holds a bunch of small programs that perform one task, all written in golang. It's a target for porting complex shell scripts into go, enabling more testing and speed. This was inspired during a conversation about the [leatherman](https://github.com/frioux/leatherman) project by fREW, which operates in a similar way.

## Tools Available

| Tool Name | Description |
| --------- | ----------- |
| `basejump` | Convert an integer from base 2, 8, 10, or 16 to an arbitrary base |
| `colors` | Display the current terminal color scheme |
| `emojis` | Generate a TSV of all emojis with :alias: names |
| `log` | A simple shell logging utility that outputs ISO8601 timestamps and terminal colors |
| `suggest-fc` | A simple spellchecking app to suggest a command you may have mistyped |
| `jsontoyaml` | Convert a JSON structure to the equivalent YAML |
| `yamltojson` | Convert a YAML structure to the equivalent JSON (assuming you're using compatible features) |

## TODO

### From dotfiles
- fileset ("settool")
- iploc()

### From scripts
- digdug.sh
  May as well make a generic "make me some ascii art"
- x86.py
- Imgur down/uploader
- lenny
