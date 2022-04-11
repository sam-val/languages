import random
import string

hangman_img = [
    0 , '\u213c','-','-','-','-',0,
    0 , '|', 0, 0,'|', 0, 0,
    0 , '|', 0, 0,'O', 0, 0,
    0 , '|',0,'/','|','\\',0,
    0 , '|', 0,0, '|',  0, 0,
    0 , '|', 0,'/',0,'\\', 0,
    0 , '|', 0,0,0,0, 0,
    0,'\u22a5', 0,0,0,0, 0,
]

guesses = word = right_letters = MAX_GUESSES = None
word_list = [
    'chicken wing',
    'breaking bad',
    'walter white',
    'inception',
]

def init_state():
    global guesses, word, right_letters
    guesses = [(5,5),(3,5),(4,4),(5,3),(3,3),(4,3),(4,2),None]
    MAX_GUESSES = len(guesses)
    word = random.choice(word_list)
    right_letters = []


def draw_hangman():
    if len(guesses) == MAX_GUESSES:
        return
    w, h = 7,8
    img = hangman_img
    for y in range(h):
        for x in range(w):
            if x == 0:
                print('')
            # print(x,y)
            val = img[w*y+x]
            if val and (x,y) not in guesses:
                print(val, end='')
            else:
                print(' ', end='')

def drawword():
    for x in word:
        if x in right_letters:
            print(f" \u0332{x} ", end='')
        else:
            if x == ' ':
                print(f" {x} ", end='')
            else:
                print(' _ ', end='')

    print("")


def getwinorlose():
    if len(guesses) == 0:
        return 'lose'
    if len(right_letters) == len(set(word.replace(' ', ''))):
        return 'win'


init_state()
if __name__ == '__main__':
    while True:
        # draw
        print("\n----------------------------\n\n\n")
        draw_hangman()
        drawword()

        # check win/lose
        winorlose = getwinorlose()
        if winorlose:
            print('you won' if winorlose == 'win' else "you've lost")
            again = input('play again? y or n: ')
            if again.strip() != 'y':
                break
            else:
                init_state()
                continue

        letter = ' '
        while letter not in string.ascii_lowercase or letter in right_letters:
            letter = input("\nEnter a valid English char: ")

        if letter in word:
            right_letters.append(letter)
        else:
            guesses.pop()




