## What is this fuckery?

This includes structured configuration files in urfavecli, and makes it play by the rules. Such rules being:

> The precedence for flag value sources is as follows (highest to lowest):
>
>1. Command line flag value from user
>2. Environment variable (if specified)
>3. Configuration file (if specified)
>4. Default defined on the flag

It contains a hardcoded file location, however this is easily replaced with Viper loaders. By doing this pattern, we are simply providing a set of sane defaults that should prevent an app from crashing.

## Supported Use Cases
- Set all flag default values from a config file
- Set a few flag defaults from a config file, let unset values use sane defaults
- Set all flag values to sane defaults

This works because it only hijacks the DefaultValue for the flags, which means it won't be touched unless flags are parsed, in which case it maintains the same order that urfavecli promises, as god intended.