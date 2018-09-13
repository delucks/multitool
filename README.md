# multitool

This is a project composed of many small tools executed under a single binary, like `busybox`. I'm creating it with the intention of implementing more command-line utilities in golang and moving some of my complex shell scripts into Go with more testing.

## Tools Available

| Tool Name | Description |
| --------- | ----------- |
| `basejump` | Convert an integer from base 2, 8, 10, or 16 to an arbitrary base |
| `colors` | Display the current terminal color scheme |
| `emojis` | Generate a TSV of all emojis with :alias: names |
| `log` | A simple shell logging utility that outputs ISO8601 timestamps and terminal colors |
| `suggest-fc` | A simple spellchecking app to suggest a command you may have mistyped |

This project is based on the [leatherman](https://github.com/frioux/leatherman) project by fREW- thanks for the inspiration!

# TODO

### New Tools
- yamltojson
- jsontoyaml

### From dotfiles
- backlight
- fileset ("settool")
- iploc()
- histogram()

### From scripts
- digdug.sh
  May as well make a generic "make me some ascii art"
- x86.py
- Imgur down/uploader
- lenny
