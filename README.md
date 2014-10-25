#Hueman

### The worlds shittiest hue client.

Working on making it less shitty.

To begin config you can press the button on your bridge, then run
`hueman --connect NAME_OF_BRIDGE`

Once it has been setup a config file ~/.hueman will be created.

You can then set the color and brightness of all your lights on that bridge


Set full intensity white:
`hueman --color "#FFFFFF"`

Set brightness to 50%
`hueman --brightness 50`

Set via HSL values (full intensity red)
`hueman --hue 0 --sat 100 --brightness 100`
