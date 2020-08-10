# Contributing to firmwares list

## How to add new firmware to list?
1. You had to know firmware name, last build date and link to last build
1. Open list with needed Android version
1. Add new line
1. Follow with CSV title. If you don't know something or it is not exist, e.g. XDA link, just pass it. If you forget some field, the table will broke.

   Example: `BestSuperIOR,,,2020-08-10,,,google.com/favicon.ico,,,`

### Additional
* in `Treble` field use only `Yes`/`No`/` ` (nothing, if you don't know)
* please use ISO date
* don't commit `https://`, `www.` and last slash at links
* try to use the shortest view, e.g. for SourceForge: `download.sourceforge.com/PROJECT NAME/FILENAME`
* save files in UTF-8 with CRLF
* keep last newline on [EOF](https://goo.gl/search/end%20of%20file)

## How to add new Android version / other OS?

1. Create `.csv` file started with `santoni-`, then place Android version letter or OS name.

   Examples: `santoni-r.csv`, `santoni-ubuntu-touch.csv`.
1. At the first line add text below:

   ```
   Name,Treble,Author,Date,Website,Folder,Download link,XDA link,Telegram link,4PDA post link
   ```
1. Populate file with firmwares
1. Add to `docs/index.html` new list item. `.html` name is same with `.csv` name (generated with Actions)

   Examples: `<li><a href="santoni-r.html">santoni-r</a></li>`, `<li><a href="santoni-ubuntu-touch.html">santoni-ubuntu-touch</a></li>`
1. Request your pull
