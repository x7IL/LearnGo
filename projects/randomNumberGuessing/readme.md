# Random Number Guessing Game

## Description

This is a command-line guessing game where the player must guess a randomly generated number within a specified range. The game provides feedback on whether the guess is higher or lower than the target number and adjusts the guessing range accordingly.

## How It Works

1. **Game Setup**:
    - A random number between `0` and `1000` is generated.
    - The player is prompted to guess the number.

2. **Gameplay**:
    - The player enters their guess.
    - The game compares the guess with the random number:
        - If the guess is higher than the random number, the upper limit of the guessing range is adjusted.
        - If the guess is lower than the random number, the lower limit of the guessing range is adjusted.
    - The game continues until the player guesses the correct number.

3. **Feedback**:
    - The game informs the player whether their guess is too high or too low and updates the guessing range.
    - When the correct number is guessed, the game congratulates the player and reveals the number.

## Requirements

- Go programming language installed (version 1.x or later).
- Basic knowledge of command-line operations.

## Additional Notes

- The random number is generated between `0` and `1000`. You can modify this range by adjusting the relevant variable in the code.
- For debugging purposes, you might uncomment lines to display the generated number.
