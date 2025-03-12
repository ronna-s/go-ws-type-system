# The Go type system and interfaces

### Before we begin 
- CLONE ME!!! [github.com/ronna-s/go-ws-type-system](github.com/ronna-s/go-ws-type-system)
- If you haven't already responded, please let me know how much experience you have with Go [here](https://app.sli.do/event/cNENGWWL5QUCR3xYJ9LELP/embed/polls/2d3b672b-998e-4756-a814-1e4f9400bd8c).
 - Install Go [here](https://go.dev/dl/)

## Task: build P&P Game
P&P™ stands for Platforms and Programmers.

P&P is a game in which a band of developers tries to take on the ultimate villain: PRODUCTION.

Today you are going to parts of the game engine, the characters (developers), and how they interact with PRODUCTION.

Game starts:
- The band of developers starts with 10 points of satisfaction each and a few basic skills.
- The band starts with 10 gold coins, which can be used to buy objects such as coffee or banana (you'll understand why in a sec), pizza, etc.
- PRODUCTION is "calm".
- Every band has a minion. Minions love PRODUCTION, the ultimate villain, and will do anything to serve it. 
- A minion's only skill is to cause bugs. A minion cannot learn new skills. A minion can be distracted using a banana.


Game flow:
- PRODUCTION state can go from "calm", "annoyed", "enraged", up to "legacy".
- The game ends if there are no more players (except for Minion, minions don't count since you can't fire minions).
- The game is won if production is calm and the band played more than 10 rounds.
- In each iteration a player has to select a move to make against PRODUCTION.
- PRODUCTION will react to the option played.
- If an unsuccessful move was made against PRODUCTION while in legacy mode, one player will be terminated at random (except for Minion, minions can't be fired).

### Your mission:
- Implement Minion (in `pnpdev/player.go`):
    - Minion can create bugs.
    - If there's a budget you can choose to buy a Banana instead (costs 1 gold coin) to distract Minion from creating a bug by buying a banana and have him eat it.
    - Bugs always upset PRODUCTION.
    - A minion doesn't have a method `Alive() bool` - we need to figure out a solution for this.
- Implement Rubyist:
    - Rubyist has one option - dark magic, which always upsets PRODUCTION unless production is in legacy mood, then it's surprisingly effective.
- Implement the game loop.
- Plug in the players to your game.
- Implement the Gopher player:
    - Gopher can create a feature.
    - Features always make PRODUCTION happy.

### You're free now to define your own game:
- Add as many more players as you'd like (for instance, a manager who can buy Pizza for everyone).
- You can addd a satisfaction system to your game.
- Add leveling up of the players.
- Whatever strikes your fancy.

