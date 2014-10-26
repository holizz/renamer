# renamer

Rename files using your text editor. Inspired by [massren](https://github.com/laurent22/massren) but with fewer features.

## Installation

    go get -u github.com/holizz/renamer

## Usage

    renamer file1 file2 file3

Your editor will open with the following lines:

    file1
    file2
    file3

Modify them (but don't move, delete, or insert any lines!) and save/quit and the files will be renamed.

Uses the EDITOR or VISUAL environment variables to determine which editor to use. Falls back to "vi" if neither option is set.

## License

MIT
