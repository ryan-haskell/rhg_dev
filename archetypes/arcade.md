---
# basic details
title: "{{ replace .Name "-" " " | title }}"
description: "Description of game"

# hero stuff
subtitle: "( shows in hero, under title )"
image: "wave"
hasPixelImage: false

# game stuff
game: "#url-to-game"
repo: "#url-to-repo"
players: "1"
genre: Game
hasControllerSupport: false

# 
date: {{ .Date }}
draft: true
---

### Controls

...