---
# basic details
title: "Tic-Tac-Total Carnage!"
description: "A mindbending version of tic-tac-toe where each space on your board is another nested tic-tac-toe game."

# hero stuff
subtitle: "( has absolutely nothing to do with carnage )"
image: "x"
hasPixelImage: false

# game stuff
game: "https://tic-tac-total-carnage.netlify.com/"
repo: "https://github.com/ryannhg/tic-tac-total-carnage"
players: "2"
genre: Local multiplayer
hasControllerSupport: false

# 
date: 2019-08-14T00:44:40-06:00
---

### How to play

The game starts with __Player 1__. They can place an "X" on any position on the grid. For example, they might place an "X" in the middle square of the bottom-left game.

Because Player 1 placed their "X" in the __middle square__, Player 2 has to play their piece in the __middle game__.

On their turn, Player 2 might place an "O" in the __top-right square__ of the middle game. This forces Player 1 to play in the __top-right game__.

Getting tic-tac-toe (three in a row) on __one of the 9 tiny boards__ will give a player a large "X" or "O" on the __overall board__. 

The game ends when a player gets __three in a row__ on the overall game. ( or their is a draw! )

### Confused?

Good– it's kinda hard to explain with text. But on the bright side, this version of the game will __automatically disable__ invalid moves.

It also takes care of alternating X and O after each turn.

### One more thing!

If you send a player to a tiny board that __already has a large piece__, they can play their next move in __any tile__. So be careful not to give them all that freedom– they will likely crush your hopes and dreams.