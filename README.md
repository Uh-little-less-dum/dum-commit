# Dum-commit

A super simple conventional commits utility.

This can be added to the `.husky/prepare-commit-msg` or `.git/hooks/prepare-commit-msg` hooks with:

```sh
dum-commit hook $1
```

This will pass the commit temporary file to `dum-commit` which will open a cli form everytime you make a commit, and hopefully enforce some sharable, structured commits. This does nothing to validate the commit message, so you can still get around this structure, but the cli itself will output commits that abide by general conventional commit practices.

You can configure this commit using the root package.json file with the following config:

```ts
type DumCommitConfig = {
    // ... the rest of you package.json
    "dum-commit": {
        "height": 8,
        "type": {
            "height": 8, // will override the general height
            "opts": [
                "wip",
                "feat",
                "chore",
                ...
            ]
        },
        "scope": {
            "height": 8, // will override the general height
            "opts": [
                "api",
                "website",
                "docs",
                ...
            ]
        }
    }
}
```

- All options are optional, and in fact, the entire config is optional, although that kinda defeats the point.
- You can omit the `opts` fields and a text input will be used.
- The height is only applied to select lists, as inputs will fit automatically thanks to the awesome work of the people at [charm](https://github.com/charmbracelet).
- You can provide a global config using the `--config` or `-c` flag which points to an equivalent json file.

## Uh Little Less Dum

Checkout the main project that this app was built to support, [UhLittleLessDum](https://uhlittlelessdum.com), a free, open-source, academic oriented note taking framework that was built to contain my own research and is now being released as a free tool that can be easily extended.

### Thanks to:

- [Charm](https://github.com/charmbracelet)
- [Cobra](https://github.com/spf13/cobra)
- [Viper](https://github.com/spf13/viper)

> As a matter of fact, the only thing that's really different from the Charm example is that I'm grabbing some info from a config with viper, and exposing a global command with cobra. Honestly, I feel a little guilty even releasing this having done almost none of the work.
