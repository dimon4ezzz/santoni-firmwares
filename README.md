# Santoni Firmwares list
[![Pages](https://img.shields.io/badge/Pages-%2Fdocs-green)](/docs)

[Marshmallow](santoni-m.csv) • [Nougat](santoni-n.csv) • [Oreo](santoni-o.csv) • [Pie](santoni-p.csv) • [Q](santoni-q.csv) • [R]()

## CSV Fields list
1. Name (**required**) — firmware name (without version name, suffixes `OS`, `ROM`)
1. Treble — is firmware treble (`Yes` or `No`)
1. Author — real or supposed firmware author (team)
   * please, use full name, nicknames divided with `/`
1. Date (**required**) — last update date
   * please, use ISO date: `yyyy-mm-dd`
1. Website — official landing page
1. Folder — updatable folder with firmware ROMs
1. Download link (**required**) — link to download ROM at specified date
   * please, for sourceforge use typical `download.sourceforge.com/PROJECT NAME/FILENAME`, see examples in files
1. XDA link — link to XDA topic
   * please, use shorten view, matches regex: `^forum.xda-developers.com\/.+\/-t\d+$`
1. Telegram link — link to Telegram Announcement Channel or Group
   * please, use shorten view, matches regex: `^t\.me\/\w+$`
1. 4PDA link — link to 4PDA post
   * please, use shorten view, matches regex: `^4pda\.ru\/forum\/index\.php\?act=findpost&pid=\d+$`

## Links
* 4PDA forum: https://4pda.ru/forum/index.php?showtopic=821670
* XDA Developers forum: https://forum.xda-developers.com/xiaomi-redmi-4x/development
* **Redmi 4X** tag on Xiaomi Firmware: https://xiaomifirmware.com/tag/redmi-4x/
* **santoni** tag on Android File Host: https://androidfilehost.com/?w=developers&did=2459

## Cleaning CSV
There are regular expressions, used for cleaning csv files (make it smaller). Used `$1` notation reference group for VS Code.
1. `(?:https:\/\/)?sourceforge\.net\/projects\/([^\/]+)\/files[^,]*\/([^\/]+.zip)\/download` to `download.sourceforge.net/$1/$2`
1. `(?:https:\/\/)?forum\.xda-developers\.com\/([^,]*)\/[^,]*-t(\d+)` to `forum.xda-developers.com/$1/-t$2`
1. `(?:https:\/\/)?4pda\.ru\/forum\/index.php\?s=&showtopic=[\d]+&view=findpost&p=(\d+)` to `4pda.ru/forum/index.php?act=findpost&pid=$1`
1. remove `https://`
1. `\/,` to `,`
