# Pokedex Terminal/REPL
## Getting started as a user

Prerequisites:
- You must have Go installed, at least version 1.24.1

To get started, run `go build`, after that you should have an executable named pokedex | pokedex.exe

After running it, you should see a prompt, `Pokedex >`
From here, you can run the `help` command to see all of the currently accepted commands.

## Worth noting
Currently, the "captured" pokemon in your pokedex only last as long as you keep the program running.
There is no persisting memory. That might get added later on, or would be fairly trivial for you to add
if you'd like.

Also, If you experience any issues, please submit an issue report so I can see to them being fixed.

Thank you to the developers at Pokeapi.co for building such an extensive Rest Api. This Api, provided
by them for free, is what facilitates the fetching of all pokemon data.

## Contributing or Branching
If you'd like to develop on this, I'll break down how I organized the code

Firstly, I've only used std lib packages in this project, so no dependencies.

Next, I've implemented two Api's in the internal/ folder. One for the http.Client, 
providing you with functions like `GetPokemon` to fetch pokemon data, and more. It has a
struct call `PokeClient` that contains a `client` http.Client and a `cache` pokeapi.Cache,
from the other Api in this project.

There is also an Api for the cache implemented within the pokeapi package. It exposes 
functions for getting and adding to the cache, and encapsulates all of the safety management.
You'll want to use the `NewCache` function, as it starts a "reapLoop", to continuously remove
old items in the cache.

All the commands are in the main package of the project. They're pretty easy to read and understand.
To add commands, you'll have to add the commmand to the return map of the `GetCommands()` func.
